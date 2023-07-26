package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSalesInvoiceLinesGet(t *testing.T) {
	req := client.NewSalesInvoiceLinesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req.PathParams().SalesInvoiceID = "07610c2e-4219-ee11-9cc4-6045bdac9485"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

