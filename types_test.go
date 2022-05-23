package vismaonline_test

import (
	"encoding/json"
	"log"
	"testing"

	vismaonline "github.com/omniboost/go-visma.net"
)

func TestValueString(t *testing.T) {
	// t1 := vismaonline.ValueString("")
	// b, err := json.Marshal(t1)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// var t2 *vismaonline.ValueString
	// b, err = json.Marshal(t2)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// b = []byte(`
	// 	{"T": null}
	// `)
	// t3 := struct {
	// 	T *vismaonline.ValueString
	// }{}
	// err = json.Unmarshal(b, &t3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// b, err = json.Marshal(t3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	t1 := vismaonline.ValueNullString{nil}
	b, err := json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))

	var s vismaonline.ValueString
	s = ""
	t1 = vismaonline.ValueNullString{&s}
	b, err = json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
}
