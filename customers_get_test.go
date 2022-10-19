package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomersGet(t *testing.T) {
	req := client.NewCustomersGet()
	req.QueryParams().Filter.Set("ExternalCode eq 10020")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
