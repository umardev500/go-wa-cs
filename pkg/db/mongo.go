package db

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	db *mongo.Database
}

func NewMongo() *Mongo {
	log.Info().Msg("connecting to mongodb")
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal().Msg("MONGODB_URI env is missing")
	}

	// Set connection options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongodb")
	}

	// Ping
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping mongodb")
	}

	database := client.Database(os.Getenv("MONGODB_DATABASE"))
	mongo := &Mongo{
		db: database,
	}

	log.Info().Msg("connected to mongodb")
	return mongo
}
