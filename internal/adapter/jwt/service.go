package jwt

import (
	"time"
	jwtport "github.com/bishal-dhakal/project-management/internal/core/port/token"
)

type jwtService struct {
	secretKey string
}

func NewService(secret string) jwtport.Service {
	return &jwtService{secretKey: secret}
}

func (s *jwtService) GenerateToken(userID string, email string, duration time.Duration)(string, error){
	return "signed.jwt.token", nil
}

func (s *jwtService) ValidateToken(token string) (string, string, error){
	return "user-123", "test@test.com", nil
}
