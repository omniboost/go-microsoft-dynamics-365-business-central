package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJournalLinesGet(t *testing.T) {
	req := client.NewJournalLinesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req.PathParams().JournalID = "61e3d254-9913-ec11-86bc-000d3ac818b6"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
