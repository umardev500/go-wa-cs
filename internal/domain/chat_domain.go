package domain

type ChatList struct {
	Status           string      `json:"status"`
	LastMessage      LastMessage `json:"lastMessage"`
	MessageTimestamp *int64      `json:"messageTimestamp"`
	RemoteJID        string      `json:"remotejid"`
	CSID             string      `json:"csid"`
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
