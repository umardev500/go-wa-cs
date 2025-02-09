package websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/pkg/constants"
	"github.com/umardev500/chat/pkg/utils"
)

type WsHandler interface {
	Handle(c *websocket.Conn)
}

type wsHandler struct{}

func NewWsHandler() WsHandler {
	return &wsHandler{}
}

func (h *wsHandler) Handle(c *websocket.Conn) {
	userId := c.Locals(constants.KeyUserId)
	utils.WsAddClient(userId.(string), c)

	var (
		// mt  int
		msg []byte
		err error
	)

	for {
		_, msg, err = c.ReadMessage()
		if err != nil {
			log.Error().Err(err).Msg("failed to read message")
			break
		}

		log.Info().Msgf("Received message: %s", string(msg))
	}
}
