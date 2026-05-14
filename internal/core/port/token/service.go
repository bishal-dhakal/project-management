package token

import "time"

type Service interface {
	GenerateToken(userID string, email string, duration time.Duration)(string, error)
	ValidateToken(token string) (userID string, email string, err error)
}