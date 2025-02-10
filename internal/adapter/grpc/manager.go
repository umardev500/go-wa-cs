package grpc

import (
	"sync"

	"github.com/umardev500/chat/api/proto"
)

// Create a global channel to push messages to clients
var streamChan = make(chan *proto.StreamMessageResponse)
var mu sync.Mutex
var streamClients []proto.WhatsAppService_StreamMessageServer

func GetStreamChan() chan<- *proto.StreamMessageResponse {
	return streamChan
}

func GetStreamClients() []proto.WhatsAppService_StreamMessageServer {
	return streamClients
}
