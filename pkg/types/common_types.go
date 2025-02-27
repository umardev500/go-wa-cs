package types

import "github.com/gofiber/fiber/v2"

// Container is an interface that defines methods for setting up API and web routes
type Container interface {
	Api(r fiber.Router)
	Web(r fiber.Router)
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
