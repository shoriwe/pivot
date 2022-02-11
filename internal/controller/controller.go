package controller

import (
	"embed"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/pivot/internal/data"
	"github.com/shoriwe/pivot/internal/data/objects"
	"github.com/shoriwe/pivot/internal/logs"
	"io/fs"
	"path"
	"regexp"
)

var (
	emailRegexp           = regexp.MustCompile(`\w+[\w\._-]*\w+@\w+[\\w\._-]*\w+`)
	personalIDRegexp      = regexp.MustCompile(`(?m)^\d+$`)
	passwordNumbersRegexp = regexp.MustCompile(`\d+`)
	passwordSpecialRegexp = regexp.MustCompile(`[[:punct:]]+`)
	nameRegexp            = regexp.MustCompile(`[\w|\s]+`)
)

func checkEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

func checkPassword(password string) bool {
	passwordLength := len(password)
	if passwordLength < 8 {
		return false
	}
	return len(passwordNumbersRegexp.FindAllString(password, -1)) > 0 && len(passwordSpecialRegexp.FindAllString(password, -1)) > 0
}

func checkPersonalID(personalID string) bool {
	return personalIDRegexp.MatchString(personalID)
}

func checkName(name string) bool {
	return nameRegexp.MatchString(name)
}

func checkUser(user *objects.User) (string, bool) {
	if !checkName(user.Name) {
		return "Invalid Name", false
	}
	if !checkName(user.LastAndMiddleName) {
		return "Invalid Middle and Last name", false
	}
	if !checkEmail(user.Email) {
		return "Invalid email", false
	}
	if !checkPassword(user.Password) {
		return "Invalid password, it should have at least 8 characters, 1 number and 1 special character", false
	}
	if !checkPersonalID(user.PersonalID) {
		return "Invalid Personal ID", false
	}
	return "Everything is ok", true
}

type Controller struct {
	Connection *data.Connection
	PagesFS    *embed.FS
	*logs.Logger
}

func (controller *Controller) Register(context *gin.Context, user *objects.User) (string, bool) {
	errorMessage, isValid := checkUser(user)
	if !isValid {
		return errorMessage, false
	}
	user.Password = hex.EncodeToString(data.CalcHash(user.Password))
	succeed, registerError := controller.Connection.DB.Register(user)
	go controller.LogRegistration(context, user, succeed)
	if succeed {
		return "", true
	} else if registerError != nil {
		go controller.LogError(context, registerError)
	}
	return "Registration failed", false
}

func (controller *Controller) Logout(context *gin.Context, cookie string) bool {
	return controller.Connection.Cache.DeleteUserSession(cookie)
}

func (controller *Controller) Login(context *gin.Context, email, password string) (*objects.User, string, bool) {
	user, succeed, loginError := controller.Connection.DB.Login(email, password)
	go controller.LogLoginAttempt(context, email, succeed)
	if succeed {
		cookies := controller.Connection.Cache.NewUserSession(user)
		return user, cookies, true
	} else if loginError != nil {
		go controller.LogError(context, loginError)
	}
	return user, "", false
}

func (controller *Controller) OpenPage(p string) (fs.File, error) {
	return controller.PagesFS.Open(path.Join("pages/", p))
}

func NewController(connection *data.Connection, logger *logs.Logger, static, pages *embed.FS) *Controller {
	return &Controller{
		Connection: connection,
		Logger:     logger,
		PagesFS:    pages,
	}
}
