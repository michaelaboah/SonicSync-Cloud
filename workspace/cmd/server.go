package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/michaelaboah/sonic-sync-cloud/database"
	handlers "github.com/michaelaboah/sonic-sync-cloud/handlers"
	"github.com/michaelaboah/sonic-sync-cloud/middleware"
)

const defaultPort = "8080"
const defaultLogPath = "./logs/log.log"
func main() {


  godotenv.Load("../workspace/.env")



	port := os.Getenv("PORT")
  fmt.Println(port)
	if port == "" {
		port = defaultPort
	}


  logFile, err := setupLogFile(defaultLogPath)
  if err != nil {
    log.Fatal(err)
  }



  log.SetOutput(logFile)
  log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
  log.Println("Log file created")


  mongoClient := database.DBInstance()
  

  r := gin.Default()


  

  r.Use(middleware.DbMiddleware(mongoClient))

  r.POST("/graphql", handlers.GrapqhlHandler(mongoClient))

  r.GET("/graphql-playground", handlers.PlaygroundHandler())

  r.GET("/", func(ctx *gin.Context) {

    ctx.JSON(http.StatusOK, gin.H{
      "Hello": "World",
    })

  })


  log.Fatal(r.Run(":" + port))

}




// create the required folder if necessary
func setupLogFile(path string) (*os.File, error) {
  
  logFile, err := os.OpenFile(path, os.O_WRONLY | os.O_APPEND | os.O_CREATE,  0644)


    if err != nil {
    log.Fatal(err)
    return nil, err 
  }
  
  return logFile, err
}
