package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestInvoicePost(t *testing.T) {
	req := client.NewInvoicePostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
