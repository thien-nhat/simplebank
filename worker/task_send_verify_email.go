package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db "github.com/thien-nhat/simplebank/db/sqlc"
	"github.com/thien-nhat/simplebank/util"
)

const TaskSendVerifyEmail = "task:send_verify_email"
type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueued task")

	return nil
}

func (processor *RedisTaskProcessor)ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload: %w", asynq.SkipRetry)

	}
	
	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exit: %w",asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	verifyEmail, err := processor.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Username: user.Username,
		Email: user.Email,
		SecretCode: util.RandomString(32),
	})

	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	subject := "Welcome to SimpleBank"
	verifyUrl := fmt.Sprintf("http:/localhost:8080/v1/verify_email?email_id=%d&secret_code=%s", 
		verifyEmail.ID, verifyEmail.SecretCode)

	content := fmt.Sprintf(`Hi %s, welcome to SimpleBank.<br/> 	
	Thank you for registering with us!<br/> 
	To verify your email, please <a href="%s">Click here</a>.`, user.Username, verifyUrl)
	
	to := []string{user.Email}
	err =processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("processed task")

	return nil
}