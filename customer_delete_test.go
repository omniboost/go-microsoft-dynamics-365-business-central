package dkplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerDelete(t *testing.T) {
	req := client.NewCustomerDeleteRequest()
	req.PathParams().Number = "4321"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
