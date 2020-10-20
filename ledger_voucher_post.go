package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVoucherPostRequest() LedgerVoucherPostRequest {
	r := LedgerVoucherPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVoucherPostQueryParams()
	r.pathParams = r.NewLedgerVoucherPostPathParams()
	r.requestBody = r.NewLedgerVoucherPostRequestBody()
	return r
}

type LedgerVoucherPostRequest struct {
	client      *Client
	queryParams *LedgerVoucherPostQueryParams
	pathParams  *LedgerVoucherPostPathParams
	method      string
	headers     http.Header
	requestBody LedgerVoucherPostRequestBody
}

func (r LedgerVoucherPostRequest) NewLedgerVoucherPostQueryParams() *LedgerVoucherPostQueryParams {
	return &LedgerVoucherPostQueryParams{}
}

type LedgerVoucherPostQueryParams struct {
	SendToLedger bool `schema:"sendToLedger"`
}

func (p LedgerVoucherPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVoucherPostRequest) QueryParams() *LedgerVoucherPostQueryParams {
	return r.queryParams
}

func (r LedgerVoucherPostRequest) NewLedgerVoucherPostPathParams() *LedgerVoucherPostPathParams {
	return &LedgerVoucherPostPathParams{}
}

type LedgerVoucherPostPathParams struct {
}

func (p *LedgerVoucherPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerVoucherPostRequest) PathParams() *LedgerVoucherPostPathParams {
	return r.pathParams
}

func (r *LedgerVoucherPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVoucherPostRequest) Method() string {
	return r.method
}

func (r LedgerVoucherPostRequest) NewLedgerVoucherPostRequestBody() LedgerVoucherPostRequestBody {
	return LedgerVoucherPostRequestBody{}
}

type LedgerVoucherPostRequestBody Voucher

func (r *LedgerVoucherPostRequest) RequestBody() *LedgerVoucherPostRequestBody {
	return &r.requestBody
}

func (r *LedgerVoucherPostRequest) SetRequestBody(body LedgerVoucherPostRequestBody) {
	r.requestBody = body
}

func (r *LedgerVoucherPostRequest) NewResponseBody() *LedgerVoucherPostResponseBody {
	return &LedgerVoucherPostResponseBody{}
}

type LedgerVoucherPostResponseBody struct{}

func (r *LedgerVoucherPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/voucher", r.PathParams())
}

func (r *LedgerVoucherPostRequest) Do() (LedgerVoucherPostResponseBody, error) {
	// fetch a new token if it isn't set already
	if r.client.token == "" {
		var err error
		r.client.token, err = r.client.NewToken()
		if err != nil {
			return *r.NewResponseBody(), err
		}
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
