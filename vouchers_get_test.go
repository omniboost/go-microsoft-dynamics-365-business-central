package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVouchersGet(t *testing.T) {
	req := client.NewVouchersGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
