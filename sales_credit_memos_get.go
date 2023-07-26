package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesCreditMemosGet() SalesCreditMemosGet {
	r := SalesCreditMemosGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesCreditMemosGet struct {
	client      *Client
	queryParams *SalesCreditMemosGetQueryParams
	pathParams  *SalesCreditMemosGetPathParams
	method      string
	headers     http.Header
	requestBody SalesCreditMemosGetBody
}

func (r SalesCreditMemosGet) NewQueryParams() *SalesCreditMemosGetQueryParams {
	selectFields, _ := utils.Fields(&SalesCreditMemo{})
	return &SalesCreditMemosGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type SalesCreditMemosGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p SalesCreditMemosGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesCreditMemosGet) QueryParams() *SalesCreditMemosGetQueryParams {
	return r.queryParams
}

func (r SalesCreditMemosGet) NewPathParams() *SalesCreditMemosGetPathParams {
	return &SalesCreditMemosGetPathParams{}
}

type SalesCreditMemosGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *SalesCreditMemosGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *SalesCreditMemosGet) PathParams() *SalesCreditMemosGetPathParams {
	return r.pathParams
}

func (r *SalesCreditMemosGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesCreditMemosGet) SetMethod(method string) {
	r.method = method
}

func (r *SalesCreditMemosGet) Method() string {
	return r.method
}

func (r *SalesCreditMemosGet) Headers() http.Header {
	return r.headers
}

func (r SalesCreditMemosGet) NewRequestBody() SalesCreditMemosGetBody {
	return SalesCreditMemosGetBody{}
}

type SalesCreditMemosGetBody struct {
}

func (r *SalesCreditMemosGet) RequestBody() *SalesCreditMemosGetBody {
	return nil
}

func (r *SalesCreditMemosGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesCreditMemosGet) SetRequestBody(body SalesCreditMemosGetBody) {
	r.requestBody = body
}

func (r *SalesCreditMemosGet) NewResponseBody() *SalesCreditMemosGetResponseBody {
	return &SalesCreditMemosGetResponseBody{}
}

type SalesCreditMemosGetResponseBody struct {
	OdataContext string           `json:"@odata.context"`
	Value        SalesCreditMemos `json:"value"`
}

func (r *SalesCreditMemosGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesCreditMemos", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesCreditMemosGet) Do() (SalesCreditMemosGetResponseBody, error) {
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
