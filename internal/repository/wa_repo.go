package repository

import (
	"context"
	"fmt"

	"github.com/umardev500/chat/pkg/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type WaRepo interface {
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
