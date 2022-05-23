package vismaonline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerCreditNoteGetAll(t *testing.T) {
	req := client.NewCustomerCreditNoteGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
