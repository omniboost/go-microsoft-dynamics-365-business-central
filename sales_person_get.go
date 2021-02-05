package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewSalesPersonGetRequest() SalesPersonGetRequest {
	r := SalesPersonGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesPersonGetRequest struct {
	client      *Client
	queryParams *SalesPersonGetQueryParams
	pathParams  *SalesPersonGetPathParams
	method      string
	headers     http.Header
	requestBody SalesPersonGetRequestBody
}

func (r SalesPersonGetRequest) NewQueryParams() *SalesPersonGetQueryParams {
	return &SalesPersonGetQueryParams{}
}

type SalesPersonGetQueryParams struct {
}

func (p SalesPersonGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesPersonGetRequest) QueryParams() *SalesPersonGetQueryParams {
	return r.queryParams
}

func (r SalesPersonGetRequest) NewPathParams() *SalesPersonGetPathParams {
	return &SalesPersonGetPathParams{}
}

type SalesPersonGetPathParams struct {
}

func (p *SalesPersonGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesPersonGetRequest) PathParams() *SalesPersonGetPathParams {
	return r.pathParams
}

func (r *SalesPersonGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesPersonGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesPersonGetRequest) Method() string {
	return r.method
}

func (r SalesPersonGetRequest) NewRequestBody() SalesPersonGetRequestBody {
	return SalesPersonGetRequestBody{}
}

type SalesPersonGetRequestBody struct {
}

func (r *SalesPersonGetRequest) RequestBody() *SalesPersonGetRequestBody {
	return &r.requestBody
}

func (r *SalesPersonGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SalesPersonGetRequest) SetRequestBody(body SalesPersonGetRequestBody) {
	r.requestBody = body
}

func (r *SalesPersonGetRequest) NewResponseBody() *SalesPersonGetResponseBody {
	return &SalesPersonGetResponseBody{}
}

type SalesPersonGetResponseBody []struct {
	Number            string `json:"Number"`
	Employee          string `json:"Employee"`
	NameOnSalesOrders string `json:"NameOnSalesOrders"`
	Modified          string `json:"Modified"`
	Created           string `json:"Created"`
	PriceGroup        int    `json:"PriceGroup"`
	Price1Closed      bool   `json:"Price1Closed"`
	Price2Closed      bool   `json:"Price2Closed"`
	Price3Closed      bool   `json:"Price3Closed"`
	CanChangeDueDate  bool   `json:"CanChangeDueDate"`
	FilterOnCustomer  bool   `json:"FilterOnCustomer"`
	Warehouse         string `json:"Warehouse,omitempty"`
	ItemCode          string `json:"ItemCode,omitempty"`
}

func (r *SalesPersonGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("sales/person", r.PathParams())
	return &u
}

func (r *SalesPersonGetRequest) Do() (SalesPersonGetResponseBody, error) {
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
