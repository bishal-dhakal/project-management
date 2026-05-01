package main

import (
	"fmt"
	"net/http"

	// chathandler "github.com/bishal-dhakal/project-management/internal/handlers/chat"

	"github.com/bishal-dhakal/project-management/internal/config"
	storage "github.com/bishal-dhakal/project-management/internal/adapter/storage"
)

func main() {
	// go chathandler.HandleMessage()
	// http.HandleFunc("/ws", chathandler.WsHandler)

	// 1. Load configuration
	cfg := config.LoadConfig()

	db := storage.NewPostgresConnection(cfg)
	fmt.Println("Loading database:", db)

	fmt.Println("Websocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Println("Hello, World!")
}
