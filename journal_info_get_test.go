package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalInfoGet(t *testing.T) {
	req := client.NewJournalInfoGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
