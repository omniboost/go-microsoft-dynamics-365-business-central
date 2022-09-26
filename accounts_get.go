package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/utils"
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
	return &AccountsGetQueryParams{}
}

type AccountsGetQueryParams struct {
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

func (r *AccountsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r AccountsGet) NewPathParams() *AccountsGetPathParams {
	return &AccountsGetPathParams{}
}

type AccountsGetPathParams struct {
}

func (p *AccountsGetPathParams) Params() map[string]string {
	return map[string]string{}
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
	Meta Meta     `json:"Meta"`
	Data Accounts `json:"Data"`
}

func (r *AccountsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/accounts", r.PathParams())
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
