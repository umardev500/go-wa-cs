package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/internal/adapter/rest"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/internal/usecase"
	"github.com/umardev500/chat/pkg/db"
	"github.com/umardev500/chat/pkg/types"
)

type chatContainer struct {
	handler rest.ChatHandler
}

func NewChatContainer(mongoDb *db.Mongo) types.Container {
	repo := repository.NewChatRepo(mongoDb)
	uc := usecase.NewChatUsecase(repo)
	handler := rest.NewChatHandler(uc)

	return &chatContainer{
		handler: handler,
	}
}

func (c *chatContainer) Api(r fiber.Router) {
	chat := r.Group("/chat")
	chat.Get("/", c.handler.GetChatList)
	chat.Get("/sse/:id", c.handler.Sse)
}

func (c *chatContainer) Web(r fiber.Router) {}
