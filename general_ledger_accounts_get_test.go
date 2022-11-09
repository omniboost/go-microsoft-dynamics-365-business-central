package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGeneralLedgerAccountsGet(t *testing.T) {
	req := client.NewGeneralLedgerAccountsGet()
	// req.QueryParams().Filter.Set("ExternalCode eq 10020")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
