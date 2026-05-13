package userservice

import (
	"context"
	"errors"

	userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
	authport "github.com/bishal-dhakal/project-management/internal/core/port/auth"
	userport "github.com/bishal-dhakal/project-management/internal/core/port/user"
	hashPassword "github.com/bishal-dhakal/project-management/internal/core/util"
)

type userService struct {
	repo userport.UserRepository
}

func New(repo userport.UserRepository) authport.AuthService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, email, password string) (string, error) {
	// TODO: validate input
	if email == " " {
		return "", errors.New("empty name")
	}

	// TODO: check email exists
	_, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// TODO: hash password
	hashedPassword, err := hashPassword.HashPassword(password)
	if err != nil {
		return "", err
	}

	var user userdomain.User

	user.Email = email
	user.Password = hashedPassword

	// TODO: save user
	_, err = s.repo.Save(ctx, user)
	if err != nil {
		return "", err
	}

	return "user created successfully", nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	// TODO: find user by email
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Add this check to prevent the panic
	if user == nil {
		return "", errors.New("user not found")
	}

	// TODO: compare password
	verified := hashPassword.VerifyPasswordHash(password, user.Password)

	if !verified {
		return " ", errors.New("email/password incorrect")
	}
	// TODO: generate JWT
	return "", nil
}
