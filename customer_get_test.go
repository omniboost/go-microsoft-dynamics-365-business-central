package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerGet(t *testing.T) {
	req := client.NewCustomerGetRequest()
	// req.QueryParams().Count = 1
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
