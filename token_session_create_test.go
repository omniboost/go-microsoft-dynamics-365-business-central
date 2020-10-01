package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestTokenSessionCreate(t *testing.T) {
	req := client.NewTokenSessionCreateRequest()
	// req.QueryParams().ExpirationDate = tripletex.Date{time.Now().AddDate(0, 0, 1)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
