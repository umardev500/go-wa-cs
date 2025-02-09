package container

import (
	fiberws "github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/internal/adapter/websocket"
	"github.com/umardev500/chat/pkg/middleware"
	"github.com/umardev500/chat/pkg/types"
)

type wsContainer struct {
	handler websocket.WsHandler
}

func NewWsContainer() types.Container {
	handler := websocket.NewWsHandler()

	return &wsContainer{
		handler: handler,
	}
}

func (wc *wsContainer) Api(r fiber.Router) {
	ws := r.Group("/ws")

	// Middleware to upgrade the connection to WebSocket
	ws.Use("/", middleware.WsAuthMiddleware)

	// WebSocket endpoint
	ws.Get("/", fiberws.New(wc.handler.Handle))
}

func (wc *wsContainer) Web(r fiber.Router) {}
