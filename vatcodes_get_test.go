package vismaonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVatcodesGet(t *testing.T) {
	req := client.NewVatcodesGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
