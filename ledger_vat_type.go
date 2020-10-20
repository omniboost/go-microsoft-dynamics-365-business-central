package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVATTypeRequest() LedgerVATTypeRequest {
	r := LedgerVATTypeRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVATTypeQueryParams()
	r.pathParams = r.NewLedgerVATTypePathParams()
	r.requestBody = r.NewLedgerVATTypeRequestBody()
	return r
}

type LedgerVATTypeRequest struct {
	client      *Client
	queryParams *LedgerVATTypeQueryParams
	pathParams  *LedgerVATTypePathParams
	method      string
	headers     http.Header
	requestBody LedgerVATTypeRequestBody
}

func (r LedgerVATTypeRequest) NewLedgerVATTypeQueryParams() *LedgerVATTypeQueryParams {
	return &LedgerVATTypeQueryParams{
		From:  0,
		Count: 1000,
	}
}

type LedgerVATTypeQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p LedgerVATTypeQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVATTypeRequest) QueryParams() *LedgerVATTypeQueryParams {
	return r.queryParams
}

func (r LedgerVATTypeRequest) NewLedgerVATTypePathParams() *LedgerVATTypePathParams {
	return &LedgerVATTypePathParams{}
}

type LedgerVATTypePathParams struct {
}

func (p *LedgerVATTypePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerVATTypeRequest) PathParams() *LedgerVATTypePathParams {
	return r.pathParams
}

func (r *LedgerVATTypeRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVATTypeRequest) Method() string {
	return r.method
}

func (r LedgerVATTypeRequest) NewLedgerVATTypeRequestBody() LedgerVATTypeRequestBody {
	return LedgerVATTypeRequestBody{}
}

type LedgerVATTypeRequestBody struct{}

func (r *LedgerVATTypeRequest) RequestBody() *LedgerVATTypeRequestBody {
	return &r.requestBody
}

func (r *LedgerVATTypeRequest) SetRequestBody(body LedgerVATTypeRequestBody) {
	r.requestBody = body
}

func (r *LedgerVATTypeRequest) NewResponseBody() *LedgerVATTypeResponseBody {
	return &LedgerVATTypeResponseBody{}
}

type LedgerVATTypeResponseBody struct {
	FullResultSize int      `json:"fullResultSize"`
	From           int      `json:"from"`
	Count          int      `json:"count"`
	VersionDigest  string   `json:"versionDigest"`
	Values         VATTypes `json:"values"`
}

func (r *LedgerVATTypeRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/vatType", r.PathParams())
}

func (r *LedgerVATTypeRequest) Do() (LedgerVATTypeResponseBody, error) {
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
