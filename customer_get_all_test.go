package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerGetAll(t *testing.T) {
	req := client.NewCustomerGetAll()
	// req.QueryParams().PeriodID = "202103"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
