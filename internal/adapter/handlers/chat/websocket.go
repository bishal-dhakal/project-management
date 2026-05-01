package chat

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"fmt"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan []byte)			 // Broadcast channel
var mutex = &sync.Mutex{}					 // Protect clients map



func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}

	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil{
			mutex.Lock()
			delete(clients,conn)
			mutex.Unlock()
			break
		}
		broadcast <- message
	}
}

func HandleMessage(){
	for {
		message := <- broadcast

		mutex.Lock()

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message); 
			if err != nil {
			client.Close()
			delete(clients, client)
		}
		}
		mutex.Unlock()
	}
}