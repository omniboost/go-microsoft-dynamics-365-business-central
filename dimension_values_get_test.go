package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDimensionValuesGet(t *testing.T) {
	req := client.NewDimensionValuesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req.PathParams().DimensionID = "b7eb4f33-068b-ed11-aad6-000d3abcd607"
	// req.QueryParams().Filter.Set("code eq 'JV'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
