package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLocationsGet(t *testing.T) {
	req := client.NewLocationsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
