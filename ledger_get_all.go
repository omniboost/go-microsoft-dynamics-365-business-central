package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewLedgerGetAll() LedgerGetAll {
	r := LedgerGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LedgerGetAll struct {
	client      *Client
	queryParams *LedgerGetAllQueryParams
	pathParams  *LedgerGetAllPathParams
	method      string
	headers     http.Header
	requestBody LedgerGetAllBody
}

func (r LedgerGetAll) NewQueryParams() *LedgerGetAllQueryParams {
	return &LedgerGetAllQueryParams{}
}

type LedgerGetAllQueryParams struct{}

func (p LedgerGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *LedgerGetAll) QueryParams() *LedgerGetAllQueryParams {
	return r.queryParams
}

func (r LedgerGetAll) NewPathParams() *LedgerGetAllPathParams {
	return &LedgerGetAllPathParams{}
}

type LedgerGetAllPathParams struct {
}

func (p *LedgerGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerGetAll) PathParams() *LedgerGetAllPathParams {
	return r.pathParams
}

func (r *LedgerGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *LedgerGetAll) SetMethod(method string) {
	r.method = method
}

func (r *LedgerGetAll) Method() string {
	return r.method
}

func (r LedgerGetAll) NewRequestBody() LedgerGetAllBody {
	return LedgerGetAllBody{}
}

type LedgerGetAllBody struct {
}

func (r *LedgerGetAll) RequestBody() *LedgerGetAllBody {
	return nil
}

func (r *LedgerGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *LedgerGetAll) SetRequestBody(body LedgerGetAllBody) {
	r.requestBody = body
}

func (r *LedgerGetAll) NewResponseBody() *LedgerGetAllResponseBody {
	return &LedgerGetAllResponseBody{}
}

type LedgerGetAllResponseBody Ledgers

func (r *LedgerGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/ledger", r.PathParams())
	return &u
}

func (r *LedgerGetAll) Do() (LedgerGetAllResponseBody, error) {
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
