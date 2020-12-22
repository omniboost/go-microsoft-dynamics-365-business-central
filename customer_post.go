package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCustomerPostRequest() CustomerPostRequest {
	r := CustomerPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
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

func (r CustomerPostRequest) NewQueryParams() *CustomerPostQueryParams {
	return &CustomerPostQueryParams{}
}

type CustomerPostQueryParams struct {
}

func (p CustomerPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerPostRequest) QueryParams() *CustomerPostQueryParams {
	return r.queryParams
}

func (r CustomerPostRequest) NewPathParams() *CustomerPostPathParams {
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

func (r *CustomerPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerPostRequest) Method() string {
	return r.method
}

func (r CustomerPostRequest) NewRequestBody() CustomerPostRequestBody {
	return CustomerPostRequestBody{}
}

type CustomerPostRequestBody struct {
	Number      string `json:"Number"`
	Name        string `json:"Name"`
	Address1    string `json:"Address1"`
	Address2    string `json:"Address2"`
	Zipcode     string `json:"Zipcode"`
	Email       string `json:"Email"`
	Phone       string `json:"Phone"`
	PhoneLocal  string `json:"PhoneLocal"`
	PhoneMobile string `json:"PhoneMobile"`
}

func (r *CustomerPostRequest) RequestBody() *CustomerPostRequestBody {
	return &r.requestBody
}

func (r *CustomerPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerPostRequest) SetRequestBody(body CustomerPostRequestBody) {
	r.requestBody = body
}

func (r *CustomerPostRequest) NewResponseBody() *CustomerPostResponseBody {
	return &CustomerPostResponseBody{}
}

type CustomerPostResponseBody struct {
}

func (r *CustomerPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("customer", r.PathParams())
	return &u
}

func (r *CustomerPostRequest) Do() (CustomerPostResponseBody, error) {
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
