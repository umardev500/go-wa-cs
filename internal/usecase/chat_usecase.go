package usecase

import (
	"context"

	"github.com/umardev500/chat/internal/domain"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/pkg/types"
)

type ChatUsecase interface {
	GetChatList(ctx context.Context) *types.Response
	PushChat(ctx context.Context, req *domain.PushChat)
}

type chatUsecase struct {
	repo repository.ChatRepo
}

func NewChatUsecase(repo repository.ChatRepo) ChatUsecase {
	return &chatUsecase{
		repo: repo,
	}
}

func (c *chatUsecase) PushChat(ctx context.Context, req *domain.PushChat) {}

func (c *chatUsecase) GetChatList(ctx context.Context) *types.Response {
	chats, err := c.repo.GetChatList(ctx)
	if err != nil {
		return &types.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &types.Response{
		Success: true,
		Message: "success",
		Data:    chats,
	}
}
