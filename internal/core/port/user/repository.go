package user

import (
	"context"
	"database/sql"

	userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
)

type Repository interface {
	Save(ctx context.Context, user userdomain.User) (sql.Result, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
}
