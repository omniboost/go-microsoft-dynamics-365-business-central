package guestline_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	guestline "github.com/omniboost/go-guestline"
)

var (
	client *guestline.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	siteID := os.Getenv("SITE_ID")
	interfaceID := os.Getenv("INTERFACE_ID")
	operatorCode := os.Getenv("OPERATOR_CODE")
	password := os.Getenv("PASSWORD")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = guestline.NewClient(nil, siteID, interfaceID, operatorCode, password)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	m.Run()
}
