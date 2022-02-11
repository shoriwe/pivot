package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/pivot/internal/controller"
	"github.com/shoriwe/pivot/internal/data/objects"
	"github.com/shoriwe/pivot/internal/web/values"
	"io"
	"net/http"
)

func GetRegister(c *controller.Controller) gin.HandlerFunc {
	return func(context *gin.Context) {
		file, openError := c.OpenPage("register.html")
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

func PostRegister(c *controller.Controller) gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			password             = context.PostForm(values.PasswordArgument)
			passwordConfirmation = context.PostForm(values.PasswordConfirmationArgument)
			registrationError    string
			succeed              bool
		)
		if password == passwordConfirmation {
			user := &objects.User{
				Name:              context.PostForm(values.NameArgument),
				LastAndMiddleName: context.PostForm(values.LastAndMiddleNameArgument),
				PersonalID:        context.PostForm(values.PersonalID),
				Email:             context.PostForm(values.EmailArgument),
				Password:          password,
			}
			registrationError, succeed = c.Register(context, user)
			if succeed {
				context.Redirect(http.StatusFound, values.LoginLocation)
				return
			}
		} else {
			registrationError = "password doesn't match"
		}
		file, openError := c.OpenPage("failed-register.html")
		if openError != nil {
			go c.LogError(context, openError)
			return
		}
		contents, _ := io.ReadAll(file)
		final := fmt.Sprintf("<script>alert(\"%s\")</script>", registrationError)
		defer file.Close()
		_, copyError := io.Copy(context.Writer, bytes.NewReader(append(contents, []byte(final)...)))
		if copyError != nil {
			go c.LogError(context, openError)
		}
	}
}
