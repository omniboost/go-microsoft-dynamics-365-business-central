package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	tripletex "github.com/omniboost/go-tripletex"
)

func TestInvoceGet(t *testing.T) {
	req := client.NewInvoceGetRequest()
	req.QueryParams().InvoiceDateFrom = tripletex.Date{time.Now().AddDate(0, 0, -50)}
	req.QueryParams().InvoiceDateTo = tripletex.Date{time.Now().AddDate(0, 0, 0)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
