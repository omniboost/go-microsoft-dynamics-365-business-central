package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerPost(t *testing.T) {
	req := client.NewCustomerPost()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
