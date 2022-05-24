package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-vismaonline/omitempty"
	"github.com/omniboost/go-vismaonline/utils"
)

func (c *Client) NewCustomersPost() CustomersPost {
	r := CustomersPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomersPost struct {
	client      *Client
	queryParams *CustomersPostQueryParams
	pathParams  *CustomersPostPathParams
	method      string
	headers     http.Header
	requestBody CustomersPostBody
}

func (r CustomersPost) NewQueryParams() *CustomersPostQueryParams {
	return &CustomersPostQueryParams{}
}

type CustomersPostQueryParams struct {
}

func (p CustomersPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomersPost) QueryParams() *CustomersPostQueryParams {
	return r.queryParams
}

func (r CustomersPost) NewPathParams() *CustomersPostPathParams {
	return &CustomersPostPathParams{}
}

type CustomersPostPathParams struct {
}

func (p *CustomersPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomersPost) PathParams() *CustomersPostPathParams {
	return r.pathParams
}

func (r *CustomersPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomersPost) SetMethod(method string) {
	r.method = method
}

func (r *CustomersPost) Method() string {
	return r.method
}

func (r CustomersPost) NewRequestBody() CustomersPostBody {
	return CustomersPostBody{}
}

type CustomersPostBody Customer

func (v CustomersPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *CustomersPost) RequestBody() *CustomersPostBody {
	return &r.requestBody
}

func (r *CustomersPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomersPost) SetRequestBody(body CustomersPostBody) {
	r.requestBody = body
}

func (r *CustomersPost) NewResponseBody() *CustomersPostResponseBody {
	return &CustomersPostResponseBody{}
}

type CustomersPostResponseBody Customer

func (r *CustomersPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/customers", r.PathParams())
	return &u
}

func (r *CustomersPost) Do() (CustomersPostResponseBody, error) {
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
