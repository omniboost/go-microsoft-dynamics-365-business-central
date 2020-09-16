package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalEntryLineInfoListByPeriod(t *testing.T) {
	req := client.NewJournalEntryLineInfoListByPeriodRequest()
	req.PathParams().Year = 2020
	req.PathParams().StartPeriod = 0
	req.PathParams().EndPeriod = 9
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
