package database

import (
	"context"
	"log"
	"os"
  "fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
   
  mongoURL := os.Getenv("MONGODB_URL")


  if len(mongoURL) == 0 {
    fmt.Println("MONGODB_URL not specified in .env file")
    log.Fatalln("MONGODB_URL not specified in .env file")
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
