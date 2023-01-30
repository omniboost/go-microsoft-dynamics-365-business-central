package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewAccountsGet() AccountsGet {
	r := AccountsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountsGet struct {
	client      *Client
	queryParams *AccountsGetQueryParams
	pathParams  *AccountsGetPathParams
	method      string
	headers     http.Header
	requestBody AccountsGetBody
}

func (r AccountsGet) NewQueryParams() *AccountsGetQueryParams {
	selectFields, _ := utils.Fields(&Account{})
	return &AccountsGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type AccountsGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p AccountsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountsGet) QueryParams() *AccountsGetQueryParams {
	return r.queryParams
}

func (r AccountsGet) NewPathParams() *AccountsGetPathParams {
	return &AccountsGetPathParams{}
}

type AccountsGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *AccountsGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *AccountsGet) PathParams() *AccountsGetPathParams {
	return r.pathParams
}

func (r *AccountsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountsGet) SetMethod(method string) {
	r.method = method
}

func (r *AccountsGet) Method() string {
	return r.method
}

func (r AccountsGet) NewRequestBody() AccountsGetBody {
	return AccountsGetBody{}
}

type AccountsGetBody struct {
}

func (r *AccountsGet) RequestBody() *AccountsGetBody {
	return nil
}

func (r *AccountsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountsGet) SetRequestBody(body AccountsGetBody) {
	r.requestBody = body
}

func (r *AccountsGet) NewResponseBody() *AccountsGetResponseBody {
	return &AccountsGetResponseBody{}
}

type AccountsGetResponseBody struct {
	OdataContext string   `json:"@odata.context"`
	Value        Accounts `json:"value"`
}

func (r *AccountsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/accounts", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *AccountsGet) Do() (AccountsGetResponseBody, error) {
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
