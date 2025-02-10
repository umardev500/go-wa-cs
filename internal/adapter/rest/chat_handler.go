package rest

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/configs"
	"github.com/umardev500/chat/internal/sse"
	"github.com/umardev500/chat/internal/usecase"
	"github.com/valyala/fasthttp"
)

type ChatHandler interface {
	GetChatList(c *fiber.Ctx) error
	Sse(c *fiber.Ctx) error
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
	time.Sleep(2 * time.Second)

	return c.JSON(resp)
}

func (ch *chatHandler) Sse(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	// TODO: get the user id
	userId := "12345"

	// Create a new channel for this client
	statusChannel := sse.AddClient(userId, configs.SSE_CHAT)

	log.Info().Msg("client connected")

	c.Status(fiber.StatusOK).
		Context().
		SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			for msg := range statusChannel {
				fmt.Fprintf(w, "data: %s\n\n", msg)
				if err := w.Flush(); err != nil {
					log.Error().Err(err).Msg("failed to flush sse")
					close(statusChannel)
					break
				}
			}
		}))

	return nil
}
