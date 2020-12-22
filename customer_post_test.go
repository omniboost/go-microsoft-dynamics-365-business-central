package dkplus_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-dkplus"
)

func TestCustomerPost(t *testing.T) {
	req := client.NewCustomerPostRequest()
	req.SetRequestBody(dkplus.CustomerPostRequestBody{
		Number:      "4321",
		Name:        "Leon Bogaert",
		Address1:    "Stadhuisplein 3",
		Address2:    "Terneuzen",
		Zipcode:     "4531 GZ",
		Email:       "",
		Phone:       "",
		PhoneLocal:  "",
		PhoneMobile: "",
	})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
