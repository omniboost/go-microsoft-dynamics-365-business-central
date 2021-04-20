package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerCreditNoteV2Post(t *testing.T) {
	req := client.NewCustomerCreditNoteV2Post()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
