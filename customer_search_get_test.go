package dkplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerSearchGet(t *testing.T) {
	req := client.NewCustomerSearchGetRequest()
	req.PathParams().SearchString = "Bogaert"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
