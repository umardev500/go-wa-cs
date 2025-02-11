package domain

import "github.com/umardev500/chat/api/proto"

type ChatList struct {
	Status      string      `json:"status"`
	UnreadCount int         `json:"unreadCount"`
	LastMessage LastMessage `json:"lastMessage"`
	RemoteJID   string      `json:"remotejid"`
	CSID        string      `json:"csid"`
}

type LastMessage struct {
	Conversation string   `json:"conversation"`
	PushName     string   `json:"pushname"`
	Timestamp    int64    `json:"timestamp"`
	Metadata     Metadata `json:"metadata"`
}

type Metadata struct {
	RemoteJID string `json:"remotejid"`
	FromMe    bool   `json:"fromme"`
	ID        string `json:"id"`
}

type PushChat struct {
	IsInitial   bool                      `json:"isInitial"`
	TextMessage *proto.TextMessageRequest `json:"textMessage"`
}
