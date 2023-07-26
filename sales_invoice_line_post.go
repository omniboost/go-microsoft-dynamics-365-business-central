package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesInvoiceLinePost() SalesInvoiceLinePost {
	r := SalesInvoiceLinePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoiceLinePost struct {
	client      *Client
	queryParams *SalesInvoiceLinePostQueryParams
	pathParams  *SalesInvoiceLinePostPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoiceLinePostBody
}

func (r SalesInvoiceLinePost) NewQueryParams() *SalesInvoiceLinePostQueryParams {
	return &SalesInvoiceLinePostQueryParams{}
}

type SalesInvoiceLinePostQueryParams struct {
}

func (p SalesInvoiceLinePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoiceLinePost) QueryParams() QueryParams {
	return r.queryParams
}

func (r SalesInvoiceLinePost) NewPathParams() *SalesInvoiceLinePostPathParams {
	return &SalesInvoiceLinePostPathParams{}
}

type SalesInvoiceLinePostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	SalesInvoiceID  string `schema:"salesInvoiceID"`
}

func (p *SalesInvoiceLinePostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"salesInvoiceID":  p.SalesInvoiceID,
	}
}

func (r *SalesInvoiceLinePost) PathParams() *SalesInvoiceLinePostPathParams {
	return r.pathParams
}

func (r *SalesInvoiceLinePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoiceLinePost) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoiceLinePost) Method() string {
	return r.method
}

func (r *SalesInvoiceLinePost) Headers() http.Header {
	return r.headers
}

func (r SalesInvoiceLinePost) NewRequestBody() SalesInvoiceLinePostBody {
	return SalesInvoiceLinePostBody{}
}

type SalesInvoiceLinePostBody struct {
	SalesInvoiceLine
}

func (r *SalesInvoiceLinePost) RequestBody() *SalesInvoiceLinePostBody {
	return &r.requestBody
}

func (r *SalesInvoiceLinePost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *SalesInvoiceLinePost) SetRequestBody(body SalesInvoiceLinePostBody) {
	r.requestBody = body
}

func (r *SalesInvoiceLinePost) NewResponseBody() *SalesInvoiceLinePostResponseBody {
	return &SalesInvoiceLinePostResponseBody{}
}

type SalesInvoiceLinePostResponseBody struct {
	OdataContext string           `json:"@odata.context"`
	Value        SalesInvoiceLine `json:"value"`
}

func (r *SalesInvoiceLinePost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesInvoices({{.salesInvoiceID}})/salesInvoiceLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesInvoiceLinePost) Do() (SalesInvoiceLinePostResponseBody, error) {
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
