package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different type of errors
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")

)
// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Valid implements jwt.Claims.
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// NewPayload creates a new token Payload
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
