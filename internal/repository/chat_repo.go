package repository

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/internal/domain"
	"github.com/umardev500/chat/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ChatRepo interface {
	InitializeChat(remoteJid, csId string) (bool, error)
	GetChatList(ctx context.Context) ([]domain.ChatList, error)
	CheckExist(ctx context.Context, remoteJid string) (bool, error)
	CheckExistByCsIdAndRemoteJid(ctx context.Context, userId, remoteJid string) (bool, error)
	PushMessge(remoteJid, csid string, message interface{}) error
}

type chatRepo struct {
	mongoDb *db.Mongo
}

func NewChatRepo(mongoDb *db.Mongo) ChatRepo {
	return &chatRepo{
		mongoDb: mongoDb,
	}
}

func (r *chatRepo) InitializeChat(remoteJid, csId string) (bool, error) {
	// Get the collection
	coll := r.mongoDb.Db.Collection("messages")

	// Create a filter with both remotejid and customer_service_jid
	filter := bson.D{
		{Key: "remotejid", Value: remoteJid},
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
				{Key: "unread_count", Value: 0},
				{Key: "messages", Value: []string{}}, // Empty message array
			}

			_, err := coll.InsertOne(context.Background(), newDocument)
			if err != nil {
				log.Error().Err(err).Msg("failed to insert new chat")
			}

			log.Info().Msgf("New chat created for remotejid: %s", remoteJid)
			return true, nil
		}
		return false, nil
	}

	log.Info().Msgf("Chat found for remotejid: %s", remoteJid)

	return false, nil
}

func (c *chatRepo) CheckExist(ctx context.Context, remoteJid string) (bool, error) {
	coll := c.mongoDb.Db.Collection("messages")
	filter := bson.D{{Key: "remotejid", Value: remoteJid}}
	var result bson.M
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c *chatRepo) CheckExistByCsIdAndRemoteJid(ctx context.Context, csId, remoteJid string) (bool, error) {
	coll := c.mongoDb.Db.Collection("messages")
	filter := bson.D{
		{Key: "remotejid", Value: remoteJid},
		{Key: "csid", Value: csId},
	}
	var result bson.M
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c *chatRepo) GetChatList(ctx context.Context) ([]domain.ChatList, error) {
	coll := c.mongoDb.Db.Collection("messages")

	aggregationPipeline := mongo.Pipeline{
		bson.D{
			{Key: "$unwind", Value: "$messages"},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "messages.timestamp", Value: -1},
			}},
		},
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "remotejid", Value: "$remotejid"},
					{Key: "csid", Value: "$csid"},
				}},
				{Key: "status", Value: bson.D{
					{Key: "$last", Value: "$status"},
				}},
				{Key: "unreadCount", Value: bson.D{
					{Key: "$last", Value: "$unreadCount"},
				}},
				{Key: "lastMessage", Value: bson.D{
					{Key: "$last", Value: "$messages"},
				}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "remotejid", Value: "$_id.remotejid"},
				{Key: "csid", Value: "$_id.csid"},
				{Key: "status", Value: 1},
				{Key: "unreadCount", Value: 1},
				{Key: "lastMessage", Value: 1},
			}},
		},
		// ✅ FINAL SORT: Ensure list is sorted by latest message timestamp
		bson.D{{Key: "$sort", Value: bson.D{
			{Key: "lastMessage.timestamp", Value: -1},
		}}},
	}

	cur, err := coll.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var chatList []domain.ChatList
	if err := cur.All(ctx, &chatList); err != nil {
		return nil, err
	}

	return chatList, nil
}

func (r *chatRepo) PushMessge(remoteJid, csid string, message interface{}) error {
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

	log.Info().Msgf("Message added to chat for remotejid: %s and csid: %s message: %v", remoteJid, csid, message)
	return nil
}
