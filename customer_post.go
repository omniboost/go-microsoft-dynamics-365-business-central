package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewCustomerPostRequest() CustomerPostRequest {
	r := CustomerPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerPostQueryParams()
	r.pathParams = r.NewCustomerPostPathParams()
	r.requestBody = r.NewCustomerPostRequestBody()
	return r
}

type CustomerPostRequest struct {
	client      *Client
	queryParams *CustomerPostQueryParams
	pathParams  *CustomerPostPathParams
	method      string
	headers     http.Header
	requestBody CustomerPostRequestBody
}

func (r CustomerPostRequest) NewCustomerPostQueryParams() *CustomerPostQueryParams {
	return &CustomerPostQueryParams{}
}

type CustomerPostQueryParams struct{}

func (p CustomerPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerPostRequest) QueryParams() *CustomerPostQueryParams {
	return r.queryParams
}

func (r CustomerPostRequest) NewCustomerPostPathParams() *CustomerPostPathParams {
	return &CustomerPostPathParams{}
}

type CustomerPostPathParams struct {
}

func (p *CustomerPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerPostRequest) PathParams() *CustomerPostPathParams {
	return r.pathParams
}

func (r *CustomerPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerPostRequest) Method() string {
	return r.method
}

func (r CustomerPostRequest) NewCustomerPostRequestBody() CustomerPostRequestBody {
	return CustomerPostRequestBody{}
}

type CustomerPostRequestBody Customer

func (r *CustomerPostRequest) RequestBody() *CustomerPostRequestBody {
	return &r.requestBody
}

func (r *CustomerPostRequest) SetRequestBody(body CustomerPostRequestBody) {
	r.requestBody = body
}

func (r *CustomerPostRequest) NewResponseBody() *CustomerPostResponseBody {
	return &CustomerPostResponseBody{}
}

type CustomerPostResponseBody struct {
	Value Customer `json:"value"`
}

func (r *CustomerPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customer", r.PathParams())
}

func (r *CustomerPostRequest) Do() (CustomerPostResponseBody, error) {
	// fetch a new token if it isn't set already
	if r.client.token == "" {
		var err error
		r.client.token, err = r.client.NewToken()
		if err != nil {
			return *r.NewResponseBody(), err
		}
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
