package multivers_test

import (
	"encoding/json"
	"log"
	"testing"

	multivers "github.com/omniboost/go-unit4-multivers"
)

func TestCustomerPost(t *testing.T) {
	req := client.NewCustomerPostRequest()
	req.SetRequestBody(multivers.CustomerPostRequestBody{
		ShortName: "TEST",
		Name:      "TEST",
	})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
