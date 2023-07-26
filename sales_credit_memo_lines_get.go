package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesCreditMemoLinesGet() SalesCreditMemoLinesGet {
	r := SalesCreditMemoLinesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesCreditMemoLinesGet struct {
	client      *Client
	queryParams *SalesCreditMemoLinesGetQueryParams
	pathParams  *SalesCreditMemoLinesGetPathParams
	method      string
	headers     http.Header
	requestBody SalesCreditMemoLinesGetBody
}

func (r SalesCreditMemoLinesGet) NewQueryParams() *SalesCreditMemoLinesGetQueryParams {
	selectFields, _ := utils.Fields(&SalesCreditMemoLine{})
	return &SalesCreditMemoLinesGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type SalesCreditMemoLinesGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p SalesCreditMemoLinesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesCreditMemoLinesGet) QueryParams() *SalesCreditMemoLinesGetQueryParams {
	return r.queryParams
}

func (r SalesCreditMemoLinesGet) NewPathParams() *SalesCreditMemoLinesGetPathParams {
	return &SalesCreditMemoLinesGetPathParams{}
}

type SalesCreditMemoLinesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	SalesCreditMemoID  string `schema:"salesCreditMemoID"`
}

func (p *SalesCreditMemoLinesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName":   p.EnvironmentName,
		"companyID":         p.CompanyID,
		"salesCreditMemoID": p.SalesCreditMemoID,
	}
}

func (r *SalesCreditMemoLinesGet) PathParams() *SalesCreditMemoLinesGetPathParams {
	return r.pathParams
}

func (r *SalesCreditMemoLinesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesCreditMemoLinesGet) SetMethod(method string) {
	r.method = method
}

func (r *SalesCreditMemoLinesGet) Method() string {
	return r.method
}

func (r *SalesCreditMemoLinesGet) Headers() http.Header {
	return r.headers
}

func (r SalesCreditMemoLinesGet) NewRequestBody() SalesCreditMemoLinesGetBody {
	return SalesCreditMemoLinesGetBody{}
}

type SalesCreditMemoLinesGetBody struct {
}

func (r *SalesCreditMemoLinesGet) RequestBody() *SalesCreditMemoLinesGetBody {
	return nil
}

func (r *SalesCreditMemoLinesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesCreditMemoLinesGet) SetRequestBody(body SalesCreditMemoLinesGetBody) {
	r.requestBody = body
}

func (r *SalesCreditMemoLinesGet) NewResponseBody() *SalesCreditMemoLinesGetResponseBody {
	return &SalesCreditMemoLinesGetResponseBody{}
}

type SalesCreditMemoLinesGetResponseBody struct {
	OdataContext string               `json:"@odata.context"`
	Value        SalesCreditMemoLines `json:"value"`
}

func (r *SalesCreditMemoLinesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesCreditMemos({{.salesCreditMemoID}})/salesCreditMemoLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesCreditMemoLinesGet) Do() (SalesCreditMemoLinesGetResponseBody, error) {
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
