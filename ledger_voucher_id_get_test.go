package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLedgerVoucherIDGet(t *testing.T) {
	req := client.NewLedgerVoucherIDGetRequest()
	req.PathParams().ID = 122633538
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
