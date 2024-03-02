package e2e

import (
	"context"
	"os"

	"github.com/EzequielBPullolil/cart_service/src"
	"github.com/EzequielBPullolil/cart_service/src/cart"
	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"go.mongodb.org/mongo-driver/bson"
)

var app = src.CreateApp()

func init() {
	dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection.DeleteMany(context.Background(), bson.M{})
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type DataResponse struct {
	Status string    `json:"status"`
	Data   cart.Cart `json:"data"`
}
