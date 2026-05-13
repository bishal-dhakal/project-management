package authport

import (
	"context"
	// userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
}