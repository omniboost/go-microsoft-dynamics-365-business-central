package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-vismaonline/utils"
)

func (c *Client) NewCustomerinvoicesGet() CustomerinvoicesGet {
	r := CustomerinvoicesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerinvoicesGet struct {
	client      *Client
	queryParams *CustomerinvoicesGetQueryParams
	pathParams  *CustomerinvoicesGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerinvoicesGetBody
}

func (r CustomerinvoicesGet) NewQueryParams() *CustomerinvoicesGetQueryParams {
	return &CustomerinvoicesGetQueryParams{}
}

type CustomerinvoicesGetQueryParams struct {
}

func (p CustomerinvoicesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerinvoicesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomerinvoicesGet) NewPathParams() *CustomerinvoicesGetPathParams {
	return &CustomerinvoicesGetPathParams{}
}

type CustomerinvoicesGetPathParams struct {
}

func (p *CustomerinvoicesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerinvoicesGet) PathParams() *CustomerinvoicesGetPathParams {
	return r.pathParams
}

func (r *CustomerinvoicesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerinvoicesGet) SetMethod(method string) {
	r.method = method
}

func (r *CustomerinvoicesGet) Method() string {
	return r.method
}

func (r CustomerinvoicesGet) NewRequestBody() CustomerinvoicesGetBody {
	return CustomerinvoicesGetBody{}
}

type CustomerinvoicesGetBody struct {
}

func (r *CustomerinvoicesGet) RequestBody() *CustomerinvoicesGetBody {
	return nil
}

func (r *CustomerinvoicesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerinvoicesGet) SetRequestBody(body CustomerinvoicesGetBody) {
	r.requestBody = body
}

func (r *CustomerinvoicesGet) NewResponseBody() *CustomerinvoicesGetResponseBody {
	return &CustomerinvoicesGetResponseBody{}
}

type CustomerinvoicesGetResponseBody struct {
	Meta Meta             `json:"Meta"`
	Data CustomerInvoices `json:"Data"`
}

func (r *CustomerinvoicesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/customerinvoices", r.PathParams())
	return &u
}

func (r *CustomerinvoicesGet) Do() (CustomerinvoicesGetResponseBody, error) {
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
