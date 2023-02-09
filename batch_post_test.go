package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestBatchPost(t *testing.T) {
	req := client.NewBatchPost()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")

	req2 := client.NewJournalLinesPost()
	req2.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	req2.PathParams().CompanyID = os.Getenv("COMPANY_ID")
	req2.PathParams().JournalID = os.Getenv("COMPANY_ID")

	bresp, err := req.FromRequest(&req2)
	if err != nil {
		t.Error(err)
	}
	req.RequestBody().Requests = append(req.RequestBody().Requests, bresp)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// func TestBatchPost(t *testing.T) {
// 	req := client.NewBatchPost()
// 	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")

// 	req2 := client.NewJournalLinesPost()
// 	req2.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
// 	req2.PathParams().CompanyID = os.Getenv("COMPANY_ID")
// 	req2.PathParams().JournalID = os.Getenv("COMPANY_ID")
// 	hreq2, err := client.NewRequest(nil, &req2)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// req2 := client.NewEnvironmentsGet()
// 	// hreq2, err := client.NewRequest(nil, &req2)
// 	// if err != nil {
// 	// 	t.Error(err)
// 	// }

// 	// Process query parameters
// 	err = utils.AddQueryParamsToRequest(req2.QueryParams(), hreq2, false)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	bresp, err := req.FromHTTPRequest(hreq2)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	req.RequestBody().Requests = append(req.RequestBody().Requests, bresp)
// 	// req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
// 	// req.PathParams().CompanyID = os.Getenv("COMPANY_ID")
// 	// req.PathParams().JournalID = "61e3d254-9913-ec11-86bc-000d3ac818b6"

// 	resp, err := req.Do()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	b, _ := json.MarshalIndent(resp, "", "  ")
// 	fmt.Println(string(b))
// }
