package dkplus

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCustomerSearchGetRequest() CustomerSearchGetRequest {
	r := CustomerSearchGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerSearchGetRequest struct {
	client      *Client
	queryParams *CustomerSearchGetQueryParams
	pathParams  *CustomerSearchGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerSearchGetRequestBody
}

func (r CustomerSearchGetRequest) NewQueryParams() *CustomerSearchGetQueryParams {
	return &CustomerSearchGetQueryParams{}
}

type CustomerSearchGetQueryParams struct {
}

func (p CustomerSearchGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerSearchGetRequest) QueryParams() *CustomerSearchGetQueryParams {
	return r.queryParams
}

func (r CustomerSearchGetRequest) NewPathParams() *CustomerSearchGetPathParams {
	return &CustomerSearchGetPathParams{}
}

type CustomerSearchGetPathParams struct {
	SearchString string `schema:"searchstring"`
}

func (p *CustomerSearchGetPathParams) Params() map[string]string {
	return map[string]string{
		"searchstring": p.SearchString,
	}
}

func (r *CustomerSearchGetRequest) PathParams() *CustomerSearchGetPathParams {
	return r.pathParams
}

func (r *CustomerSearchGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerSearchGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerSearchGetRequest) Method() string {
	return r.method
}

func (r CustomerSearchGetRequest) NewRequestBody() CustomerSearchGetRequestBody {
	return CustomerSearchGetRequestBody{}
}

type CustomerSearchGetRequestBody struct {
}

func (r *CustomerSearchGetRequest) RequestBody() *CustomerSearchGetRequestBody {
	return &r.requestBody
}

func (r *CustomerSearchGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerSearchGetRequest) SetRequestBody(body CustomerSearchGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerSearchGetRequest) NewResponseBody() *CustomerSearchGetResponseBody {
	return &CustomerSearchGetResponseBody{}
}

type CustomerSearchGetResponseBody []struct {
	Contacts   []interface{} `json:"Contacts"`
	Recivers   []interface{} `json:"Recivers"`
	Memos      []interface{} `json:"Memos"`
	Properties []interface{} `json:"Properties"`
	Changes    []struct {
		Modified time.Time `json:"Modified"`
		By       string    `json:"By"`
		Fields   []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Fields"`
	} `json:"Changes"`
	Attachments []interface{} `json:"Attachments"`
	Conditions  struct {
		CreditLimit   float64 `json:"CreditLimit"`
		CreditMax     float64 `json:"CreditMax"`
		DisableSale   bool    `json:"DisableSale"`
		DisableRetail bool    `json:"DisableRetail"`
	} `json:"Conditions"`
	SendTo struct {
		Printer          bool `json:"Printer"`
		ClaimToPrinter   bool `json:"ClaimToPrinter"`
		Email            bool `json:"Email"`
		EDIInvoice       bool `json:"EDIInvoice"`
		PublishingSystem bool `json:"PublishingSystem"`
	} `json:"SendTo"`
	RecordID          int       `json:"RecordID"`
	Number            string    `json:"Number"`
	Name              string    `json:"Name"`
	Alias             string    `json:"Alias"`
	Address1          string    `json:"Address1"`
	Address2          string    `json:"Address2"`
	Address3          string    `json:"Address3"`
	ZipCode           string    `json:"ZipCode"`
	BalanceAmount     float64   `json:"BalanceAmount"`
	Phone             string    `json:"Phone"`
	PhoneLocal        string    `json:"PhoneLocal"`
	PhoneMobile       string    `json:"PhoneMobile"`
	PhoneFax          string    `json:"PhoneFax"`
	CountryCode       string    `json:"CountryCode"`
	OriginCountryCode string    `json:"OriginCountryCode"`
	Email             string    `json:"Email"`
	Password          string    `json:"Password"`
	Group             string    `json:"Group"`
	SalesPerson       string    `json:"SalesPerson"`
	Discount          float64   `json:"Discount"`
	UseItemRecivers   bool      `json:"UseItemRecivers"`
	PaymentTerm       string    `json:"PaymentTerm"`
	PaymentMode       string    `json:"PaymentMode"`
	CurrencyCode      string    `json:"CurrencyCode"`
	NoVat             bool      `json:"NoVat"`
	LedgerCode        string    `json:"LedgerCode"`
	Blocked           bool      `json:"Blocked"`
	Gender            int       `json:"Gender"`
	PriceGroup        int       `json:"PriceGroup"`
	BillingFee        float64   `json:"BillingFee"`
	Modified          time.Time `json:"Modified"`
}

func (r *CustomerSearchGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("customer/search/{{.searchstring}}", r.PathParams())
	return &u
}

func (r *CustomerSearchGetRequest) Do() (CustomerSearchGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
