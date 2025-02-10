package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/soheilhy/cmux"
	"github.com/umardev500/chat/api/proto"
	"github.com/umardev500/chat/internal/container"
	"github.com/umardev500/chat/pkg/db"
	"github.com/umardev500/chat/pkg/router"
	googleRpc "google.golang.org/grpc"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("failed to load env")
	}
}

func startRest(ctx context.Context, lis net.Listener) {
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
		log.Info().Msgf("Rest started")
		ch <- app.Listener(lis)
	}()

	select {
	case err := <-ch:
		log.Fatal().Err(err).Msg("failed to start rest server")

	case <-ctx.Done():
		log.Info().Msg("rest server stopped")
		app.Shutdown()
	}

	fmt.Println("dopne")

}

func startGrpc(lis net.Listener) {
	grpcServer := googleRpc.NewServer()
	proto.RegisterWhatsAppServiceServer(grpcServer, container.NewWaContainer())

	log.Info().Msgf("Grpc started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	m := cmux.New(lis)

	// Match gRPC and HTTP connections
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	// Run gRPC & Fiber REST servers concurrently
	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
		defer cancel()
		startRest(ctx, httpL)
	}()

	go func() {
		startGrpc(grpcL)
	}()

	log.Info().Msg("🚀 Running gRPC & REST API on the same port: " + port)
	if err := m.Serve(); err != nil {
		log.Error().Err(err).Msg("❌ cmux Error")
	}
}
