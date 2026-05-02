package handlers

import (
    "net/http"
    authhandler "github.com/bishal-dhakal/project-management/internal/adapter/handlers/auth"
)

func NewRouter(
	authHandler *authhandler.Handler,
) http.Handler{
	mux := http.NewServeMux()
	
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	return mux
}