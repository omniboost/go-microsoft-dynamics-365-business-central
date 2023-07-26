package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSalesCreditMemoLinesGet(t *testing.T) {
	req := client.NewSalesCreditMemoLinesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req.PathParams().SalesCreditMemoID = "89284bda-2d26-ee11-9cbf-000d3ab979cf"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}


