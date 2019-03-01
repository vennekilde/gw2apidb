package gw2api

import (
	"fmt"
	"os"
	"testing"
)

func ExampleNewGW2Api() {
	api := NewGW2Api()
	build, _ := api.Build()
	fmt.Printf("%d\n", build)
}

func TestNewGW2Api(t *testing.T) {
	api := NewGW2Api()
	if ver, err := api.Build(); err != nil {
		t.Error(err)
	} else if ver < 1 {
		t.Error("Build version is corrupted")
	}
}

func ExampleNewAuthenticatedGW2Api() {
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api("<APIKEY>"); err != nil {
		fmt.Printf("Failed to connect to the API: %s", err)
	}
	if api.HasPermission(PermAccount) {
		var acc Account
		if acc, err = api.Account(); err != nil {
			fmt.Printf("API did not answer correctly: %s", err)
		}
		fmt.Printf("%s\n", acc.Name)
	}
}

func TestNewAuthenticatedGW2Api(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermAccount) {
		t.Error("Failed to obtain basic permission")
	}
}

func TestSetAuthentication(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}

	api := NewGW2Api()
	if _, err := api.Account(); err == nil {
		t.Error("Should be able to call authenticated endpoint")
	}

	if err := api.SetAuthentication("FOOBAR"); err == nil {
		t.Error("Failed to catch error for incorrectly formatted APIKEY")
	}

	if err := api.SetAuthentication(apikey); err != nil {
		t.Error("Failed to authenticate: ", err)
	}

	if _, err := api.Account(); err != nil {
		t.Error("Authentication should be available now")
	}
}

func TestAPIError(t *testing.T) {
	type ProbablyError struct {
		Foo int `json:"error"`
	}
	api := NewGW2Api()
	var foo ProbablyError
	if err := api.fetchEndpoint("v2", "foobar", nil, &foo); err == nil {
		t.Error("Should not have been nil")
	}
}
