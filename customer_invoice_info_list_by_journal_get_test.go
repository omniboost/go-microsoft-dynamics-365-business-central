package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerInvoiceInfoListByJournalGet(t *testing.T) {
	req := client.NewCustomerInvoiceInfoListByJournalGetRequest()
	req.PathParams().FiscalYear = 2020
	req.PathParams().JournalID = "V"
	req.PathParams().InvoiceState = 0
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
