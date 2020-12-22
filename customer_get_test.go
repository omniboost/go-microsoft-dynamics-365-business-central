package dkplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerGet(t *testing.T) {
	req := client.NewCustomerGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
