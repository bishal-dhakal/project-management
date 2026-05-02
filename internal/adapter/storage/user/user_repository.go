package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	// userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
	"github.com/bishal-dhakal/project-management/internal/core/domain/user"
	userport "github.com/bishal-dhakal/project-management/internal/core/port/user"
)

type userRepository struct {
	db *sqlx.DB
}

// FindByEmail implements [ports.UserRepository].
func (u *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	panic("unimplemented")
}

// Save implements [ports.UserRepository].
func (u *userRepository) Save(ctx context.Context, user user.User) error {
	panic("unimplemented")
}

func NewUserRepository(db *sqlx.DB) userport.UserRepository {
	return &userRepository{db: db}
}
