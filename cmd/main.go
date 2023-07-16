package main

import (
	"fmt"
	"github.com/punkestu/hello-websocket/internal/handler"
	"github.com/punkestu/hello-websocket/internal/service"
	"log"
	"net/http"
)

func main() {
	s := service.Service{}
	app := handler.Newhandler(s)
	http.HandleFunc("/", app.Index)
	http.HandleFunc("/ws", app.WebSocket)
	fmt.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err.Error())
	}
}
