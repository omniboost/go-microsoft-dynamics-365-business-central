package dkplus_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/Azure/go-ntlmssp"
	guestline "github.com/omniboost/go-dkplus"
)

var (
	client *guestline.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	token := os.Getenv("TOKEN")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = guestline.NewClient(nil, token)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	client.SetBeforeRequestDo(func(httpClient *http.Client, req *http.Request, body interface{}) {
		tr := &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			ForceAttemptHTTP2: false,
		}

		ntlmTransport := ntlmssp.Negotiator{
			RoundTripper: tr,
		}

		httpClient.Transport = ntlmTransport
	})

	m.Run()
}
