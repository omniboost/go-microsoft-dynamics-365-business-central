package tripletex_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	tripletex "github.com/omniboost/go-tripletex"
)

var (
	client *tripletex.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	consumerToken := os.Getenv("CONSUMER_TOKEN")
	employeeToken := os.Getenv("EMPLOYEE_TOKEN")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = tripletex.NewClient(nil, consumerToken, employeeToken)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
