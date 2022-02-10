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

func checkLogin(c *controller.Controller) gin.HandlerFunc {
	return func(context *gin.Context) {
		value, getError := context.Cookie(values.CookieName)
		if getError == nil {
			_, found := c.Connection.Cache.GetUserSession(value)
			if found {
				context.Redirect(http.StatusFound, values.DashboardLocation)
			}
		}
	}
}

func requiresLogin(c *controller.Controller) gin.HandlerFunc {
	return func(context *gin.Context) {
		value, getError := context.Cookie(values.CookieName)
		if getError == nil {
			_, found := c.Connection.Cache.GetUserSession(value)
			if found {
				return
			}
		}
		context.Redirect(http.StatusFound, values.LoginLocation)
	}
}

func NewEngine(connection *data.Connection, logger *logs.Logger) *gin.Engine {
	router := gin.Default()
	c := controller.NewController(connection, logger, &staticFiles, &pagesFiles)
	router.Use(gin.Logger())
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
	router.GET(values.LoginLocation, checkLogin(c), handlers.GetLogin(c))
	router.POST(values.LoginLocation, checkLogin(c), handlers.PostLogin(c))
	router.GET(values.RegisterLocation, checkLogin(c), handlers.GetRegister(c))
	router.POST(values.RegisterLocation, checkLogin(c), handlers.PostRegister(c))
	router.GET(values.DashboardLocation, requiresLogin(c), handlers.GetDashboard(c))
	router.GET(values.LogoutLocation, requiresLogin(c), handlers.Logout(c))
	router.POST(values.LogoutLocation, requiresLogin(c), handlers.Logout(c))
	return router
}
