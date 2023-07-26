package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSalesCreditMemoLinePost(t *testing.T) {
	req := client.NewSalesCreditMemoLinePost()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req.PathParams().SalesCreditMemoID = "52eb588e-5722-ee11-9cbf-000d3ab001d2"

	req.RequestBody().AccountID = ""
	req.RequestBody().AmountExcludingTax = 84.4
	req.RequestBody().Description = ""
	req.RequestBody().DocumentID = "52eb588e-5722-ee11-9cbf-000d3ab001d2"
	req.RequestBody().LineObjectNumber = "34510"
	req.RequestBody().LineType = "Account"
	req.RequestBody().TaxCode = ""

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}


