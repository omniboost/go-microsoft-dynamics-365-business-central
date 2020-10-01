package tripletex

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerPostingIDGetRequest() LedgerPostingIDGetRequest {
	r := LedgerPostingIDGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerPostingIDGetQueryParams()
	r.pathParams = r.NewLedgerPostingIDGetPathParams()
	r.requestBody = r.NewLedgerPostingIDGetRequestBody()
	return r
}

type LedgerPostingIDGetRequest struct {
	client      *Client
	queryParams *LedgerPostingIDGetQueryParams
	pathParams  *LedgerPostingIDGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerPostingIDGetRequestBody
}

func (r LedgerPostingIDGetRequest) NewLedgerPostingIDGetQueryParams() *LedgerPostingIDGetQueryParams {
	return &LedgerPostingIDGetQueryParams{}
}

type LedgerPostingIDGetQueryParams struct {
}

func (p LedgerPostingIDGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerPostingIDGetRequest) QueryParams() *LedgerPostingIDGetQueryParams {
	return r.queryParams
}

func (r LedgerPostingIDGetRequest) NewLedgerPostingIDGetPathParams() *LedgerPostingIDGetPathParams {
	return &LedgerPostingIDGetPathParams{}
}

type LedgerPostingIDGetPathParams struct {
	ID int
}

func (p *LedgerPostingIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *LedgerPostingIDGetRequest) PathParams() *LedgerPostingIDGetPathParams {
	return r.pathParams
}

func (r *LedgerPostingIDGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerPostingIDGetRequest) Method() string {
	return r.method
}

func (r LedgerPostingIDGetRequest) NewLedgerPostingIDGetRequestBody() LedgerPostingIDGetRequestBody {
	return LedgerPostingIDGetRequestBody{}
}

type LedgerPostingIDGetRequestBody struct{}

func (r *LedgerPostingIDGetRequest) RequestBody() *LedgerPostingIDGetRequestBody {
	return &r.requestBody
}

func (r *LedgerPostingIDGetRequest) SetRequestBody(body LedgerPostingIDGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerPostingIDGetRequest) NewResponseBody() *LedgerPostingIDGetResponseBody {
	return &LedgerPostingIDGetResponseBody{}
}

type LedgerPostingIDGetResponseBody struct {
	Value struct {
		ID      int    `json:"id"`
		Version int    `json:"version"`
		URL     string `json:"url"`
		Voucher struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"voucher"`
		Date        string `json:"date"`
		Description string `json:"description"`
		Account     struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"account"`
		Customer struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"customer"`
		Supplier interface{} `json:"supplier"`
		Employee struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"employee"`
		Project    interface{} `json:"project"`
		Product    interface{} `json:"product"`
		Department interface{} `json:"department"`
		VatType    struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"vatType"`
		Amount              float64 `json:"amount"`
		AmountCurrency      float64 `json:"amountCurrency"`
		AmountGross         float64 `json:"amountGross"`
		AmountGrossCurrency float64 `json:"amountGrossCurrency"`
		Currency            struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"currency"`
		CloseGroup      interface{} `json:"closeGroup"`
		InvoiceNumber   string      `json:"invoiceNumber"`
		TermOfPayment   string      `json:"termOfPayment"`
		Row             int         `json:"row"`
		SystemGenerated bool        `json:"systemGenerated"`
	} `json:"value"`
}

func (r *LedgerPostingIDGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/posting/{{.id}}", r.PathParams())
}

func (r *LedgerPostingIDGetRequest) Do() (LedgerPostingIDGetResponseBody, error) {
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
