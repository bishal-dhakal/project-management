package userservice

import (
	"context"
	userport "github.com/bishal-dhakal/project-management/internal/core/port/user"
)

type userService struct {
	repo userport.UserRepository
}

func New(repo userport.UserRepository) userport.AuthService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, email, password string) (string, error) {
    // TODO: validate input
    // TODO: check email exists
    // TODO: hash password
    // TODO: save user
    return "nepal", nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
    // TODO: find user by email
    // TODO: compare password
    // TODO: generate JWT
    return "", nil
}