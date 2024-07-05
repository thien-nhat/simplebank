package token

import "time"

// TokenMaker is interface for managing tokens
type TokenMaker interface {
	// CreateToken create a new token
	CreateToken(username string, duration time.Duration) (string,*Payload, error)

	// VerifyToken verify a token
	VerifyToken(token string) (*Payload, error)
}