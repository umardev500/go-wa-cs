package rest

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/internal/usecase"
)

type ChatHandler interface {
	GetChatList(c *fiber.Ctx) error
}

type chatHandler struct {
	uc usecase.ChatUsecase
}

func NewChatHandler(uc usecase.ChatUsecase) ChatHandler {
	return &chatHandler{
		uc: uc,
	}
}

func (ch *chatHandler) GetChatList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp := ch.uc.GetChatList(ctx)

	return c.JSON(resp)
}
