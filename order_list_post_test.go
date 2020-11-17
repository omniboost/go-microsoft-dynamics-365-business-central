package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestOrderListPost(t *testing.T) {
	req := client.NewOrderListPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
