package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestAccountsGet(t *testing.T) {
	req := client.NewAccountsGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	// req.QueryParams().Filter.Set("displayName eq 'TOTAL NON OPERATING INCOME & EXPENSES'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
