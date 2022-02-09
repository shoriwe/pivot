package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/pivot/internal/data"
	"github.com/shoriwe/pivot/internal/data/memory"
	"github.com/shoriwe/pivot/internal/logs"
	"github.com/shoriwe/pivot/internal/web"
	"log"
	"os"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	database := memory.NewMemory()
	connection := data.NewConnection(database)
	logger := logs.NewLogger(os.Stderr)
	engine := web.NewEngine(connection, logger)
	executionError := engine.Run(":8080")
	if executionError != nil {
		log.Fatal(executionError)
	}
	fmt.Println("Everything is Fine :)")
}
