package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	central "github.com/omniboost/go-microsoft-dynamics-365-business-central"
)

func TestJournalsGet(t *testing.T) {
	req := client.NewJournalsGet()
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

func TestJournalsGetAllCompanies(t *testing.T) {
	out := []struct {
		Company  central.Company
		Journals central.Journals
	}{}

	req := client.NewCompaniesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	for k, c := range resp.Value {
		if len(out) <= k {
			out = append(out, struct {
				Company  central.Company
				Journals central.Journals
			}{
				Company:  c,
				Journals: central.Journals{},
			})
		}

		req := client.NewJournalsGet()
		req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
		req.PathParams().CompanyID = c.ID
		// req.QueryParams().Filter.Set("code eq 'JV'")
		resp, err := req.Do()
		if err != nil {
			t.Error(err)
		}

		out[k].Journals = resp.Value
	}

	b, _ := json.MarshalIndent(out, "", "   ")
	fmt.Println(string(b))
}
