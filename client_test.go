package multivers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/mux"
	multivers "github.com/tim-online/go-unit4-multivers"
)

const (
	username = "info@tim-online.nl"
	password = "mysecret"
	database = "TEST"
)

var (
	router *mux.Router
	client *multivers.Client
	server *httptest.Server
)

func setup() {
	router = mux.NewRouter()
	server = httptest.NewServer(router)

	// set custom http client
	httpClient := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set testing endpoint
	baseURL, _ := url.Parse(server.URL)
	client = multivers.NewClient(httpClient, baseURL)

	// enable debugging
	client.SetDebug(true)
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	got := r.Method
	if expected != got {
		t.Errorf("Request method = %v, expected %v", got, expected)
	}
}

func testHeader(t *testing.T, r *http.Request, key string, expected string) {
	got := r.Header.Get(key)
	if expected != got {
		t.Errorf("Request header %v = %v, expected %v", key, got, expected)
	}
}
