package tests

import (
	"github.com/shoriwe/pivot/internal/web/values"
	"net/http"
	"net/url"
	"testing"
)

func TestValidRegister(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{"Antonio"},
			values.LastAndMiddleNameArgument:    []string{"Donis"},
			values.PersonalID:                   []string{"2347283486425"},
			values.EmailArgument:                []string{"antonio@upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-password1@"},
			values.PasswordConfirmationArgument: []string{"wrong-password1@"},
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

func TestInvalidRegisterName(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{""},
			values.LastAndMiddleNameArgument:    []string{"Donis"},
			values.PersonalID:                   []string{"2347283486425"},
			values.EmailArgument:                []string{"antonio@upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-password1@"},
			values.PasswordConfirmationArgument: []string{"wrong-password1@"},
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
	} else if location.Path != values.RegisterLocation {
		t.Fatal(location.Path)
	}
}

func TestInvalidRegisterLastAndMiddleName(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{"Antonio"},
			values.LastAndMiddleNameArgument:    []string{""},
			values.PersonalID:                   []string{"2347283486425"},
			values.EmailArgument:                []string{"antonio@upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-password1@"},
			values.PasswordConfirmationArgument: []string{"wrong-password1@"},
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
	} else if location.Path != values.RegisterLocation {
		t.Fatal(location.Path)
	}
}

func TestInvalidRegisterPersonalID(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{"Antonio"},
			values.LastAndMiddleNameArgument:    []string{"Donis"},
			values.PersonalID:                   []string{"Invalid ID"},
			values.EmailArgument:                []string{"antonio@upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-password1@"},
			values.PasswordConfirmationArgument: []string{"wrong-password1@"},
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
	} else if location.Path != values.RegisterLocation {
		t.Fatal(location.Path)
	}
}

func TestInvalidRegisterEmail(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{"Antonio"},
			values.LastAndMiddleNameArgument:    []string{"Donis"},
			values.PersonalID:                   []string{"2347283486425"},
			values.EmailArgument:                []string{"antonio.upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-password1@"},
			values.PasswordConfirmationArgument: []string{"wrong-password1@"},
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
	} else if location.Path != values.RegisterLocation {
		t.Fatal(location.Path)
	}
}

func TestInvalidRegisterPassword(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{"Antonio"},
			values.LastAndMiddleNameArgument:    []string{"Donis"},
			values.PersonalID:                   []string{"2347283486425"},
			values.EmailArgument:                []string{"antonio.upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-password"},
			values.PasswordConfirmationArgument: []string{"wrong-password"},
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
	} else if location.Path != values.RegisterLocation {
		t.Fatal(location.Path)
	}
}

func TestInvalidRegisterPasswordDoesntMatch(t *testing.T) {
	baseUrl, listener := Serve()
	defer listener.Close()
	client := Client()
	response, requestError := client.PostForm(
		baseUrl+values.RegisterLocation,
		url.Values{
			values.NameArgument:                 []string{"Antonio"},
			values.LastAndMiddleNameArgument:    []string{"Donis"},
			values.PersonalID:                   []string{"2347283486425"},
			values.EmailArgument:                []string{"antonio.upb.motors.co"},
			values.PasswordArgument:             []string{"wrong-passwor1@"},
			values.PasswordConfirmationArgument: []string{"wrong-password1@"},
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
	} else if location.Path != values.RegisterLocation {
		t.Fatal(location.Path)
	}
}
