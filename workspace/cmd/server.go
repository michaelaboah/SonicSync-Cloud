package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
  handlers "github.com/michaelaboah/sonic-sync-cloud/handlers"
)

const defaultPort = "8080"
const defaultLogPath = "./logs/log.log"
func main() {
	port := os.Getenv("PORT")
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

  r := gin.Default()

  r.POST("/graphql", handlers.GrapqhlHandler())

  r.GET("/graphql-playground", handlers.PlaygroundHandler())

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
