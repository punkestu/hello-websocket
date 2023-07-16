package service

import (
	"fmt"
	"github.com/punkestu/hello-websocket/internal/entity"
	"github.com/punkestu/hello-websocket/internal/lib"
	"log"
	"strings"
)

type Service struct {
	connections *[]*entity.WebSocketConnection
}

func (s *Service) SetConnections(connections *[]*entity.WebSocketConnection) {
	s.connections = connections
}

func (s *Service) HandleIO(currentConn *entity.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	lib.BroadcastMessage(currentConn, lib.MESSAGE_NEW_USER, "", *s.connections)

	for {
		payload := entity.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				lib.BroadcastMessage(currentConn, lib.MESSAGE_LEAVE, "", *s.connections)
				lib.EjectConnection(currentConn, *s.connections)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		lib.BroadcastMessage(currentConn, lib.MESSAGE_CHAT, payload.Message, *s.connections)
	}
}
