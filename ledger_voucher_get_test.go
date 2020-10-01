package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	tripletex "github.com/omniboost/go-tripletex"
)

func TestLedgerVoucherGet(t *testing.T) {
	req := client.NewLedgerVoucherGetRequest()
	req.QueryParams().DateFrom = tripletex.Date{time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)}
	req.QueryParams().DateTo = tripletex.Date{time.Now()}
	req.QueryParams().Count = 1000
	// req.QueryParams().ExpirationDate = tripletex.Date{time.Now().AddDate(0, 0, 1)}
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
