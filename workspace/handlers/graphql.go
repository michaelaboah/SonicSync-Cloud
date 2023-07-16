package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/michaelaboah/sonic-sync-cloud/graph"
)

func GrapqhlHandler() gin.HandlerFunc  {
    
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

  return func(ctx *gin.Context) {
    h.ServeHTTP(ctx.Writer, ctx.Request)
  }

}

func PlaygroundHandler() gin.HandlerFunc  {
 
  h := playground.Handler("GraphQL playground", "/query")

  return func(ctx *gin.Context) {
    
    h.ServeHTTP(ctx.Writer, ctx.Request)
    
  }

}
