package user

import (
	"context"
	"database/sql"
	"errors"

	userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
	userport "github.com/bishal-dhakal/project-management/internal/core/port/user"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

// FindByEmail implements [ports.UserRepository].
func (u *userRepository) FindByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	var userRow userdomain.User
	err := u.db.QueryRowContext(ctx, "SELECT * FROM users where email = $1", email).Scan(&userRow.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // caller should check for nil
		}
		return nil, err
	}

	return &userRow, nil
}

// Save implements [ports.UserRepository].

func (u *userRepository) Save(ctx context.Context, user userdomain.User) (sql.Result, error) {
	// Ensure the SQL string is wrapped in backticks or quotes
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`

	result, err := u.db.ExecContext(ctx, query, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUserRepository(db *sqlx.DB) userport.Repository {
	return &userRepository{db: db}
}
