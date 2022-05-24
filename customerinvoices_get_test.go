package vismaonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerinvoicesGet(t *testing.T) {
	req := client.NewCustomerinvoicesGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
