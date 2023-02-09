package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDimensionsGet(t *testing.T) {
	req := client.NewDimensionsGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	// req.QueryParams().Filter.Set("code eq 'JV'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
