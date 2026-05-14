package userservice

import (
	"context"
	"errors"
	"time"

	userdomain "github.com/bishal-dhakal/project-management/internal/core/domain/user"
	authport "github.com/bishal-dhakal/project-management/internal/core/port/auth"
	userport "github.com/bishal-dhakal/project-management/internal/core/port/user"
	hashPassword "github.com/bishal-dhakal/project-management/internal/core/util"
	token "github.com/bishal-dhakal/project-management/internal/core/port/token"

)

type userService struct {
	repo userport.Repository
	token token.Service
}

func New(repo userport.Repository, token token.Service) authport.AuthService {
	return &userService{
		repo: repo,
		token: token,
	}
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

	token, err := s.token.GenerateToken(user.ID, user.Email, 24 * time.Hour)
	// TODO: generate JWT
	return token, nil
}
