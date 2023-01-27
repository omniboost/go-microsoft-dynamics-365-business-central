package central_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJournalLinesGet(t *testing.T) {
	req := client.NewJournalLinesGet()
	req.PathParams().EnvironmentName = "DG_Sandbox001"
	req.PathParams().CompanyID = "5dfedb69-2021-ec11-8f46-00224856209b"
	req.PathParams().JournalID = "45a7023d-e311-ec11-86bc-000d3ac818b6"

	req.PathParams().EnvironmentName = "production"
	req.PathParams().CompanyID = "7da8c511-282d-ec11-8f46-0022485628fc"
	req.PathParams().JournalID = "61e3d254-9913-ec11-86bc-000d3ac818b6"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
