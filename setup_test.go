package poweroffice_test

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	poweroffice "github.com/omniboost/go-poweroffice"
)

var (
	client *poweroffice.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	// refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := poweroffice.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client = poweroffice.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		u, err := url.Parse(baseURL)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*u)
		oauthConfig.SetBaseURL(u)
	}

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.TokenURL = tokenURL
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
