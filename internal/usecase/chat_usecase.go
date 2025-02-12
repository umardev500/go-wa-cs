package usecase

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/api/proto"
	"github.com/umardev500/chat/configs"
	"github.com/umardev500/chat/internal/domain"
	grpcManager "github.com/umardev500/chat/internal/grpc"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/pkg/types"
	"github.com/umardev500/chat/pkg/utils"
)

type ChatUsecase interface {
	GetChatList(ctx context.Context, csid string) *types.Response
	UpdateUnreadCounter(ctx context.Context, csId string, jid string) types.Response
	PushChat(ctx context.Context, csId string, req *domain.PushChat) error
	GetProfilePic(ctx context.Context, jid string) *types.Response
}

type chatUsecase struct {
	repo repository.ChatRepo
}

func NewChatUsecase(repo repository.ChatRepo) ChatUsecase {
	return &chatUsecase{
		repo: repo,
	}
}

func (c *chatUsecase) GetProfilePic(ctx context.Context, jid string) *types.Response {
	client := grpcManager.GetPicClient()

	var url string

	client.MsgChan <- &proto.SubscribeProfilePicResponse{
		Jid: jid,
	}

	url = <-client.ResultChan

	return &types.Response{
		Success: true,
		Message: "Subscribed to profile pic",
		Data:    url,
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
	var messagePayload interface{}

	if req.Data.TextMessage != nil {
		jid = req.Data.TextMessage.Metadata.RemoteJid
		messagePayload = req.Data.TextMessage
	} else {
		log.Error().Msg("invalid message type")
		return fmt.Errorf("invalid message type")
	}

	isInitial, err := c.repo.InitializeChat(jid, csId)
	if err != nil {
		log.Err(err).Msg("failed to initialize chat")
		return err
	}

	// Push message
	err = c.repo.PushMessge(jid, csId, messagePayload)
	if err != nil {
		log.Err(err).Msg("failed to push message")
		return err
	}

	if isInitial {
		// TODO: Fetch chat list for initialize chat
		req.Data.IsInitial = isInitial

		chats, err := c.getChatList(ctx, &jid, csId)
		if err != nil {
			log.Err(err).Msg("failed to fetch chat list")
			return err
		}

		req.Data.InitialChats = chats
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

func (c *chatUsecase) getChatList(ctx context.Context, remoteJid *string, csid string) ([]domain.ChatList, error) {
	chats, err := c.repo.GetChatList(ctx, remoteJid, csid)
	if err != nil {
		return nil, err
	}

	// Subscribe chats to presense channel
	var jids []string
	for _, chat := range chats {
		jid := chat.RemoteJID
		jids = append(jids, jid)
	}

	for _, client := range grpcManager.GetPresenceClients() {
		client.MsgChan <- &proto.SubscribePresenseResponse{
			Mt:  string(configs.MessageTypeStatus),
			Jid: jids,
		}
	}

	return chats, nil
}

func (c *chatUsecase) GetChatList(ctx context.Context, csid string) *types.Response {
	chats, err := c.getChatList(ctx, nil, csid)
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

func (c *chatUsecase) UpdateUnreadCounter(ctx context.Context, csId string, jid string) types.Response {
	updated, err := c.repo.UpdateUnreadCounter(ctx, csId, jid, 1)
	if err != nil {
		return types.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return types.Response{
		Success: updated,
		Message: "Update unread counter",
		Data:    nil,
	}
}
