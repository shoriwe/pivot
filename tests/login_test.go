package tests

import (
	"github.com/shoriwe/pivot/internal/web/values"
	"net/http"
	"net/url"
	"testing"
)

func TestSuccessLogin(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.LoginLocation,
		url.Values{
			values.EmailArgument:    []string{"admin@upb.motors.co"},
			values.PasswordArgument: []string{"admin"},
		},
	)
	if requestError != nil {
		t.Fatal(requestError)
	}
	cookies := response.Header.Get("Set-Cookie")
	if len(cookies) == 0 {
		t.Fatal("Login failed")
	}
	cookies = cookies[6:]
	if len(cookies) == 0 {
		t.Fatal("Login failed")
	}
}

func TestFailLogin(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.LoginLocation,
		url.Values{
			values.EmailArgument:    []string{"admin@upb.motors.co"},
			values.PasswordArgument: []string{"wrong-password"},
		},
	)
	if requestError != nil {
		t.Fatal(requestError)
	}
	if response.StatusCode != http.StatusFound {
		t.Fatal("Status code was not expected")
	}
	if location, err := response.Location(); err != nil {
		t.Fatal(err)
	} else if location.Path != values.LoginLocation {
		t.Fatal(location.Path)
	}
}
