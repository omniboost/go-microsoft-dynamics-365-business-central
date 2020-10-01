package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerGetRequest() LedgerGetRequest {
	r := LedgerGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerGetQueryParams()
	r.pathParams = r.NewLedgerGetPathParams()
	r.requestBody = r.NewLedgerGetRequestBody()
	return r
}

type LedgerGetRequest struct {
	client      *Client
	queryParams *LedgerGetQueryParams
	pathParams  *LedgerGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerGetRequestBody
}

func (r LedgerGetRequest) NewLedgerGetQueryParams() *LedgerGetQueryParams {
	return &LedgerGetQueryParams{}
}

type LedgerGetQueryParams struct {
	DateFrom Date `schema:"dateFrom"`
	DateTo   Date `schema:"dateTo"`
}

func (p LedgerGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerGetRequest) QueryParams() *LedgerGetQueryParams {
	return r.queryParams
}

func (r LedgerGetRequest) NewLedgerGetPathParams() *LedgerGetPathParams {
	return &LedgerGetPathParams{}
}

type LedgerGetPathParams struct {
}

func (p *LedgerGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerGetRequest) PathParams() *LedgerGetPathParams {
	return r.pathParams
}

func (r *LedgerGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerGetRequest) Method() string {
	return r.method
}

func (r LedgerGetRequest) NewLedgerGetRequestBody() LedgerGetRequestBody {
	return LedgerGetRequestBody{}
}

type LedgerGetRequestBody struct{}

func (r *LedgerGetRequest) RequestBody() *LedgerGetRequestBody {
	return &r.requestBody
}

func (r *LedgerGetRequest) SetRequestBody(body LedgerGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerGetRequest) NewResponseBody() *LedgerGetResponseBody {
	return &LedgerGetResponseBody{}
}

type LedgerGetResponseBody struct {
}

func (r *LedgerGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger", r.PathParams())
}

func (r *LedgerGetRequest) Do() (LedgerGetResponseBody, error) {
	// fetch a new token if it isn't set already
	if r.client.token == "" {
		var err error
		r.client.token, err = r.client.NewToken()
		if err != nil {
			return *r.NewResponseBody(), err
		}
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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
