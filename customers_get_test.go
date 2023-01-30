package central_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomersGet(t *testing.T) {
	req := client.NewCustomersGet()
	req.PathParams().EnvironmentName = "production"
	req.PathParams().CompanyID = "7da8c511-282d-ec11-8f46-0022485628fc"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
