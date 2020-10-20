package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLedgerVoucherPost(t *testing.T) {
	req := client.NewLedgerVoucherPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
