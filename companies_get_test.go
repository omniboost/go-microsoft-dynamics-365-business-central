package central_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestCompaniesGet(t *testing.T) {
	req := client.NewCompaniesGet()
	req.PathParams().EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
