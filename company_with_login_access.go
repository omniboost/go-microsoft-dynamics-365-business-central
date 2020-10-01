package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewCompanyWithLoginAccessRequest() CompanyWithLoginAccessRequest {
	r := CompanyWithLoginAccessRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCompanyWithLoginAccessQueryParams()
	r.pathParams = r.NewCompanyWithLoginAccessPathParams()
	r.requestBody = r.NewCompanyWithLoginAccessRequestBody()
	return r
}

type CompanyWithLoginAccessRequest struct {
	client      *Client
	queryParams *CompanyWithLoginAccessQueryParams
	pathParams  *CompanyWithLoginAccessPathParams
	method      string
	headers     http.Header
	requestBody CompanyWithLoginAccessRequestBody
}

func (r CompanyWithLoginAccessRequest) NewCompanyWithLoginAccessQueryParams() *CompanyWithLoginAccessQueryParams {
	return &CompanyWithLoginAccessQueryParams{}
}

type CompanyWithLoginAccessQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p CompanyWithLoginAccessQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CompanyWithLoginAccessRequest) QueryParams() *CompanyWithLoginAccessQueryParams {
	return r.queryParams
}

func (r CompanyWithLoginAccessRequest) NewCompanyWithLoginAccessPathParams() *CompanyWithLoginAccessPathParams {
	return &CompanyWithLoginAccessPathParams{}
}

type CompanyWithLoginAccessPathParams struct {
}

func (p *CompanyWithLoginAccessPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompanyWithLoginAccessRequest) PathParams() *CompanyWithLoginAccessPathParams {
	return r.pathParams
}

func (r *CompanyWithLoginAccessRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompanyWithLoginAccessRequest) Method() string {
	return r.method
}

func (r CompanyWithLoginAccessRequest) NewCompanyWithLoginAccessRequestBody() CompanyWithLoginAccessRequestBody {
	return CompanyWithLoginAccessRequestBody{}
}

type CompanyWithLoginAccessRequestBody struct{}

func (r *CompanyWithLoginAccessRequest) RequestBody() *CompanyWithLoginAccessRequestBody {
	return &r.requestBody
}

func (r *CompanyWithLoginAccessRequest) SetRequestBody(body CompanyWithLoginAccessRequestBody) {
	r.requestBody = body
}

func (r *CompanyWithLoginAccessRequest) NewResponseBody() *CompanyWithLoginAccessResponseBody {
	return &CompanyWithLoginAccessResponseBody{}
}

type CompanyWithLoginAccessResponseBody struct {
	FullResultSize int           `json:"fullResultSize"`
	From           int           `json:"from"`
	Count          int           `json:"count"`
	VersionDigest  interface{}   `json:"versionDigest"`
	Values         []interface{} `json:"values"`
}

func (r *CompanyWithLoginAccessRequest) URL() url.URL {
	return r.client.GetEndpointURL("/company/>withLoginAccess", r.PathParams())
}

func (r *CompanyWithLoginAccessRequest) Do() (CompanyWithLoginAccessResponseBody, error) {
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
