package guestline_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	guestline "github.com/omniboost/go-dkplus"
)

func TestGetArrivals(t *testing.T) {
	req := client.NewGetArrivalsRequest()
	req.RequestBody().ArrivalDate = guestline.DateTime{time.Now().AddDate(0, 0, -7)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
