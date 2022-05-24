package vismaonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomersGet(t *testing.T) {
	req := client.NewCustomersGet()
	req.QueryParams().Filter.Set("Name eq 'Leon Bogaert'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
