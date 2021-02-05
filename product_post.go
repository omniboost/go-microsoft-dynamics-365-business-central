package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewProductPostRequest() ProductPostRequest {
	r := ProductPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ProductPostRequest struct {
	client      *Client
	queryParams *ProductPostQueryParams
	pathParams  *ProductPostPathParams
	method      string
	headers     http.Header
	requestBody ProductPostRequestBody
}

func (r ProductPostRequest) NewQueryParams() *ProductPostQueryParams {
	return &ProductPostQueryParams{}
}

type ProductPostQueryParams struct {
	Post bool `schema:"post"`
}

func (p ProductPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ProductPostRequest) QueryParams() *ProductPostQueryParams {
	return r.queryParams
}

func (r ProductPostRequest) NewPathParams() *ProductPostPathParams {
	return &ProductPostPathParams{}
}

type ProductPostPathParams struct {
}

func (p *ProductPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProductPostRequest) PathParams() *ProductPostPathParams {
	return r.pathParams
}

func (r *ProductPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ProductPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProductPostRequest) Method() string {
	return r.method
}

func (r ProductPostRequest) NewRequestBody() ProductPostRequestBody {
	return ProductPostRequestBody{}
}

type ProductPostRequestBody struct {
	ItemCode               string  `json:"ItemCode"`
	AliasItemCode          string  `json:"AliasItemCode,omitempty"`
	Group                  string  `json:"Group,omitempty"`
	Description            string  `json:"Description,omitempty"`
	Description2           string  `json:"Description2,omitempty"`
	Price1                 float64 `json:"Price1,omitempty"`
	Price2                 float64 `json:"Price2,omitempty"`
	Price3                 float64 `json:"Price3,omitempty"`
	CostPrice              float64 `json:"CostPrice,omitempty"`
	TaxPercent             float64 `json:"TaxPercent"`
	Inactive               bool    `json:"Inactive,omitempty"`
	PurchacePrice          float64 `json:"PurchasePrice,omitempty"`
	AllowDiscount          bool    `json:"AllowDiscount,omitempty"`
	Discount               float64 `json:"Discount,omitempty"`
	ExtraDesc1             string  `json:"ExtraDesc1,omitempty"`
	ExtraDesc2             string  `json:"ExtraDesc2,omitempty"`
	AllowNegativeInventory bool    `json:"AllowNegativeInventory"`
	ShowItemInWebShop      bool    `json:"ShowItemInWebShop"`
	SalesLedgerCode        string  `json:"SalesLedgerCode"`
}

func (r *ProductPostRequest) RequestBody() *ProductPostRequestBody {
	return &r.requestBody
}

func (r *ProductPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ProductPostRequest) SetRequestBody(body ProductPostRequestBody) {
	r.requestBody = body
}

func (r *ProductPostRequest) NewResponseBody() *ProductPostResponseBody {
	return &ProductPostResponseBody{}
}

type ProductPostResponseBody Product

func (r *ProductPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("product", r.PathParams())
	return &u
}

func (r *ProductPostRequest) Do() (ProductPostResponseBody, error) {
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
