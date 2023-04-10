package mdb

import (
	"context"

	"github.com/idprm/go-telenor-pk/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(cfg *config.Secret) *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Mdb.Uri))
	if err != nil {
		panic(err)
	}
	return client.Database(cfg.Mdb.Name).Collection(cfg.Mdb.Collection)
}
