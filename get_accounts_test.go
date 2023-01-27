package central_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountsGet(t *testing.T) {
	req := client.NewAccountsGet()
	req.PathParams().EnvironmentName = "DG_Sandbox001"
	req.PathParams().CompanyID = "5dfedb69-2021-ec11-8f46-00224856209b"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
