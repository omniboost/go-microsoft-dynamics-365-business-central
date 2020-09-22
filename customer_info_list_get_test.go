package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerInfoListGet(t *testing.T) {
	req := client.NewCustomerInfoListGetRequest()
	req.QueryParams().Filter.Set("Email eq 'info@melbicks.h-p.co.uk' or Name eq 'Inkoopcombinatie CO-OP'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
