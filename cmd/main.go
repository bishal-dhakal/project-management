package main

import (
	"fmt"
	"net/http"

	// chathandler "github.com/bishal-dhakal/project-management/internal/handlers/chat"

	"github.com/bishal-dhakal/project-management/internal/adapter/handlers"
	authhandler "github.com/bishal-dhakal/project-management/internal/adapter/handlers/auth"
	migrator "github.com/bishal-dhakal/project-management/internal/adapter/storage"
	storage "github.com/bishal-dhakal/project-management/internal/adapter/storage/postgres"
	userrepo "github.com/bishal-dhakal/project-management/internal/adapter/storage/user"
	"github.com/bishal-dhakal/project-management/internal/config"
	authservice "github.com/bishal-dhakal/project-management/internal/core/services/user"
	jwtservice "github.com/bishal-dhakal/project-management/internal/adapter/jwt"
)

func main() {
	// go chathandler.HandleMessage()
	// http.HandleFunc("/ws", chathandler.WsHandler)

	// 1. Load configuration
	cfg := config.LoadConfig()

	db := storage.NewPostgresConnection(cfg)

	err := migrator.RunMigrations(db, "./internal/adapter/storage/migrations/")
	if err != nil {
		fmt.Printf("Something went wrong :", err)
	}
	userRepo := userrepo.NewUserRepository(db)

	jwt := jwtservice.NewService("nepal123")
	authSvc := authservice.New(userRepo,jwt)
	authHandler := authhandler.NewHandler(authSvc)

	mux := handlers.NewRouter(authHandler)

	fmt.Println("Server started on :" + cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Println("Hello, World!")
}
