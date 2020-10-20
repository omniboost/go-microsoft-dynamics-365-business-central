package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLedgerPostingIDGet(t *testing.T) {
	req := client.NewLedgerPostingIDGetRequest()
	req.PathParams().ID = 603915264
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
