package grpc

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/api/proto"
)

type PresenceClient struct {
	Stream  proto.WhatsAppService_SubscribePresenseServer
	MsgChan chan *proto.SubscribePresenseResponse
}

var (
	mu             sync.Mutex
	presenceClient []*PresenceClient
)

func AddPresenceClient(client *PresenceClient) {
	mu.Lock()
	defer mu.Unlock()
	presenceClient = append(presenceClient, client)
}

func RemovePresenceClient(client *PresenceClient) {
	mu.Lock()
	defer mu.Unlock()

	for i, c := range presenceClient {
		if c == client {
			presenceClient = append(presenceClient[:i], presenceClient[i+1:]...)
			break
		}
	}

	log.Info().Msgf("removed presence client: %v", presenceClient)
}

func GetPresenceClients() []*PresenceClient {
	mu.Lock()
	defer mu.Unlock()

	return presenceClient
}
