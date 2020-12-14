package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCompaniesGetRequest() CompaniesGetRequest {
	r := CompaniesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompaniesGetRequest struct {
	client      *Client
	queryParams *CompaniesGetQueryParams
	pathParams  *CompaniesGetPathParams
	method      string
	headers     http.Header
	requestBody CompaniesGetRequestBody
}

func (r CompaniesGetRequest) NewQueryParams() *CompaniesGetQueryParams {
	return &CompaniesGetQueryParams{}
}

type CompaniesGetQueryParams struct {
}

func (p CompaniesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CompaniesGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CompaniesGetRequest) NewPathParams() *CompaniesGetPathParams {
	return &CompaniesGetPathParams{}
}

type CompaniesGetPathParams struct {
}

func (p *CompaniesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompaniesGetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CompaniesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompaniesGetRequest) Method() string {
	return r.method
}

func (r CompaniesGetRequest) NewRequestBody() CompaniesGetRequestBody {
	return CompaniesGetRequestBody{}
}

type CompaniesGetRequestBody struct {
}

func (r *CompaniesGetRequest) RequestBody() *CompaniesGetRequestBody {
	return &r.requestBody
}

func (r *CompaniesGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CompaniesGetRequest) SetRequestBody(body CompaniesGetRequestBody) {
	r.requestBody = body
}

func (r *CompaniesGetRequest) NewResponseBody() *CompaniesGetResponseBody {
	return &CompaniesGetResponseBody{}
}

type CompaniesGetResponseBody struct {
}

func (r *CompaniesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("companies", r.PathParams())
	return &u
}

func (r *CompaniesGetRequest) Do() (CompaniesGetResponseBody, error) {
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
