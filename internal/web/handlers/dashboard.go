package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/pivot/internal/controller"
	"io"
)

func GetDashboard(c *controller.Controller) gin.HandlerFunc {
	return func(context *gin.Context) {
		file, openError := c.OpenPage("dashboard/dashboard.html")
		if openError != nil {
			go c.LogError(context, openError)
			return
		}
		defer file.Close()
		_, copyError := io.Copy(context.Writer, file)
		if copyError != nil {
			go c.LogError(context, openError)
		}
	}
}
