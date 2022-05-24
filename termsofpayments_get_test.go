package vismaonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTermsofpaymentsGet(t *testing.T) {
	req := client.NewTermsofpaymentsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
