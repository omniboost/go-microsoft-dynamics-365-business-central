package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewSalesCreditMemoPost() SalesCreditMemoPost {
	r := SalesCreditMemoPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesCreditMemoPost struct {
	client      *Client
	queryParams *SalesCreditMemoPostQueryParams
	pathParams  *SalesCreditMemoPostPathParams
	method      string
	headers     http.Header
	requestBody SalesCreditMemoPostBody
}

func (r SalesCreditMemoPost) NewQueryParams() *SalesCreditMemoPostQueryParams {
	return &SalesCreditMemoPostQueryParams{}
}

type SalesCreditMemoPostQueryParams struct {
}

func (p SalesCreditMemoPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesCreditMemoPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r SalesCreditMemoPost) NewPathParams() *SalesCreditMemoPostPathParams {
	return &SalesCreditMemoPostPathParams{}
}

type SalesCreditMemoPostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *SalesCreditMemoPostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *SalesCreditMemoPost) PathParams() *SalesCreditMemoPostPathParams {
	return r.pathParams
}

func (r *SalesCreditMemoPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesCreditMemoPost) SetMethod(method string) {
	r.method = method
}

func (r *SalesCreditMemoPost) Method() string {
	return r.method
}

func (r *SalesCreditMemoPost) Headers() http.Header {
	return r.headers
}

func (r SalesCreditMemoPost) NewRequestBody() SalesCreditMemoPostBody {
	return SalesCreditMemoPostBody{}
}

type SalesCreditMemoPostBody struct {
	SalesCreditMemo
}

func (r *SalesCreditMemoPost) RequestBody() *SalesCreditMemoPostBody {
	return &r.requestBody
}

func (r *SalesCreditMemoPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *SalesCreditMemoPost) SetRequestBody(body SalesCreditMemoPostBody) {
	r.requestBody = body
}

func (r *SalesCreditMemoPost) NewResponseBody() *SalesCreditMemoPostResponseBody {
	return &SalesCreditMemoPostResponseBody{}
}

type SalesCreditMemoPostResponseBody struct {
	OdataContext string           `json:"@odata.context"`
	Value        SalesCreditMemos `json:"value"`
}

func (r *SalesCreditMemoPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/salesCreditMemos", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *SalesCreditMemoPost) Do() (SalesCreditMemoPostResponseBody, error) {
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
