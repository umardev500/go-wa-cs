package repository

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type WaRepo interface {
	InitializeChat(remoteJid, csId string) error
	SaveMessage(data interface{}) error
}

type waRepo struct {
	mongoDb *db.Mongo
}

func NewWaRepo(db *db.Mongo) WaRepo {
	return &waRepo{
		mongoDb: db,
	}
}

func (r *waRepo) InitializeChat(remoteJid, csId string) error {
	// Get the collection
	coll := r.mongoDb.Db.Collection("messages")

	// Create a filter with both remotejid and customer_service_jid
	filter := bson.D{
		{Key: "remotejid", Value: remoteJid},
		{Key: "csid", Value: csId},
	}

	// Find the document
	var result bson.M
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newDocument := bson.D{
				{Key: "remotejid", Value: remoteJid},
				{Key: "csid", Value: csId},
				{Key: "messages", Value: []string{}}, // Empty message array
			}

			_, err := coll.InsertOne(context.Background(), newDocument)
			if err != nil {
				log.Error().Err(err).Msg("failed to insert new chat")
			}

			log.Info().Msgf("New chat created for remotejid: %s", remoteJid)
		}
		return err
	}

	log.Info().Msgf("Chat found for remotejid: %s", remoteJid)

	return nil
}

func (r *waRepo) SaveMessage(data interface{}) error {
	bsonData, err := bson.Marshal(data)
	if err != nil {
		return err
	}

	coll := r.mongoDb.Db.Collection("messages")
	_, err = coll.InsertOne(context.Background(), bsonData)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}
