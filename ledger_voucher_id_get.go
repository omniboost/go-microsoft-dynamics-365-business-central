package tripletex

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVoucherIDGetRequest() LedgerVoucherIDGetRequest {
	r := LedgerVoucherIDGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVoucherIDGetQueryParams()
	r.pathParams = r.NewLedgerVoucherIDGetPathParams()
	r.requestBody = r.NewLedgerVoucherIDGetRequestBody()
	return r
}

type LedgerVoucherIDGetRequest struct {
	client      *Client
	queryParams *LedgerVoucherIDGetQueryParams
	pathParams  *LedgerVoucherIDGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerVoucherIDGetRequestBody
}

func (r LedgerVoucherIDGetRequest) NewLedgerVoucherIDGetQueryParams() *LedgerVoucherIDGetQueryParams {
	return &LedgerVoucherIDGetQueryParams{}
}

type LedgerVoucherIDGetQueryParams struct {
}

func (p LedgerVoucherIDGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVoucherIDGetRequest) QueryParams() *LedgerVoucherIDGetQueryParams {
	return r.queryParams
}

func (r LedgerVoucherIDGetRequest) NewLedgerVoucherIDGetPathParams() *LedgerVoucherIDGetPathParams {
	return &LedgerVoucherIDGetPathParams{}
}

type LedgerVoucherIDGetPathParams struct {
	ID int
}

func (p *LedgerVoucherIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *LedgerVoucherIDGetRequest) PathParams() *LedgerVoucherIDGetPathParams {
	return r.pathParams
}

func (r *LedgerVoucherIDGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVoucherIDGetRequest) Method() string {
	return r.method
}

func (r LedgerVoucherIDGetRequest) NewLedgerVoucherIDGetRequestBody() LedgerVoucherIDGetRequestBody {
	return LedgerVoucherIDGetRequestBody{}
}

type LedgerVoucherIDGetRequestBody struct{}

func (r *LedgerVoucherIDGetRequest) RequestBody() *LedgerVoucherIDGetRequestBody {
	return &r.requestBody
}

func (r *LedgerVoucherIDGetRequest) SetRequestBody(body LedgerVoucherIDGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerVoucherIDGetRequest) NewResponseBody() *LedgerVoucherIDGetResponseBody {
	return &LedgerVoucherIDGetResponseBody{}
}

type LedgerVoucherIDGetResponseBody struct {
	Value struct {
		ID          int    `json:"id"`
		Version     int    `json:"version"`
		URL         string `json:"url"`
		Date        string `json:"date"`
		Number      int    `json:"number"`
		Year        int    `json:"year"`
		Description string `json:"description"`
		VoucherType struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"voucherType"`
		ReverseVoucher interface{} `json:"reverseVoucher"`
		Postings       []struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"postings"`
		Document struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"document"`
		Attachment  interface{} `json:"attachment"`
		EdiDocument struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"ediDocument"`
	} `json:"value"`
}

func (r *LedgerVoucherIDGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/voucher/{{.id}}", r.PathParams())
}

func (r *LedgerVoucherIDGetRequest) Do() (LedgerVoucherIDGetResponseBody, error) {
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
