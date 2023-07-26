package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewCustomerPost() CustomerPost {
	r := CustomerPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerPost struct {
	client      *Client
	queryParams *CustomerPostQueryParams
	pathParams  *CustomerPostPathParams
	method      string
	headers     http.Header
	requestBody CustomerPostBody
}

func (r CustomerPost) NewQueryParams() *CustomerPostQueryParams {
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

func (r *CustomerPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomerPost) NewPathParams() *CustomerPostPathParams {
	return &CustomerPostPathParams{}
}

type CustomerPostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *CustomerPostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *CustomerPost) PathParams() *CustomerPostPathParams {
	return r.pathParams
}

func (r *CustomerPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerPost) SetMethod(method string) {
	r.method = method
}

func (r *CustomerPost) Method() string {
	return r.method
}

func (r *CustomerPost) Headers() http.Header {
	return r.headers
}

func (r CustomerPost) NewRequestBody() CustomerPostBody {
	return CustomerPostBody{}
}

type CustomerPostBody struct {
	Customer
}

func (r *CustomerPost) RequestBody() *CustomerPostBody {
	return &r.requestBody
}

func (r *CustomerPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *CustomerPost) SetRequestBody(body CustomerPostBody) {
	r.requestBody = body
}

func (r *CustomerPost) NewResponseBody() *CustomerPostResponseBody {
	return &CustomerPostResponseBody{}
}

type CustomerPostResponseBody struct {
	OdataContext string       `json:"@odata.context"`
	Value        Customer `json:"value"`
}

func (r *CustomerPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/customers", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CustomerPost) Do() (CustomerPostResponseBody, error) {
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


