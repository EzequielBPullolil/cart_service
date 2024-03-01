package dbmanager

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConnection struct {
	CartCollection *mongo.Collection
}

func ConnectDB(db_uri, db_name string) *DbConnection {
	var DbConnection DbConnection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db_uri))

	if err != nil {
		panic(err)
	}
	DbConnection.CartCollection = client.Database(db_name).Collection("cart")

	return &DbConnection
}
