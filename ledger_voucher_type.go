package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVoucherTypeRequest() LedgerVoucherTypeRequest {
	r := LedgerVoucherTypeRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVoucherTypeQueryParams()
	r.pathParams = r.NewLedgerVoucherTypePathParams()
	r.requestBody = r.NewLedgerVoucherTypeRequestBody()
	return r
}

type LedgerVoucherTypeRequest struct {
	client      *Client
	queryParams *LedgerVoucherTypeQueryParams
	pathParams  *LedgerVoucherTypePathParams
	method      string
	headers     http.Header
	requestBody LedgerVoucherTypeRequestBody
}

func (r LedgerVoucherTypeRequest) NewLedgerVoucherTypeQueryParams() *LedgerVoucherTypeQueryParams {
	return &LedgerVoucherTypeQueryParams{
		From:  0,
		Count: 1000,
	}
}

type LedgerVoucherTypeQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p LedgerVoucherTypeQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVoucherTypeRequest) QueryParams() *LedgerVoucherTypeQueryParams {
	return r.queryParams
}

func (r LedgerVoucherTypeRequest) NewLedgerVoucherTypePathParams() *LedgerVoucherTypePathParams {
	return &LedgerVoucherTypePathParams{}
}

type LedgerVoucherTypePathParams struct {
}

func (p *LedgerVoucherTypePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerVoucherTypeRequest) PathParams() *LedgerVoucherTypePathParams {
	return r.pathParams
}

func (r *LedgerVoucherTypeRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVoucherTypeRequest) Method() string {
	return r.method
}

func (r LedgerVoucherTypeRequest) NewLedgerVoucherTypeRequestBody() LedgerVoucherTypeRequestBody {
	return LedgerVoucherTypeRequestBody{}
}

type LedgerVoucherTypeRequestBody struct{}

func (r *LedgerVoucherTypeRequest) RequestBody() *LedgerVoucherTypeRequestBody {
	return &r.requestBody
}

func (r *LedgerVoucherTypeRequest) SetRequestBody(body LedgerVoucherTypeRequestBody) {
	r.requestBody = body
}

func (r *LedgerVoucherTypeRequest) NewResponseBody() *LedgerVoucherTypeResponseBody {
	return &LedgerVoucherTypeResponseBody{}
}

type LedgerVoucherTypeResponseBody struct {
	FullResultSize int    `json:"fullResultSize"`
	From           int    `json:"from"`
	Count          int    `json:"count"`
	VersionDigest  string `json:"versionDigest"`
	Values         []struct {
		ID      int    `json:"id"`
		Version int    `json:"version"`
		URL     string `json:"url"`
		Name    string `json:"name"`
	} `json:"values"`
}

func (r *LedgerVoucherTypeRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/voucherType", r.PathParams())
}

func (r *LedgerVoucherTypeRequest) Do() (LedgerVoucherTypeResponseBody, error) {
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
