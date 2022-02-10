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
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s IP:PORT", os.Args[0])
		return
	}
	database := memory.NewMemory()
	connection := data.NewConnection(database)
	logger := logs.NewLogger(os.Stderr)
	engine := web.NewEngine(connection, logger)
	executionError := engine.Run(os.Args[1])
	if executionError != nil {
		log.Fatal(executionError)
	}
	fmt.Println("Everything is Fine :)")
}
