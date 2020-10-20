package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerAccountGetRequest() LedgerAccountGetRequest {
	r := LedgerAccountGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerAccountGetQueryParams()
	r.pathParams = r.NewLedgerAccountGetPathParams()
	r.requestBody = r.NewLedgerAccountGetRequestBody()
	return r
}

type LedgerAccountGetRequest struct {
	client      *Client
	queryParams *LedgerAccountGetQueryParams
	pathParams  *LedgerAccountGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerAccountGetRequestBody
}

func (r LedgerAccountGetRequest) NewLedgerAccountGetQueryParams() *LedgerAccountGetQueryParams {
	return &LedgerAccountGetQueryParams{
		From:  0,
		Count: 1000,
	}
}

type LedgerAccountGetQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p LedgerAccountGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerAccountGetRequest) QueryParams() *LedgerAccountGetQueryParams {
	return r.queryParams
}

func (r LedgerAccountGetRequest) NewLedgerAccountGetPathParams() *LedgerAccountGetPathParams {
	return &LedgerAccountGetPathParams{}
}

type LedgerAccountGetPathParams struct {
}

func (p *LedgerAccountGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerAccountGetRequest) PathParams() *LedgerAccountGetPathParams {
	return r.pathParams
}

func (r *LedgerAccountGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerAccountGetRequest) Method() string {
	return r.method
}

func (r LedgerAccountGetRequest) NewLedgerAccountGetRequestBody() LedgerAccountGetRequestBody {
	return LedgerAccountGetRequestBody{}
}

type LedgerAccountGetRequestBody struct{}

func (r *LedgerAccountGetRequest) RequestBody() *LedgerAccountGetRequestBody {
	return &r.requestBody
}

func (r *LedgerAccountGetRequest) SetRequestBody(body LedgerAccountGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerAccountGetRequest) NewResponseBody() *LedgerAccountGetResponseBody {
	return &LedgerAccountGetResponseBody{}
}

type LedgerAccountGetResponseBody struct {
	FullResultSize int      `json:"fullResultSize"`
	From           int      `json:"from"`
	Count          int      `json:"count"`
	VersionDigest  string   `json:"versionDigest"`
	Values         Accounts `json:"values"`
}

func (r *LedgerAccountGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/account", r.PathParams())
}

func (r *LedgerAccountGetRequest) Do() (LedgerAccountGetResponseBody, error) {
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
