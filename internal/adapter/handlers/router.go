package handlers

import (
    "net/http"
    authhandler "github.com/bishal-dhakal/project-management/internal/adapter/handlers/auth"
)

func NewRouter(
	authHandler *authhandler.Handler,
) http.Handler{
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	return mux
}