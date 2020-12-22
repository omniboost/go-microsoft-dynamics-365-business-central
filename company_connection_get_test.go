package dkplus_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCompanyConnectionGet(t *testing.T) {
	req := client.NewCompanyConnectionGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
