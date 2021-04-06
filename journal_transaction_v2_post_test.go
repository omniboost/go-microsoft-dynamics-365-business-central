package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalTransactionV2Post(t *testing.T) {
	req := client.NewJournalTransactionV2Post()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
