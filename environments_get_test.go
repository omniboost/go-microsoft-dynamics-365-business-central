package central_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEnvironmentsGet(t *testing.T) {
	req := client.NewEnvironmentsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
