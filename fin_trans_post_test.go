package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestFinTransPost(t *testing.T) {
	req := client.NewFinTransPostRequest()
	req.SetRequestBody(multivers.FinTransPostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
