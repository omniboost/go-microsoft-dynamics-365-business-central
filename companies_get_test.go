package central_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCompaniesGet(t *testing.T) {
	req := client.NewCompaniesGet()
	req.PathParams().EnvironmentName = "DG_Sandbox001"
	req.PathParams().EnvironmentName = "production"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
