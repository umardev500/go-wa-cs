package configs

type Presense string

var (
	PresenseOnline      Presense = "available"
	PresenseOffline     Presense = "unavailable"
	PresenseOnlineText  Presense = "online"
	PresenseOfflineText Presense = "offline"
)
