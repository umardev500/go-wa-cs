package container

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/api/proto"
	"github.com/umardev500/chat/configs"
	"github.com/umardev500/chat/internal/adapter/grpc"
	"github.com/umardev500/chat/internal/adapter/rest"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/internal/sse"
	"github.com/umardev500/chat/internal/usecase"
	"github.com/umardev500/chat/pkg/db"
	"github.com/umardev500/chat/pkg/types"
	"github.com/umardev500/chat/pkg/utils"
)

type chatContainer struct {
	handler rest.ChatHandler
}

func NewChatContainer(mongoDb *db.Mongo) types.Container {
	repo := repository.NewChatRepo(mongoDb)
	uc := usecase.NewChatUsecase(repo)
	handler := rest.NewChatHandler(uc)

	return &chatContainer{
		handler: handler,
	}
}

var (
	currentStatus = "online"
)

func (c *chatContainer) Api(r fiber.Router) {

	chat := r.Group("/chat")
	chat.Get("/", c.handler.GetChatList)
	chat.Get("/sse/:id", c.handler.Sse)
	chat.Get("/test", func(c *fiber.Ctx) error {
		userId := "12345"
		channel := sse.GetSseResource(userId, configs.SSE_CHAT)
		if channel == nil {
			return c.JSON("nil")
		}

		if currentStatus == "online" {
			currentStatus = "offline"
		} else {
			currentStatus = "online"
		}

		// Define the structured JSON message
		message := map[string]interface{}{
			"type": "status",
			"data": []map[string]string{
				{
					"status":    currentStatus,
					"remotejid": "6285123456791@s.whatsapp.net",
				},
				{
					"status":    currentStatus,
					"remotejid": "6285123456781@s.whatsapp.net",
				},
			},
		}

		// Convert the Go struct to a JSON string
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to encode JSON"})
		}

		channel <- string(jsonMessage)

		return c.JSON(currentStatus)
	})
	chat.Get("/ws/test", func(c *fiber.Ctx) error {
		if currentStatus == "online" {
			currentStatus = "offline"
		} else {
			currentStatus = "online"
		}

		for _, conn := range utils.WsGetClients() {
			var data = map[string]interface{}{
				"mt": "status",
				"data": []map[string]interface{}{
					{
						"status":    currentStatus,
						"remotejid": "6285123456791@s.whatsapp.net",
					},
					{
						"status":    currentStatus,
						"remotejid": "6285123456781@s.whatsapp.net",
					},
				},
			}
			jsonByte, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				return err
			}

			fmt.Println("send", string(jsonByte))

			conn.WriteJSON(data)
		}

		return nil
	})
	chat.Get("/stream", func(c *fiber.Ctx) error {
		grpc.GetStreamChan() <- &proto.SubscribePresenseResponse{
			Mt:  "status",
			Jid: []string{"6283142765573@s.whatsapp.net"},
		}

		return c.JSON("ok")
	})
	chat.Get("/ping", func(c *fiber.Ctx) error {
		fmt.Println(grpc.GetStreamClients())
		return c.JSON("pong")
	})

	// update unread counter
	chat.Get("/update-unread-count", c.handler.UpdateUnreadCounter)
}

func (c *chatContainer) Web(r fiber.Router) {}
