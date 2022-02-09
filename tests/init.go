package tests

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/pivot/internal/data"
	"github.com/shoriwe/pivot/internal/data/memory"
	"github.com/shoriwe/pivot/internal/logs"
	"github.com/shoriwe/pivot/internal/web"
	"net"
	"net/http"
)

var (
	testHost = "127.0.0.1:8080"
)

func Serve() (string, net.Listener) {
	output := &bytes.Buffer{}
	gin.DefaultWriter = output
	gin.DefaultErrorWriter = output
	logger := logs.NewLogger(output)
	database := memory.NewMemory()
	engine := web.NewEngine(data.NewConnection(database), logger)
	listener, listenError := net.Listen("tcp", testHost)
	if listenError != nil {
		panic(listenError)
	}
	go engine.RunListener(listener)
	return fmt.Sprintf("http://%s", testHost), listener
}

func Client() *http.Client {
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return client
}
