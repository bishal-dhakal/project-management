package main

import (
	"fmt"
	"net/http"

	// chathandler "github.com/bishal-dhakal/project-management/internal/handlers/chat"

	"github.com/bishal-dhakal/project-management/internal/adapter/handlers"
	authhandler "github.com/bishal-dhakal/project-management/internal/adapter/handlers/auth"
	storage "github.com/bishal-dhakal/project-management/internal/adapter/storage/postgres"
	userrepo "github.com/bishal-dhakal/project-management/internal/adapter/storage/user"
	"github.com/bishal-dhakal/project-management/internal/config"
	authservice "github.com/bishal-dhakal/project-management/internal/core/services/user"
)

func main() {
	// go chathandler.HandleMessage()
	// http.HandleFunc("/ws", chathandler.WsHandler)

	// 1. Load configuration
	cfg := config.LoadConfig()

	db := storage.NewPostgresConnection(cfg)
	fmt.Println("Loading database:", db)

	userRepo := userrepo.NewUserRepository(db)
	authSvc := authservice.New(userRepo)
	authHandler := authhandler.NewHandler(authSvc)

	mux := handlers.NewRouter(authHandler)

	fmt.Println("Server started on :" + cfg.Server_Port)
	err := http.ListenAndServe(":"+cfg.Server_Port, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Println("Hello, World!")
}
