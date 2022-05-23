package vismaonline_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	vismaonline "github.com/omniboost/go-visma.net"
)

var (
	client *vismaonline.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	accessToken := os.Getenv("ACCESS_TOKEN")
	companyID := os.Getenv("COMPANY_ID")
	applicationType := os.Getenv("APPLICATION_TYPE")
	debug := os.Getenv("DEBUG")

	client = vismaonline.NewClient(nil, accessToken, companyID, applicationType)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
