package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerGetAll() CustomerGetAll {
	r := CustomerGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerGetAll struct {
	client      *Client
	queryParams *CustomerGetAllQueryParams
	pathParams  *CustomerGetAllPathParams
	method      string
	headers     http.Header
	requestBody CustomerGetAllBody
}

func (r CustomerGetAll) NewQueryParams() *CustomerGetAllQueryParams {
	return &CustomerGetAllQueryParams{}
}

type CustomerGetAllQueryParams struct {
	Name              string `schema:"name,omitempty"`
	VATRegistrationID string `schema:"vatRegistrationId,omitempty"`
	Email             string `schema:"email,omitempty"`
	Phone             string `schema:"phone,omitempty"`
}

func (p CustomerGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerGetAll) QueryParams() *CustomerGetAllQueryParams {
	return r.queryParams
}

func (r CustomerGetAll) NewPathParams() *CustomerGetAllPathParams {
	return &CustomerGetAllPathParams{}
}

type CustomerGetAllPathParams struct {
}

func (p *CustomerGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerGetAll) PathParams() *CustomerGetAllPathParams {
	return r.pathParams
}

func (r *CustomerGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerGetAll) SetMethod(method string) {
	r.method = method
}

func (r *CustomerGetAll) Method() string {
	return r.method
}

func (r CustomerGetAll) NewRequestBody() CustomerGetAllBody {
	return CustomerGetAllBody{}
}

type CustomerGetAllBody struct {
}

func (r *CustomerGetAll) RequestBody() *CustomerGetAllBody {
	return nil
}

func (r *CustomerGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerGetAll) SetRequestBody(body CustomerGetAllBody) {
	r.requestBody = body
}

func (r *CustomerGetAll) NewResponseBody() *CustomerGetAllResponseBody {
	return &CustomerGetAllResponseBody{}
}

type CustomerGetAllResponseBody Customers

func (r *CustomerGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customer", r.PathParams())
	return &u
}

func (r *CustomerGetAll) Do() (CustomerGetAllResponseBody, error) {
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
