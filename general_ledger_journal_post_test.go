package dkplus_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-dkplus"
)

func TestGeneralLedgerJournalPost(t *testing.T) {
	req := client.NewGeneralLedgerJournalPostRequest()
	req.SetRequestBody(dkplus.GeneralLedgerJournalPostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
