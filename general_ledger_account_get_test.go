package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGeneralLedgerAccountGet(t *testing.T) {
	req := client.NewGeneralLedgerAccountGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
