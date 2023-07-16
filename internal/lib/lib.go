package lib

import (
	"github.com/punkestu/hello-websocket/internal/entity"
	"log"
)

const (
	MESSAGE_NEW_USER = "New User"
	MESSAGE_CHAT     = "Chat"
	MESSAGE_LEAVE    = "Leave"
)

func EjectConnection(currentConn *entity.WebSocketConnection, connections []*entity.WebSocketConnection) {
	for i, connection := range connections {
		if connection.Conn == currentConn.Conn {
			connections = append(connections[:i], connections[i+1:]...)
		}
	}
}
func BroadcastMessage(currentConn *entity.WebSocketConnection, kind, message string, connections []*entity.WebSocketConnection) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		if err := eachConn.WriteJSON(entity.SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		}); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
