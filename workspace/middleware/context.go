package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DbMiddleware (mongoClient *mongo.Client) func (c *gin.Context) {
  return func (c *gin.Context) {
    ctx := context.WithValue(c.Request.Context(), "mongoClient", mongoClient)
    c.Request = c.Request.WithContext(ctx)
    c.Next()
  }
}
