package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerInvoiceGet(t *testing.T) {
	req := client.NewCustomerInvoiceGetRequest()
	req.PathParams().InvoiceID = 2020002
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
