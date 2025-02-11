package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/configs"
	"github.com/umardev500/chat/internal/domain"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/pkg/types"
	"github.com/umardev500/chat/pkg/utils"
)

type ChatUsecase interface {
	GetChatList(ctx context.Context) *types.Response
	PushChat(ctx context.Context, csId string, req *domain.PushChat) error
}

type chatUsecase struct {
	repo repository.ChatRepo
}

func NewChatUsecase(repo repository.ChatRepo) ChatUsecase {
	return &chatUsecase{
		repo: repo,
	}
}

func (c *chatUsecase) broadcastChat(req *domain.PushChat, csId string) {
	conn := utils.WsGetClient(csId)
	if conn == nil {
		log.Info().Msgf("connection not found: %s", csId)
		return
	}
	if req.Data.TextMessage != nil {
		req.Mt = string(configs.MessageTypeMessage)
		conn.WriteJSON(req)
	}

}

func (c *chatUsecase) PushChat(ctx context.Context, csId string, req *domain.PushChat) error {
	var jid string

	if req.Data.TextMessage != nil {
		jid = req.Data.TextMessage.Metadata.RemoteJid
	}

	isInitial, err := c.repo.InitializeChat(jid, csId)
	if err != nil {
		log.Err(err).Msg("failed to initialize chat")
		return err
	}

	if isInitial {
		req.Data.IsInitial = isInitial
	}

	exist, err := c.repo.CheckExistByCsIdAndRemoteJid(ctx, csId, jid)
	if err != nil {
		log.Err(err).Msg("failed to check exist")
		return err
	}

	if exist {
		c.broadcastChat(req, csId)
	}

	return nil
}

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
