package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesInvoicesGet() SalesInvoicesGet {
	r := SalesInvoicesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicesGet struct {
	client      *Client
	queryParams *SalesInvoicesGetQueryParams
	pathParams  *SalesInvoicesGetPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoicesGetBody
}

func (r SalesInvoicesGet) NewQueryParams() *SalesInvoicesGetQueryParams {
	selectFields, _ := utils.Fields(&SalesInvoice{})
	return &SalesInvoicesGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type SalesInvoicesGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p SalesInvoicesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoicesGet) QueryParams() *SalesInvoicesGetQueryParams {
	return r.queryParams
}

func (r SalesInvoicesGet) NewPathParams() *SalesInvoicesGetPathParams {
	return &SalesInvoicesGetPathParams{}
}

type SalesInvoicesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *SalesInvoicesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *SalesInvoicesGet) PathParams() *SalesInvoicesGetPathParams {
	return r.pathParams
}

func (r *SalesInvoicesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicesGet) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicesGet) Method() string {
	return r.method
}

func (r *SalesInvoicesGet) Headers() http.Header {
	return r.headers
}

func (r SalesInvoicesGet) NewRequestBody() SalesInvoicesGetBody {
	return SalesInvoicesGetBody{}
}

type SalesInvoicesGetBody struct {
}

func (r *SalesInvoicesGet) RequestBody() *SalesInvoicesGetBody {
	return nil
}

func (r *SalesInvoicesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesInvoicesGet) SetRequestBody(body SalesInvoicesGetBody) {
	r.requestBody = body
}

func (r *SalesInvoicesGet) NewResponseBody() *SalesInvoicesGetResponseBody {
	return &SalesInvoicesGetResponseBody{}
}

type SalesInvoicesGetResponseBody struct {
	OdataContext string       `json:"@odata.context"`
	Value        SalesInvoices `json:"value"`
}

func (r *SalesInvoicesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesInvoices", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesInvoicesGet) Do() (SalesInvoicesGetResponseBody, error) {
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

