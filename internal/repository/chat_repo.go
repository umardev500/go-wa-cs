package repository

import (
	"context"

	"github.com/umardev500/chat/internal/domain"
	"github.com/umardev500/chat/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ChatRepo interface {
	GetChatList(ctx context.Context) ([]domain.ChatList, error)
}

type chatRepo struct {
	mongoDb *db.Mongo
}

func NewChatRepo(mongoDb *db.Mongo) ChatRepo {
	return &chatRepo{
		mongoDb: mongoDb,
	}
}

func (c *chatRepo) GetChatList(ctx context.Context) ([]domain.ChatList, error) {
	coll := c.mongoDb.Db.Collection("messages")

	aggregationPipeline := mongo.Pipeline{
		bson.D{
			{Key: "$unwind", Value: "$messages"},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "messages.metadata.timestamp", Value: -1},
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
				{Key: "lastMessage", Value: bson.D{
					{Key: "$last", Value: "$messages"},
				}},
				{Key: "messageTimestamp", Value: bson.D{
					{Key: "$last", Value: "$messages.metadata.timestamp"},
				}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "remotejid", Value: "$_id.remotejid"},
				{Key: "csid", Value: "$_id.csid"},
				{Key: "status", Value: 1},
				{Key: "lastMessage", Value: 1},
				{Key: "messageTimestamp", Value: 1},
			}},
		},
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
