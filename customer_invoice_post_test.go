package multivers_test

import (
	"encoding/json"
	"log"
	"testing"

	multivers "github.com/omniboost/go-unit4-multivers"
)

func TestCustomerInvoicePost(t *testing.T) {
	req := client.NewCustomerInvoicePostRequest()
	req.SetRequestBody(multivers.CustomerInvoicePostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
