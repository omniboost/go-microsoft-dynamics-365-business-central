package multivers_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	multivers "github.com/tim-online/go-unit4-multivers"
)

func TestProductInfoList(t *testing.T) {
	setup()
	defer teardown()

	// create the request
	requestParams := multivers.NewProductInfoListGetParams()
	requestParams.Description = "description"
	requestParams.ProductGroupID = "productGroupID"
	requestParams.ProductID = "productID"
	requestParams.ShortName = "shortName"

	// Test request
	router.HandleFunc("/api/TEST/ProductInfoList{.json}", func(w http.ResponseWriter, r *http.Request) {
		// Test meta data of incoming request
		testMethod(t, r, "GET")
		testHeader(t, r, "Content-Type", "application/json; charset=utf-8")

		// convert incoming query parameters back to struct
		got := multivers.NewProductInfoListGetParams()
		err := got.FromQueryParams(r.URL.Query())
		if err != nil {
			t.Error(err)
		}

		want := requestParams

		// compare them
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got=%#v\nwant=%#v", got, want)
		}

		// Send response as in API documentation
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, `[{
			"vatCodeId" : 2,
			"eanCode" : "EAN13",
			"productState" : 9,
			"dateCreated" : "28-1-2008",
			"quantityScale" : 2,
			"shortName" : "DELAN W.",
			"priceExclVat" : 99.12,
			"description" : "DELAN W.G. 1 KG (10001N)",
			"productGroupId" : "1",
			"projectSurcharge" : 1,
			"priceInclVat" : 100,
			"productId" : "10001N1",
			"technicalStock" : 99,
			"lastUpdate" : "9-3-2017",
			"stockTransferPrice" : 0.5,
			"discountAccountId" : "8013",
			"unit" : "K",
			"projectEntryType" : "",
			"productType" : 2,
			"accountId" : "8013",
			"pricePer" : 1
		}]`)
	})

	got, err := client.ProductInfoList.Get(database, requestParams, nil)
	if err != nil {
		t.Errorf("service.ListVar returned error: %v", err)
	}

	// create wanted response
	want := &multivers.ProductInfoListGetResponse{
		multivers.ProductInfo{
			VatCodeID:          2,
			EanCode:            "EAN13",
			ProductState:       9,
			DateCreated:        multivers.NewDateNLNL(2008, 1, 28),
			QuantityScale:      2,
			ShortName:          "DELAN W.",
			PriceExclVat:       99.12,
			Description:        "DELAN W.G. 1 KG (10001N)",
			ProductGroupID:     "1",
			ProjectSurcharge:   1,
			PriceInclVat:       100,
			ProductID:          "10001N1",
			TechnicalStock:     99,
			LastUpdate:         multivers.NewDateNLNL(2017, 3, 9),
			StockTransferPrice: 0.5,
			DiscountAccountID:  "8013",
			Unit:               "K",
			ProjectEntryType:   "",
			ProductType:        2,
			AccountID:          "8013",
			PricePer:           1,
		},
	}

	// compare them
	if !reflect.DeepEqual(got, want) {
		t.Errorf("service.ListVar\n got=%#v\nwant=%#v", got, want)
	}
}
