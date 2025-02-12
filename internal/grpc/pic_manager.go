package grpc

import (
	"sync"

	"github.com/umardev500/chat/api/proto"
)

type PicClient struct {
	Stream     proto.WhatsAppService_SubscribeProfilePicServer
	MsgChan    chan *proto.SubscribeProfilePicResponse
	ResultChan chan string
}

var (
	picClient *PicClient
	picMutex  sync.RWMutex
)

func AddPicClient(client *PicClient) {
	picMutex.Lock()
	defer picMutex.Unlock()
	picClient = client
}

func RemovePicClient(client *PicClient) {
	picMutex.Lock()
	defer picMutex.Unlock()

	picClient = nil
}

func GetPicClient() *PicClient {
	picMutex.Lock()
	defer picMutex.Unlock()

	return picClient
}
