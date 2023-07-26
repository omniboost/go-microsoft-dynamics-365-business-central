package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesInvoiceLinesGet() SalesInvoiceLinesGet {
	r := SalesInvoiceLinesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoiceLinesGet struct {
	client      *Client
	queryParams *SalesInvoiceLinesGetQueryParams
	pathParams  *SalesInvoiceLinesGetPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoiceLinesGetBody
}

func (r SalesInvoiceLinesGet) NewQueryParams() *SalesInvoiceLinesGetQueryParams {
	selectFields, _ := utils.Fields(&SalesInvoiceLine{})
	return &SalesInvoiceLinesGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type SalesInvoiceLinesGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p SalesInvoiceLinesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoiceLinesGet) QueryParams() *SalesInvoiceLinesGetQueryParams {
	return r.queryParams
}

func (r SalesInvoiceLinesGet) NewPathParams() *SalesInvoiceLinesGetPathParams {
	return &SalesInvoiceLinesGetPathParams{}
}

type SalesInvoiceLinesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	SalesInvoiceID  string `schema:"salesInvoiceID"`
}

func (p *SalesInvoiceLinesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"salesInvoiceID":  p.SalesInvoiceID,
	}
}

func (r *SalesInvoiceLinesGet) PathParams() *SalesInvoiceLinesGetPathParams {
	return r.pathParams
}

func (r *SalesInvoiceLinesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoiceLinesGet) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoiceLinesGet) Method() string {
	return r.method
}

func (r *SalesInvoiceLinesGet) Headers() http.Header {
	return r.headers
}

func (r SalesInvoiceLinesGet) NewRequestBody() SalesInvoiceLinesGetBody {
	return SalesInvoiceLinesGetBody{}
}

type SalesInvoiceLinesGetBody struct {
}

func (r *SalesInvoiceLinesGet) RequestBody() *SalesInvoiceLinesGetBody {
	return nil
}

func (r *SalesInvoiceLinesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesInvoiceLinesGet) SetRequestBody(body SalesInvoiceLinesGetBody) {
	r.requestBody = body
}

func (r *SalesInvoiceLinesGet) NewResponseBody() *SalesInvoiceLinesGetResponseBody {
	return &SalesInvoiceLinesGetResponseBody{}
}

type SalesInvoiceLinesGetResponseBody struct {
	OdataContext string            `json:"@odata.context"`
	Value        SalesInvoiceLines `json:"value"`
}

func (r *SalesInvoiceLinesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesInvoices({{.salesInvoiceID}})/salesInvoiceLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesInvoiceLinesGet) Do() (SalesInvoiceLinesGetResponseBody, error) {
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
