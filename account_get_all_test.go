package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAccountGetAll(t *testing.T) {
	req := client.NewAccountGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
