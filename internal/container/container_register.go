package container

import (
	"github.com/umardev500/chat/pkg/db"
	"github.com/umardev500/chat/pkg/types"
)

func NewContainerRegister(mongoDb *db.Mongo) []types.Container {
	return []types.Container{
		NewChatContainer(mongoDb),
		NewWsContainer(),
	}
}
