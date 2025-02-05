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
	FindActiveChat(remoteJid string) (string, error)
	InitializeChat(remoteJid, csId string) error
	SaveMessage(data interface{}) error
	PushMessge(remoteJid, csid string, message interface{}) error
}

type waRepo struct {
	mongoDb *db.Mongo
}

func NewWaRepo(db *db.Mongo) WaRepo {
	return &waRepo{
		mongoDb: db,
	}
}

func (r *waRepo) FindActiveChat(remoteJid string) (string, error) {
	// Get the collection
	coll := r.mongoDb.Db.Collection("messages")

	// Create a filter to find a document with the given remotejid and status "active"
	filter := bson.D{
		{Key: "remotejid", Value: remoteJid},
		{Key: "status", Value: "active"},
	}

	// Define a variable to hold the result
	var result bson.M

	// Query the collection for one matching document
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info().Msgf("No active chat found for remotejid: %s", remoteJid)
			return "", nil // Return nil if no document is found
		}
		log.Error().Err(err).Msg("failed to find an active chat")
		return "", err
	}

	// Return the csid
	csid, ok := result["csid"].(string)
	if !ok {
		log.Error().Msg("csid is not a string")
		return "", fmt.Errorf("csid is not a string")
	}

	return csid, nil
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
				{Key: "status", Value: "queueing"},
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

func (r *waRepo) PushMessge(remoteJid, csid string, message interface{}) error {
	// Get the collection
	coll := r.mongoDb.Db.Collection("messages")

	// Create a filter with both remotejid and customer_service_jid
	filter := bson.D{
		{Key: "remotejid", Value: remoteJid},
		{Key: "csid", Value: csid},
	}

	// Define the update operation to push the message into the "messages" array
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "messages", Value: message},
		}},
	}

	// Perform the update
	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to push message: %w", err)
	}

	// Check if any document was updated
	if result.MatchedCount == 0 {
		log.Info().Msgf("No chat found for remotejid: %s and csid: %s, message not added", remoteJid, csid)
		return nil
	}

	log.Info().Msgf("Message added to chat for remotejid: %s and csid: %s", remoteJid, csid)
	return nil
}
