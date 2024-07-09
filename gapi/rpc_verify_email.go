package gapi

import (
	"context"

	db "github.com/thien-nhat/simplebank/db/sqlc"
	"github.com/thien-nhat/simplebank/pb"
	"github.com/thien-nhat/simplebank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := ValidateVerifyEmailRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}
	txResult, err := server.store.VerifyEmailTx(ctx ,db.VerifyEmailTxParams{
		EmailId: req.GetEmailId(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email: %v", err)
	}
	rsp := &pb.VerifyEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	return rsp, nil
}

func ValidateVerifyEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmailID(req.GetEmailId()); err != nil {
		violations = append(violations, fieldViolation("email_id", err.Error()))
	}
	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err.Error()))

	}
	return violations
}