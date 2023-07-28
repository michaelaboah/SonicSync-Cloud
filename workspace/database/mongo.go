package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ClientsDB = "clients-db"
const UserCol = "users"
const EquipDB = "equipment-inventory"
const ItemsCol = "items"

func DBInstance() (*mongo.Client, error) {
  

  log.Println("[MongoDB] Starting Database Connection Instance")
  fmt.Println("[MongoDB] Starting Database Connection Instance")


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
    log.Println(err) 
  }


  log.Println("[~ Status] Connection to Mongo Database Sucessful")
  fmt.Println("[~ Status] Connection to Mongo Database Sucessful")

  log.Println("[~ Status] Starting Database Setup and Pre-Checks")
  fmt.Println("[~ Status] Starting Database Setup and Pre-Checks")
  // Setup + Pre-Check
  uniqueIndexs(client)

  


  return client, err
}


func uniqueIndexs(client *mongo.Client) {

  itemsCol := client.Database(EquipDB).Collection(ItemsCol)
  indexModel := mongo.IndexModel {
    Keys: bson.M{"model": 1},
    Options: options.Index().SetUnique(true),
  }

  indexName, err :=itemsCol.Indexes().CreateOne(context.Background(), indexModel)
  if err != nil {
    fmt.Println(err)
    log.Fatalln(err)
  }

  fmt.Println("[Pre-Check] Created Unique Index: ", indexName)
}

