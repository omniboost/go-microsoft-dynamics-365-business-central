package vismaonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerledgeritemswithvoucherPost(t *testing.T) {
	req := client.NewCustomerledgeritemswithvoucherPost()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
