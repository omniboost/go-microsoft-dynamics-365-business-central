package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVoucherGetRequest() LedgerVoucherGetRequest {
	r := LedgerVoucherGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVoucherGetQueryParams()
	r.pathParams = r.NewLedgerVoucherGetPathParams()
	r.requestBody = r.NewLedgerVoucherGetRequestBody()
	return r
}

type LedgerVoucherGetRequest struct {
	client      *Client
	queryParams *LedgerVoucherGetQueryParams
	pathParams  *LedgerVoucherGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerVoucherGetRequestBody
}

func (r LedgerVoucherGetRequest) NewLedgerVoucherGetQueryParams() *LedgerVoucherGetQueryParams {
	return &LedgerVoucherGetQueryParams{
		From:  0,
		Count: 100,
	}
}

type LedgerVoucherGetQueryParams struct {
	From     int  `schema:"from"`
	Count    int  `schema:"count"`
	DateFrom Date `schema:"dateFrom"`
	DateTo   Date `schema:"dateTo"`
}

func (p LedgerVoucherGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVoucherGetRequest) QueryParams() *LedgerVoucherGetQueryParams {
	return r.queryParams
}

func (r LedgerVoucherGetRequest) NewLedgerVoucherGetPathParams() *LedgerVoucherGetPathParams {
	return &LedgerVoucherGetPathParams{}
}

type LedgerVoucherGetPathParams struct {
}

func (p *LedgerVoucherGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerVoucherGetRequest) PathParams() *LedgerVoucherGetPathParams {
	return r.pathParams
}

func (r *LedgerVoucherGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVoucherGetRequest) Method() string {
	return r.method
}

func (r LedgerVoucherGetRequest) NewLedgerVoucherGetRequestBody() LedgerVoucherGetRequestBody {
	return LedgerVoucherGetRequestBody{}
}

type LedgerVoucherGetRequestBody struct{}

func (r *LedgerVoucherGetRequest) RequestBody() *LedgerVoucherGetRequestBody {
	return &r.requestBody
}

func (r *LedgerVoucherGetRequest) SetRequestBody(body LedgerVoucherGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerVoucherGetRequest) NewResponseBody() *LedgerVoucherGetResponseBody {
	return &LedgerVoucherGetResponseBody{}
}

type LedgerVoucherGetResponseBody struct {
	FullResultSize int    `json:"fullResultSize"`
	From           int    `json:"from"`
	Count          int    `json:"count"`
	VersionDigest  string `json:"versionDigest"`
	Values         []struct {
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
		ReverseVoucher interface{}   `json:"reverseVoucher"`
		Postings       []interface{} `json:"postings"`
		Document       interface{}   `json:"document"`
		Attachment     struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"attachment"`
		EdiDocument struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"ediDocument"`
	} `json:"values"`
	TotalNumberOfPostings int `json:"totalNumberOfPostings"`
}

func (r *LedgerVoucherGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/voucher", r.PathParams())
}

func (r *LedgerVoucherGetRequest) Do() (LedgerVoucherGetResponseBody, error) {
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

func (r *LedgerVoucherGetRequest) All() (LedgerVoucherGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := LedgerVoucherGetResponseBody{}
	concat.Count = resp.Count
	concat.From = resp.From
	concat.FullResultSize = resp.FullResultSize
	concat.TotalNumberOfPostings = resp.TotalNumberOfPostings
	concat.Values = resp.Values
	concat.VersionDigest = resp.VersionDigest

	for concat.From+concat.Count < concat.FullResultSize {
		r.QueryParams().From = r.QueryParams().From + r.QueryParams().Count
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		concat.Count = resp.Count
		concat.From = resp.From
		concat.FullResultSize = resp.FullResultSize
		concat.TotalNumberOfPostings = resp.TotalNumberOfPostings
		concat.Values = append(concat.Values, resp.Values...)
		concat.VersionDigest = resp.VersionDigest
	}

	return concat, nil
}
