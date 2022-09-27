package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestManualJournalVoucherPost(t *testing.T) {
	req := client.NewManualJournalVoucherPost()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
