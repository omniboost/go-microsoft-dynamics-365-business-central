package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesCreditMemoLinePost() SalesCreditMemoLinePost {
	r := SalesCreditMemoLinePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesCreditMemoLinePost struct {
	client      *Client
	queryParams *SalesCreditMemoLinePostQueryParams
	pathParams  *SalesCreditMemoLinePostPathParams
	method      string
	headers     http.Header
	requestBody SalesCreditMemoLinePostBody
}

func (r SalesCreditMemoLinePost) NewQueryParams() *SalesCreditMemoLinePostQueryParams {
	return &SalesCreditMemoLinePostQueryParams{}
}

type SalesCreditMemoLinePostQueryParams struct {
}

func (p SalesCreditMemoLinePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesCreditMemoLinePost) QueryParams() QueryParams {
	return r.queryParams
}

func (r SalesCreditMemoLinePost) NewPathParams() *SalesCreditMemoLinePostPathParams {
	return &SalesCreditMemoLinePostPathParams{}
}

type SalesCreditMemoLinePostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	SalesCreditMemoID  string `schema:"salesCreditMemoID"`
}

func (p *SalesCreditMemoLinePostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"salesCreditMemoID":  p.SalesCreditMemoID,
	}
}

func (r *SalesCreditMemoLinePost) PathParams() *SalesCreditMemoLinePostPathParams {
	return r.pathParams
}

func (r *SalesCreditMemoLinePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesCreditMemoLinePost) SetMethod(method string) {
	r.method = method
}

func (r *SalesCreditMemoLinePost) Method() string {
	return r.method
}

func (r *SalesCreditMemoLinePost) Headers() http.Header {
	return r.headers
}

func (r SalesCreditMemoLinePost) NewRequestBody() SalesCreditMemoLinePostBody {
	return SalesCreditMemoLinePostBody{}
}

type SalesCreditMemoLinePostBody struct {
	SalesCreditMemoLine
}

func (r *SalesCreditMemoLinePost) RequestBody() *SalesCreditMemoLinePostBody {
	return &r.requestBody
}

func (r *SalesCreditMemoLinePost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *SalesCreditMemoLinePost) SetRequestBody(body SalesCreditMemoLinePostBody) {
	r.requestBody = body
}

func (r *SalesCreditMemoLinePost) NewResponseBody() *SalesCreditMemoLinePostResponseBody {
	return &SalesCreditMemoLinePostResponseBody{}
}

type SalesCreditMemoLinePostResponseBody struct {
	OdataContext string           `json:"@odata.context"`
	Value        SalesCreditMemoLine `json:"value"`
}

func (r *SalesCreditMemoLinePost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesCreditMemos({{.salesCreditMemoID}})/salesCreditMemoLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesCreditMemoLinePost) Do() (SalesCreditMemoLinePostResponseBody, error) {
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

