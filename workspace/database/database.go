package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {

  mongoURL := os.Getenv("MONGODB_URL")
  fmt.Println(mongoURL)


  if mongoURL == "" {
    log.Fatal("MONGODB_URL not specified in .env file")
  }

  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)

  client, err := mongo.Connect(context.Background(), opts)

  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(context.Background(), nil)
  
  if err != nil {
    log.Fatal(err) 
  }

  return client
}
