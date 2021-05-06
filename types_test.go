package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"

	vismanet "github.com/omniboost/go-visma.net"
)

func TestValueString(t *testing.T) {
	// t1 := vismanet.ValueString("")
	// b, err := json.Marshal(t1)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// var t2 *vismanet.ValueString
	// b, err = json.Marshal(t2)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// b = []byte(`
	// 	{"T": null}
	// `)
	// t3 := struct {
	// 	T *vismanet.ValueString
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

	t1 := vismanet.ValueNullString{nil}
	b, err := json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))

	var s vismanet.ValueString
	s = ""
	t1 = vismanet.ValueNullString{&s}
	b, err = json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
}
