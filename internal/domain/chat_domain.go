package domain

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

type ContextInfo struct {
	StanzaID    string `json:"stanzaId"`
	Participant string `json:"participant"`
	// QuotedMessage  QuotedMessage `json:"quotedMessage"`
}

type MessageText struct {
	Conversation string `json:"conversation"`
}

type MessageTextExtended struct {
	Text        string       `json:"text"`
	ContextInfo *ContextInfo `json:"contextInfo"`
}

type Message struct {
	MessageText *MessageText `json:"messageText"`
	PushName    string       `json:"pushname"`
	Timestamp   int64        `json:"timestamp"`
	Metadata    Metadata     `json:"metadata"`
}

type PushChat struct {
	IsInitial bool `json:"isInitial"`
}
