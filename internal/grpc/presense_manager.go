package grpc

import (
	"github.com/umardev500/chat/api/proto"
)

// Create a global channel to push messages to clients
var presenceChan = make(chan *proto.SubscribePresenseResponse)
var presenceClients []proto.WhatsAppService_SubscribePresenseServer

func GetStreamChan() chan *proto.SubscribePresenseResponse {
	return presenceChan
}

func GetStreamClients() []proto.WhatsAppService_SubscribePresenseServer {
	return presenceClients
}
