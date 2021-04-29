package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerInvoiceGet(t *testing.T) {
	req := client.NewCustomerInvoiceGet()
	req.PathParams().InvoiceNumber = 202151433
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
