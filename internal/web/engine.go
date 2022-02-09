package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/pivot/internal/controller"
	"github.com/shoriwe/pivot/internal/data"
	"github.com/shoriwe/pivot/internal/logs"
	"github.com/shoriwe/pivot/internal/web/handlers"
	"github.com/shoriwe/pivot/internal/web/values"
	"net/http"
)

var (
	//go:embed static/*
	staticFiles embed.FS
	//go:embed pages/*
	pagesFiles embed.FS
)

func NewEngine(connection *data.Connection, logger *logs.Logger) *gin.Engine {
	router := gin.Default()
	c := controller.NewController(connection, logger, &staticFiles, &pagesFiles)
	router.Use(gin.Logger())
	router.Use(func(context *gin.Context) {
		value, getError := context.Cookie(values.CookieName)
		if getError == nil {
			_, found := c.Connection.Cache.GetUserSession(value)
			if found {
				context.Redirect(http.StatusFound, values.DashboardLocation)
			}
		}
	})
	router.GET(values.RootLocation, func(context *gin.Context) {
		context.Redirect(http.StatusFound, values.LoginLocation)
	})
	router.GET(values.IndexLocation, func(context *gin.Context) {
		context.Redirect(http.StatusFound, values.LoginLocation)
	})
	router.GET(values.StaticLocation, func(context *gin.Context) {
		context.FileFromFS(
			context.Request.URL.Path,
			http.FS(staticFiles),
		)
	})
	router.GET(values.LoginLocation, handlers.GetLogin(c))
	router.POST(values.LoginLocation, handlers.PostLogin(c))
	router.GET(values.RegisterLocation, handlers.GetRegister(c))
	router.POST(values.RegisterLocation, handlers.PostRegister(c))
	return router
}
