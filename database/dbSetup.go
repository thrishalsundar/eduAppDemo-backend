package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSetup() (*mongo.Client, context.Context) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	uri := os.Getenv("MONGOURI")
	ctx := context.TODO()
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	fmt.Print("MongoDone")
	return client, ctx
}
