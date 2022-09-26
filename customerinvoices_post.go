package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewCustomerinvoicesPost() CustomerinvoicesPost {
	r := CustomerinvoicesPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerinvoicesPost struct {
	client      *Client
	queryParams *CustomerinvoicesPostQueryParams
	pathParams  *CustomerinvoicesPostPathParams
	method      string
	headers     http.Header
	requestBody CustomerinvoicesPostBody
}

func (r CustomerinvoicesPost) NewQueryParams() *CustomerinvoicesPostQueryParams {
	return &CustomerinvoicesPostQueryParams{}
}

type CustomerinvoicesPostQueryParams struct {
}

func (p CustomerinvoicesPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerinvoicesPost) QueryParams() *CustomerinvoicesPostQueryParams {
	return r.queryParams
}

func (r CustomerinvoicesPost) NewPathParams() *CustomerinvoicesPostPathParams {
	return &CustomerinvoicesPostPathParams{}
}

type CustomerinvoicesPostPathParams struct {
}

func (p *CustomerinvoicesPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerinvoicesPost) PathParams() *CustomerinvoicesPostPathParams {
	return r.pathParams
}

func (r *CustomerinvoicesPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerinvoicesPost) SetMethod(method string) {
	r.method = method
}

func (r *CustomerinvoicesPost) Method() string {
	return r.method
}

func (r CustomerinvoicesPost) NewRequestBody() CustomerinvoicesPostBody {
	return CustomerinvoicesPostBody{}
}

type CustomerinvoicesPostBody CustomerInvoice

func (v CustomerinvoicesPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *CustomerinvoicesPost) RequestBody() *CustomerinvoicesPostBody {
	return &r.requestBody
}

func (r *CustomerinvoicesPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerinvoicesPost) SetRequestBody(body CustomerinvoicesPostBody) {
	r.requestBody = body
}

func (r *CustomerinvoicesPost) NewResponseBody() *CustomerinvoicesPostResponseBody {
	return &CustomerinvoicesPostResponseBody{}
}

type CustomerinvoicesPostResponseBody CustomerInvoice

func (r *CustomerinvoicesPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/customerinvoices", r.PathParams())
	return &u
}

func (r *CustomerinvoicesPost) Do() (CustomerinvoicesPostResponseBody, error) {
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
