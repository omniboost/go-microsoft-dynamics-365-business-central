package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCustomerDeleteRequest() CustomerDeleteRequest {
	r := CustomerDeleteRequest{
		client:  c,
		method:  http.MethodDelete,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerDeleteRequest struct {
	client      *Client
	queryParams *CustomerDeleteQueryParams
	pathParams  *CustomerDeletePathParams
	method      string
	headers     http.Header
	requestBody CustomerDeleteRequestBody
}

func (r CustomerDeleteRequest) NewQueryParams() *CustomerDeleteQueryParams {
	return &CustomerDeleteQueryParams{}
}

type CustomerDeleteQueryParams struct {
}

func (p CustomerDeleteQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerDeleteRequest) QueryParams() *CustomerDeleteQueryParams {
	return r.queryParams
}

func (r CustomerDeleteRequest) NewPathParams() *CustomerDeletePathParams {
	return &CustomerDeletePathParams{}
}

type CustomerDeletePathParams struct {
	Number string `schema:"number"`
}

func (p *CustomerDeletePathParams) Params() map[string]string {
	return map[string]string{
		"number": p.Number,
	}
}

func (r *CustomerDeleteRequest) PathParams() *CustomerDeletePathParams {
	return r.pathParams
}

func (r *CustomerDeleteRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerDeleteRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerDeleteRequest) Method() string {
	return r.method
}

func (r CustomerDeleteRequest) NewRequestBody() CustomerDeleteRequestBody {
	return CustomerDeleteRequestBody{}
}

type CustomerDeleteRequestBody struct {
}

func (r *CustomerDeleteRequest) RequestBody() *CustomerDeleteRequestBody {
	return &r.requestBody
}

func (r *CustomerDeleteRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerDeleteRequest) SetRequestBody(body CustomerDeleteRequestBody) {
	r.requestBody = body
}

func (r *CustomerDeleteRequest) NewResponseBody() *CustomerDeleteResponseBody {
	return &CustomerDeleteResponseBody{}
}

type CustomerDeleteResponseBody []struct {
}

func (r *CustomerDeleteRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("customer/{{.number}}", r.PathParams())
	return &u
}

func (r *CustomerDeleteRequest) Do() (CustomerDeleteResponseBody, error) {
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
