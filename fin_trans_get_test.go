package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestFinTransGet(t *testing.T) {
	req := client.NewFinTransGetRequest()
	req.PathParams().Year = 2020
	req.PathParams().JournalID = "M"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
