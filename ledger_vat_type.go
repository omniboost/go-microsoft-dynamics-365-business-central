package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVatTypeRequest() LedgerVatTypeRequest {
	r := LedgerVatTypeRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVatTypeQueryParams()
	r.pathParams = r.NewLedgerVatTypePathParams()
	r.requestBody = r.NewLedgerVatTypeRequestBody()
	return r
}

type LedgerVatTypeRequest struct {
	client      *Client
	queryParams *LedgerVatTypeQueryParams
	pathParams  *LedgerVatTypePathParams
	method      string
	headers     http.Header
	requestBody LedgerVatTypeRequestBody
}

func (r LedgerVatTypeRequest) NewLedgerVatTypeQueryParams() *LedgerVatTypeQueryParams {
	return &LedgerVatTypeQueryParams{
		From:  0,
		Count: 1000,
	}
}

type LedgerVatTypeQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p LedgerVatTypeQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVatTypeRequest) QueryParams() *LedgerVatTypeQueryParams {
	return r.queryParams
}

func (r LedgerVatTypeRequest) NewLedgerVatTypePathParams() *LedgerVatTypePathParams {
	return &LedgerVatTypePathParams{}
}

type LedgerVatTypePathParams struct {
}

func (p *LedgerVatTypePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerVatTypeRequest) PathParams() *LedgerVatTypePathParams {
	return r.pathParams
}

func (r *LedgerVatTypeRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVatTypeRequest) Method() string {
	return r.method
}

func (r LedgerVatTypeRequest) NewLedgerVatTypeRequestBody() LedgerVatTypeRequestBody {
	return LedgerVatTypeRequestBody{}
}

type LedgerVatTypeRequestBody struct{}

func (r *LedgerVatTypeRequest) RequestBody() *LedgerVatTypeRequestBody {
	return &r.requestBody
}

func (r *LedgerVatTypeRequest) SetRequestBody(body LedgerVatTypeRequestBody) {
	r.requestBody = body
}

func (r *LedgerVatTypeRequest) NewResponseBody() *LedgerVatTypeResponseBody {
	return &LedgerVatTypeResponseBody{}
}

type LedgerVatTypeResponseBody struct {
	FullResultSize int    `json:"fullResultSize"`
	From           int    `json:"from"`
	Count          int    `json:"count"`
	VersionDigest  string `json:"versionDigest"`
	Values         []struct {
		ID         int     `json:"id"`
		Version    int     `json:"version"`
		URL        string  `json:"url"`
		Name       string  `json:"name"`
		Number     string  `json:"number"`
		Percentage float64 `json:"percentage"`
	} `json:"values"`
}

func (r *LedgerVatTypeRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/vatType", r.PathParams())
}

func (r *LedgerVatTypeRequest) Do() (LedgerVatTypeResponseBody, error) {
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
