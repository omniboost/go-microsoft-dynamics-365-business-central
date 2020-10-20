package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewCustomerGetRequest() CustomerGetRequest {
	r := CustomerGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerGetQueryParams()
	r.pathParams = r.NewCustomerGetPathParams()
	r.requestBody = r.NewCustomerGetRequestBody()
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

func (r CustomerGetRequest) NewCustomerGetQueryParams() *CustomerGetQueryParams {
	return &CustomerGetQueryParams{
		From:  0,
		Count: 1000,
	}
}

type CustomerGetQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p CustomerGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
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

func (r CustomerGetRequest) NewCustomerGetPathParams() *CustomerGetPathParams {
	return &CustomerGetPathParams{}
}

type CustomerGetPathParams struct {
}

func (p *CustomerGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerGetRequest) PathParams() *CustomerGetPathParams {
	return r.pathParams
}

func (r *CustomerGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerGetRequest) Method() string {
	return r.method
}

func (r CustomerGetRequest) NewCustomerGetRequestBody() CustomerGetRequestBody {
	return CustomerGetRequestBody{}
}

type CustomerGetRequestBody struct{}

func (r *CustomerGetRequest) RequestBody() *CustomerGetRequestBody {
	return &r.requestBody
}

func (r *CustomerGetRequest) SetRequestBody(body CustomerGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerGetRequest) NewResponseBody() *CustomerGetResponseBody {
	return &CustomerGetResponseBody{}
}

type CustomerGetResponseBody struct {
	FullResultSize int       `json:"fullResultSize"`
	From           int       `json:"from"`
	Count          int       `json:"count"`
	VersionDigest  string    `json:"versionDigest"`
	Values         Customers `json:"values"`
}

func (r *CustomerGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customer", r.PathParams())
}

func (r *CustomerGetRequest) Do() (CustomerGetResponseBody, error) {
	// fetch a new token if it isn't set already
	if r.client.token == "" {
		var err error
		r.client.token, err = r.client.NewToken()
		if err != nil {
			return *r.NewResponseBody(), err
		}
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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

func (r *CustomerGetRequest) All() (CustomerGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := CustomerGetResponseBody{}
	concat.Count = resp.Count
	concat.From = resp.From
	concat.FullResultSize = resp.FullResultSize
	concat.Values = resp.Values
	concat.VersionDigest = resp.VersionDigest

	for concat.From+concat.Count < concat.FullResultSize {
		r.QueryParams().From = r.QueryParams().From + r.QueryParams().Count
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		concat.Count = resp.Count
		concat.From = resp.From
		concat.FullResultSize = resp.FullResultSize
		concat.Values = append(concat.Values, resp.Values...)
		concat.VersionDigest = resp.VersionDigest
	}

	return concat, nil
}
