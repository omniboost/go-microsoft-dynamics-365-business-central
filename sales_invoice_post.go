package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesInvoicePost() SalesInvoicePost {
	r := SalesInvoicePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicePost struct {
	client      *Client
	queryParams *SalesInvoicePostQueryParams
	pathParams  *SalesInvoicePostPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoicePostBody
}

func (r SalesInvoicePost) NewQueryParams() *SalesInvoicePostQueryParams {
	return &SalesInvoicePostQueryParams{}
}

type SalesInvoicePostQueryParams struct {
}

func (p SalesInvoicePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoicePost) QueryParams() QueryParams {
	return r.queryParams
}

func (r SalesInvoicePost) NewPathParams() *SalesInvoicePostPathParams {
	return &SalesInvoicePostPathParams{}
}

type SalesInvoicePostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *SalesInvoicePostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *SalesInvoicePost) PathParams() *SalesInvoicePostPathParams {
	return r.pathParams
}

func (r *SalesInvoicePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicePost) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicePost) Method() string {
	return r.method
}

func (r *SalesInvoicePost) Headers() http.Header {
	return r.headers
}

func (r SalesInvoicePost) NewRequestBody() SalesInvoicePostBody {
	return SalesInvoicePostBody{}
}

type SalesInvoicePostBody struct {
	SalesInvoice
}

func (r *SalesInvoicePost) RequestBody() *SalesInvoicePostBody {
	return &r.requestBody
}

func (r *SalesInvoicePost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *SalesInvoicePost) SetRequestBody(body SalesInvoicePostBody) {
	r.requestBody = body
}

func (r *SalesInvoicePost) NewResponseBody() *SalesInvoicePostResponseBody {
	return &SalesInvoicePostResponseBody{}
}

type SalesInvoicePostResponseBody struct {
	OdataContext string       `json:"@odata.context"`
	Value        SalesInvoices `json:"value"`
}

func (r *SalesInvoicePost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesInvoices", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesInvoicePost) Do() (SalesInvoicePostResponseBody, error) {
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

