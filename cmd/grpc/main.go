package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/api/proto"
	"github.com/umardev500/chat/internal/adapter/grpc"
	googleRpc "google.golang.org/grpc"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("failed to load env")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":"+os.Getenv("RPC_PORT"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	grpcServer := googleRpc.NewServer()
	proto.RegisterWhatsAppServiceServer(grpcServer, grpc.NewWaHandler())

	log.Info().Msgf("starting grpc server on port %s", os.Getenv("RPC_PORT"))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
