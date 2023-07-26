package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSalesInvoicesGet(t *testing.T) {
	req := client.NewSalesInvoicesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

