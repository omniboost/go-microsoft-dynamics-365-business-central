package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArticlesGet(t *testing.T) {
	req := client.NewArticlesGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
