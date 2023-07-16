package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/michaelaboah/sonic-sync-cloud/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  r := gin.Default()

  r.POST("/graphql", grapqhlHandler())

  r.GET("/graphql-playground", playgroundHandler())

  log.Fatal(r.Run(":" + port))

}

func grapqhlHandler() gin.HandlerFunc  {
    
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

  return func(ctx *gin.Context) {
    h.ServeHTTP(ctx.Writer, ctx.Request)
  }

}

func playgroundHandler() gin.HandlerFunc  {
 
  h := playground.Handler("GraphQL playground", "/query")

  return func(ctx *gin.Context) {
    
    h.ServeHTTP(ctx.Writer, ctx.Request)
    
  }
}



