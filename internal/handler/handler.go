package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/punkestu/hello-websocket/internal/entity"
	"github.com/punkestu/hello-websocket/internal/service"
	"log"
	"net/http"
	"os"
)

type Handler struct {
	connections []*entity.WebSocketConnection
	service     *service.Service
}

func Newhandler(_service service.Service) *Handler {
	h := &Handler{connections: []*entity.WebSocketConnection{}, service: &_service}
	h.service.SetConnections(&h.connections)
	return h
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("pages/index.html")
	if err != nil {
		http.Error(w, "Could not open requested file", http.StatusInternalServerError)
		return
	}
	if _, err = fmt.Fprintf(w, "%s", content); err != nil {
		log.Fatalln(err.Error())
	}
}

func (h *Handler) WebSocket(w http.ResponseWriter, r *http.Request) {
	currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	username := r.URL.Query().Get("username")
	currentConn := entity.WebSocketConnection{Conn: currentGorillaConn, Username: username}
	h.connections = append(h.connections, &currentConn)

	go h.service.HandleIO(&currentConn)
}
