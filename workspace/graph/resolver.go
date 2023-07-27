//go:generate go run generate.go
package graph

import (
	"github.com/michaelaboah/sonic-sync-cloud/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
  users []*model.User
  items []*model.Item
  DB *mongo.Client
}
