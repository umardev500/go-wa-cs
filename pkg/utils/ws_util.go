package utils

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog/log"
)

var (
	wsClients    = make(map[string]*websocket.Conn) // userid:channel
	wsClientLock sync.Mutex
)

func WsAddClient(userId string, conn *websocket.Conn) {
	wsClientLock.Lock()
	defer wsClientLock.Unlock()

	if _, ok := wsClients[userId]; !ok {
		wsClients[userId] = conn
	} else {
		log.Info().Msgf("Connection already exists: %s and will be overriden", userId)
	}

	wsClients[userId] = conn
}

func WsGetClient(userId string) *websocket.Conn {
	wsClientLock.Lock()
	defer wsClientLock.Unlock()

	return wsClients[userId]
}

func WsGetClients() map[string]*websocket.Conn {
	wsClientLock.Lock()
	defer wsClientLock.Unlock()

	return wsClients
}

func WsRemoveClient(userId string) {
	wsClientLock.Lock()
	defer wsClientLock.Unlock()

	delete(wsClients, userId)
}
