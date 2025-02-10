package configs

type Presense string

var (
	PresenseOnline        Presense = "available"
	PresenseOffline       Presense = "unavailable"
	PresenseOnlineText    Presense = "online"
	PresenseOfflineText   Presense = "offline"
	PresenseComposing     Presense = "composing"
	PresenseComposingDone Presense = "available"
)

type MessageType string

var (
	MessageTypeStatus MessageType = "status"
	MessageTypeTyping MessageType = "typing"
)
