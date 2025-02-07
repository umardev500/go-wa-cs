package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/internal/container"
	"github.com/umardev500/chat/pkg/db"
	"github.com/umardev500/chat/pkg/router"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("failed to load env")
	}
}

func start(ctx context.Context) {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	ch := make(chan error, 1)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Cache-Control",
		AllowCredentials: true,
	}))

	mongoDb := db.NewMongo()
	containers := container.NewContainerRegister(mongoDb)
	router.NewRouter(app, containers).Setup()

	go func() {
		port := os.Getenv("REST_PORT")
		addr := fmt.Sprintf(":%s", port)

		log.Info().Msgf("starting rest server on port %s", port)
		ch <- app.Listen(addr)
	}()

	select {
	case err := <-ch:
		log.Fatal().Err(err).Msg("failed to start rest server")

	case <-ctx.Done():
		log.Info().Msg("rest server stopped")
		app.Shutdown()
	}

}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	start(ctx)
}
