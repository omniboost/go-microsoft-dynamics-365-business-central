package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDimensionSetLinesGet(t *testing.T) {
	req := client.NewDimensionSetLinesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req.PathParams().Object = "journalLines"
	req.PathParams().ObjectID = "17785fa2-d0a3-ed11-9a88-000d3a6651aa"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
