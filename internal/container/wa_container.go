package container

import (
	"github.com/umardev500/chat/internal/adapter/grpc"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/internal/usecase"
	"github.com/umardev500/chat/pkg/db"
)

func NewWaContainer() *grpc.WaHandler {
	mongoDb := db.NewMongo()
	repo := repository.NewWaRepo(mongoDb)
	chatRepo := repository.NewChatRepo(mongoDb)
	chatUc := usecase.NewChatUsecase(chatRepo)
	handler := grpc.NewWaHandler(repo, chatUc)

	return handler
}
