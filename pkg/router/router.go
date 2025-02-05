package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/pkg/types"
)

type Router interface {
	Setup()
}

type routerStruct struct {
	app        *fiber.App
	containers []types.Container
}

func NewRouter(app *fiber.App, containers []types.Container) Router {
	return &routerStruct{
		app:        app,
		containers: containers,
	}
}

func (r *routerStruct) Setup() {
	web := r.app.Group("/")
	api := r.app.Group("/api")
	for _, container := range r.containers {
		container.Api(api)
		container.Web(web)
	}
}
