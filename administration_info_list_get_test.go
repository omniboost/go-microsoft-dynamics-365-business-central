package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAdministrationInfoListGet(t *testing.T) {
	req := client.NewAdministrationInfoListGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
