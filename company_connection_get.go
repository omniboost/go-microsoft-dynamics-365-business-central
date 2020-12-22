package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCompanyConnectionGetRequest() CompanyConnectionGetRequest {
	r := CompanyConnectionGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompanyConnectionGetRequest struct {
	client      *Client
	queryParams *CompanyConnectionGetQueryParams
	pathParams  *CompanyConnectionGetPathParams
	method      string
	headers     http.Header
	requestBody CompanyConnectionGetRequestBody
}

func (r CompanyConnectionGetRequest) NewQueryParams() *CompanyConnectionGetQueryParams {
	return &CompanyConnectionGetQueryParams{}
}

type CompanyConnectionGetQueryParams struct {
}

func (p CompanyConnectionGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CompanyConnectionGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CompanyConnectionGetRequest) NewPathParams() *CompanyConnectionGetPathParams {
	return &CompanyConnectionGetPathParams{}
}

type CompanyConnectionGetPathParams struct {
}

func (p *CompanyConnectionGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompanyConnectionGetRequest) PathParams() *CompanyConnectionGetPathParams {
	return r.pathParams
}

func (r *CompanyConnectionGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CompanyConnectionGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompanyConnectionGetRequest) Method() string {
	return r.method
}

func (r CompanyConnectionGetRequest) NewRequestBody() CompanyConnectionGetRequestBody {
	return CompanyConnectionGetRequestBody{}
}

type CompanyConnectionGetRequestBody struct {
}

func (r *CompanyConnectionGetRequest) RequestBody() *CompanyConnectionGetRequestBody {
	return &r.requestBody
}

func (r *CompanyConnectionGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CompanyConnectionGetRequest) SetRequestBody(body CompanyConnectionGetRequestBody) {
	r.requestBody = body
}

func (r *CompanyConnectionGetRequest) NewResponseBody() *CompanyConnectionGetResponseBody {
	var b = CompanyConnectionGetResponseBody(false)
	return &b
}

type CompanyConnectionGetResponseBody bool

func (r *CompanyConnectionGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("company/connection", r.PathParams())
	return &u
}

func (r *CompanyConnectionGetRequest) Do() (CompanyConnectionGetResponseBody, error) {
	// if r.QueryParams().ConsumerToken == "" {
	// 	r.QueryParams().ConsumerToken = r.client.ConsumerToken()
	// }

	// if r.QueryParams().EmployeeToken == "" {
	// 	r.QueryParams().EmployeeToken = r.client.EmployeeToken()
	// }

	// if r.QueryParams().ExpirationDate.IsZero() {
	// 	r.QueryParams().ExpirationDate = Date{time.Now().AddDate(0, 0, 1)}
	// }

	// fetch a new token if it isn't set already
	// if r.client.token == "" {
	// 	var err error
	// 	r.client.token, err = r.client.NewToken()
	// 	if err != nil {
	// 		return *r.NewResponseBody(), err
	// 	}
	// }

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
