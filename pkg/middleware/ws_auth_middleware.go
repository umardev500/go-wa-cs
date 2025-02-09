package middleware

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/pkg/constants"
)

func WsAuthMiddleware(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		// Get token from query parameters
		tokenString := c.Query("token")
		if tokenString == "" {
			return fiber.ErrUnauthorized
		}

		// TODO: implement token validation

		// Token is valid, allow WebSocket upgrade
		userId := tokenString // TODO: get the user id from the token
		c.Locals(constants.KeyUserId, userId)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}
