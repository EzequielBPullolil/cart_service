package tests

import (
	"context"
	"os"

	"github.com/EzequielBPullolil/cart_service/src"
	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"go.mongodb.org/mongo-driver/bson"
)

var app = src.CreateApp()

func init() {
	dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection.DeleteMany(context.Background(), bson.M{})
}
