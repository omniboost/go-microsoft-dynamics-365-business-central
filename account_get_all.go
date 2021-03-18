package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewAccountGetAll() AccountGetAll {
	r := AccountGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountGetAll struct {
	client      *Client
	queryParams *AccountGetAllQueryParams
	pathParams  *AccountGetAllPathParams
	method      string
	headers     http.Header
	requestBody AccountGetAllBody
}

func (r AccountGetAll) NewQueryParams() *AccountGetAllQueryParams {
	return &AccountGetAllQueryParams{}
}

type AccountGetAllQueryParams struct {
}

func (p AccountGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountGetAll) QueryParams() QueryParams {
	return r.queryParams
}

func (r AccountGetAll) NewPathParams() *AccountGetAllPathParams {
	return &AccountGetAllPathParams{}
}

type AccountGetAllPathParams struct {
}

func (p *AccountGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountGetAll) PathParams() *AccountGetAllPathParams {
	return r.pathParams
}

func (r *AccountGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountGetAll) SetMethod(method string) {
	r.method = method
}

func (r *AccountGetAll) Method() string {
	return r.method
}

func (r AccountGetAll) NewRequestBody() AccountGetAllBody {
	return AccountGetAllBody{}
}

type AccountGetAllBody struct {
}

func (r *AccountGetAll) RequestBody() *AccountGetAllBody {
	return nil
}

func (r *AccountGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountGetAll) SetRequestBody(body AccountGetAllBody) {
	r.requestBody = body
}

func (r *AccountGetAll) NewResponseBody() *AccountGetAllResponseBody {
	return &AccountGetAllResponseBody{}
}

type AccountGetAllResponseBody Accounts

func (r *AccountGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("controller/api/v1/account", r.PathParams())
	return &u
}

func (r *AccountGetAll) Do() (AccountGetAllResponseBody, error) {
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
