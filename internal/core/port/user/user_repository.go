package userport

import (
	"context"
	userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
)

type UserRepository interface {
	Save(ctx context.Context, user userdomain.User) error
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
}