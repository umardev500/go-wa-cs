package grpc

import (
	"sync"

	"github.com/umardev500/chat/api/proto"
)

// Create a global channel to push messages to clients
var streamChan = make(chan *proto.SubscribePresenseResponse)
var mu sync.Mutex
var streamClients []proto.WhatsAppService_SubscribePresenseServer

func GetStreamChan() chan<- *proto.SubscribePresenseResponse {
	return streamChan
}

func GetStreamClients() []proto.WhatsAppService_SubscribePresenseServer {
	return streamClients
}
