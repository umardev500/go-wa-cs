package container

import (
	"github.com/umardev500/chat/internal/adapter/grpc"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/pkg/db"
)

func NewWaContainer() *grpc.WaHandler {
	mongoDb := db.NewMongo()
	repo := repository.NewWaRepo(mongoDb)
	handler := grpc.NewWaHandler(repo)

	return handler
}
