package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	guestline "github.com/omniboost/go-guestline"
)

func TestGetDepartures(t *testing.T) {
	req := client.NewGetDeparturesRequest()
	req.RequestBody().DepartureDate = guestline.DateTime{time.Now().AddDate(0, 0, -2)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
