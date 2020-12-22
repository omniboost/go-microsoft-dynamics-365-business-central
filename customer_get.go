package dkplus

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCustomerGetRequest() CustomerGetRequest {
	r := CustomerGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerGetRequest struct {
	client      *Client
	queryParams *CustomerGetQueryParams
	pathParams  *CustomerGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerGetRequestBody
}

func (r CustomerGetRequest) NewQueryParams() *CustomerGetQueryParams {
	return &CustomerGetQueryParams{}
}

type CustomerGetQueryParams struct {
}

func (p CustomerGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerGetRequest) QueryParams() *CustomerGetQueryParams {
	return r.queryParams
}

func (r CustomerGetRequest) NewPathParams() *CustomerGetPathParams {
	return &CustomerGetPathParams{}
}

type CustomerGetPathParams struct {
	Objects bool `schema:"objects"`
}

func (p *CustomerGetPathParams) Params() map[string]string {
	return map[string]string{
		"objects": func() string {
			if p.Objects {
				return "true"
			}
			return "false"
		}(),
	}
}

func (r *CustomerGetRequest) PathParams() *CustomerGetPathParams {
	return r.pathParams
}

func (r *CustomerGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerGetRequest) Method() string {
	return r.method
}

func (r CustomerGetRequest) NewRequestBody() CustomerGetRequestBody {
	return CustomerGetRequestBody{}
}

type CustomerGetRequestBody struct {
}

func (r *CustomerGetRequest) RequestBody() *CustomerGetRequestBody {
	return &r.requestBody
}

func (r *CustomerGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerGetRequest) SetRequestBody(body CustomerGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerGetRequest) NewResponseBody() *CustomerGetResponseBody {
	return &CustomerGetResponseBody{}
}

type CustomerGetResponseBody []struct {
	Conditions struct {
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

func (r *CustomerGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("customer/{{.objects}}", r.PathParams())
	return &u
}

func (r *CustomerGetRequest) Do() (CustomerGetResponseBody, error) {
	// if r.QueryParams().ConsumerToken == "" {
	// 	r.QueryParams().ConsumerToken = r.client.ConsumerToken()
	// }

	// if r.QueryParams().EmployeeToken == "" {
	// 	r.QueryParams().EmployeeToken = r.client.EmployeeToken()
	// }

	// if r.QueryParams().ExpirationDate.IsZero() {
	// 	r.QueryParams().ExpirationDate = Date{time.Now().AddDate(0, 0, 1)}
	// }

	// fetch a new token if it isn't set already
	// if r.client.token == "" {
	// 	var err error
	// 	r.client.token, err = r.client.NewToken()
	// 	if err != nil {
	// 		return *r.NewResponseBody(), err
	// 	}
	// }

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
