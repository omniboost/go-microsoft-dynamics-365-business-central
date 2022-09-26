package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerledgeritemsGet(t *testing.T) {
	req := client.NewCustomerledgeritemsGet()
	// req.QueryParams().Filter.Set("Name eq 'Leon Bogaert'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
