package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewSalesInvoicePostRequest() SalesInvoicePostRequest {
	r := SalesInvoicePostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicePostRequest struct {
	client      *Client
	queryParams *SalesInvoicePostQueryParams
	pathParams  *SalesInvoicePostPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoicePostRequestBody
}

func (r SalesInvoicePostRequest) NewQueryParams() *SalesInvoicePostQueryParams {
	return &SalesInvoicePostQueryParams{}
}

type SalesInvoicePostQueryParams struct {
	Post bool `schema:"post"`
}

func (p SalesInvoicePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoicePostRequest) QueryParams() *SalesInvoicePostQueryParams {
	return r.queryParams
}

func (r SalesInvoicePostRequest) NewPathParams() *SalesInvoicePostPathParams {
	return &SalesInvoicePostPathParams{}
}

type SalesInvoicePostPathParams struct {
}

func (p *SalesInvoicePostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesInvoicePostRequest) PathParams() *SalesInvoicePostPathParams {
	return r.pathParams
}

func (r *SalesInvoicePostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicePostRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicePostRequest) Method() string {
	return r.method
}

func (r SalesInvoicePostRequest) NewRequestBody() SalesInvoicePostRequestBody {
	return SalesInvoicePostRequestBody{}
}

type SalesInvoicePostRequestBody struct {
	Customer struct {
		Number   string `json:"Number"`
		Name     string `json:"Name,omitempty"`
		Address1 string `json:"Address1,omitempty"`
		Address2 string `json:"Address2,omitempty"`
		ZipCode  string `json:"ZipCode,omitempty"`
		Country  string `json:"Country,omitempty"`
	} `json:"Customer"`
	SaleType  SaleType `json:"SaleType"`
	Reference string   `json:"Reference"`
	Text1     string   `json:"Text1"`
	Text2     string   `json:"Text2"`
	Term      string   `json:"Term"`
	Options   struct {
		OriginalPrices bool `json:"OriginalPrices"`
	} `json:"Options"`
	Date        DateTime `json:"Date"`
	DueDate     DateTime `json:"DueDate"`
	Currency    string   `json:"Currency"`
	SalesPerson string   `json:"SalesPerson,omitempty"`
	Exchange    float64  `json:"Exchange"`
	Lines       []struct {
		ItemCode       string  `json:"ItemCode"`
		WareHouse      string  `json:"WareHouse"`
		Text           string  `json:"Text"`
		Text2          string  `json:"Text2"`
		Quantity       float64 `json:"Quantity"`
		Reference      string  `json:"Reference"`
		IncludingVAT   bool    `json:"IncludingVAT"`
		Price          float64 `json:"Price"`
		Discount       float64 `json:"Discount"`
		DiscountAmount float64 `json:"DiscountAmount"`
		Dim1           string  `json:"Dim1"`
		Variations     []struct {
		} `json:"Variations"`
	} `json:"Lines"`
	Payments []struct {
		ID     int     `json:"ID"`
		Name   string  `json:"Name"`
		Amount float64 `json:"Amount"`
	} `json:"Payments"`
	Attachment File `json:"Attachment"`
	// Receiver   struct {
	// } `json:"Receiver"`
}

func (r *SalesInvoicePostRequest) RequestBody() *SalesInvoicePostRequestBody {
	return &r.requestBody
}

func (r *SalesInvoicePostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SalesInvoicePostRequest) SetRequestBody(body SalesInvoicePostRequestBody) {
	r.requestBody = body
}

func (r *SalesInvoicePostRequest) NewResponseBody() *SalesInvoicePostResponseBody {
	return &SalesInvoicePostResponseBody{}
}

type SalesInvoicePostResponseBody struct {
}

func (r *SalesInvoicePostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("sales/invoice", r.PathParams())
	return &u
}

func (r *SalesInvoicePostRequest) Do() (SalesInvoicePostResponseBody, error) {
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
