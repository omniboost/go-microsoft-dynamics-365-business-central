package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewInvoceGetRequest() InvoceGetRequest {
	r := InvoceGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewInvoceGetQueryParams()
	r.pathParams = r.NewInvoceGetPathParams()
	r.requestBody = r.NewInvoceGetRequestBody()
	return r
}

type InvoceGetRequest struct {
	client      *Client
	queryParams *InvoceGetQueryParams
	pathParams  *InvoceGetPathParams
	method      string
	headers     http.Header
	requestBody InvoceGetRequestBody
}

func (r InvoceGetRequest) NewInvoceGetQueryParams() *InvoceGetQueryParams {
	return &InvoceGetQueryParams{}
}

type InvoceGetQueryParams struct {
	ID              string `schema:"id,omitempty"`
	InvoiceDateFrom Date   `schema:"invoiceDateFrom"`
	InvoiceDateTo   Date   `schema:"invoiceDateTo"`
	InvoiceNumber   string `schema:"invoiceNumber,omitempty"`
	Kid             string `schema:"kid,omitempty"`
	VoucherID       string `schema:"voucherId,omitempty"`
	CustomerID      string `schema:"customerId,omitempty"`
	From            int    `schema:"from,omitempty"`
	Count           int    `schema:"count,omitempty"`
	Sorting         int    `schema:"sorting,omitempty"`
	Fields          string `schema:"fields,omitempty"`
}

func (p InvoceGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InvoceGetRequest) QueryParams() *InvoceGetQueryParams {
	return r.queryParams
}

func (r InvoceGetRequest) NewInvoceGetPathParams() *InvoceGetPathParams {
	return &InvoceGetPathParams{}
}

type InvoceGetPathParams struct {
}

func (p *InvoceGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoceGetRequest) PathParams() *InvoceGetPathParams {
	return r.pathParams
}

func (r *InvoceGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoceGetRequest) Method() string {
	return r.method
}

func (r InvoceGetRequest) NewInvoceGetRequestBody() InvoceGetRequestBody {
	return InvoceGetRequestBody{}
}

type InvoceGetRequestBody struct{}

func (r *InvoceGetRequest) RequestBody() *InvoceGetRequestBody {
	return &r.requestBody
}

func (r *InvoceGetRequest) SetRequestBody(body InvoceGetRequestBody) {
	r.requestBody = body
}

func (r *InvoceGetRequest) NewResponseBody() *InvoceGetResponseBody {
	return &InvoceGetResponseBody{}
}

type InvoceGetResponseBody struct {
	FullResultSize int    `json:"fullResultSize"`
	From           int    `json:"from"`
	Count          int    `json:"count"`
	VersionDigest  string `json:"versionDigest"`
	Values         []struct {
		ID            int    `json:"id"`
		Version       int    `json:"version"`
		URL           string `json:"url"`
		InvoiceNumber int    `json:"invoiceNumber"`
		InvoiceDate   string `json:"invoiceDate"`
		Customer      struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"customer"`
		CreditedInvoice int    `json:"creditedInvoice"`
		IsCredited      bool   `json:"isCredited"`
		InvoiceDueDate  string `json:"invoiceDueDate"`
		Kid             string `json:"kid"`
		InvoiceComment  string `json:"invoiceComment"`
		Comment         string `json:"comment"`
		Orders          []struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"orders"`
		ProjectInvoiceDetails []struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"projectInvoiceDetails"`
		Voucher struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"voucher"`
		DeliveryDate               string  `json:"deliveryDate"`
		Amount                     float64 `json:"amount"`
		AmountCurrency             float64 `json:"amountCurrency"`
		AmountExcludingVat         float64 `json:"amountExcludingVat"`
		AmountExcludingVatCurrency float64 `json:"amountExcludingVatCurrency"`
		AmountRoundoff             float64 `json:"amountRoundoff"`
		AmountRoundoffCurrency     float64 `json:"amountRoundoffCurrency"`
		AmountOutstanding          float64 `json:"amountOutstanding"`
		AmountOutstandingTotal     float64 `json:"amountOutstandingTotal"`
		SumRemits                  int     `json:"sumRemits"`
		Currency                   struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"currency"`
		IsCreditNote bool `json:"isCreditNote"`
		IsCharged    bool `json:"isCharged"`
		IsApproved   bool `json:"isApproved"`
		Postings     []struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"postings"`
		Reminders      []interface{} `json:"reminders"`
		InvoiceRemarks string        `json:"invoiceRemarks"`
	} `json:"values"`
}

func (r *InvoceGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoice", r.PathParams())
}

func (r *InvoceGetRequest) Do() (InvoceGetResponseBody, error) {
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
