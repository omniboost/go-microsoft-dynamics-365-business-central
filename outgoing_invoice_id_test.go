package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOutgoingInvoiceIDGet(t *testing.T) {
	req := client.NewOutgoingInvoiceIDGet()
	req.PathParams().ID = "5b6d33cc-f3f2-426b-8ab7-617227a6fdde"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
