package vismaonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCompanysettingsGet(t *testing.T) {
	req := client.NewCompanysettingsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
