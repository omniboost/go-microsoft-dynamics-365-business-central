package dkplus

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewGeneralLedgerTransactionGetRequest() GeneralLedgerTransactionGetRequest {
	r := GeneralLedgerTransactionGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GeneralLedgerTransactionGetRequest struct {
	client      *Client
	queryParams *GeneralLedgerTransactionGetQueryParams
	pathParams  *GeneralLedgerTransactionGetPathParams
	method      string
	headers     http.Header
	requestBody GeneralLedgerTransactionGetRequestBody
}

func (r GeneralLedgerTransactionGetRequest) NewQueryParams() *GeneralLedgerTransactionGetQueryParams {
	return &GeneralLedgerTransactionGetQueryParams{}
}

type GeneralLedgerTransactionGetQueryParams struct {
	Account       string   `schema:"account,omitempty"`
	Dim1          string   `schema:"dim1,omitempty"`
	Voucher       string   `schema:"voucher,omitempty"`
	Reference     string   `schema:"reference,omitempty"`
	CreatedAfter  DateTime `schema:"createdAfter,omitempty"`
	CreatedBefore DateTime `schema:"createdBefore,omitempty"`
	DueAfter      DateTime `schema:"dueAfter,omitempty"`
}

func (p GeneralLedgerTransactionGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GeneralLedgerTransactionGetRequest) QueryParams() *GeneralLedgerTransactionGetQueryParams {
	return r.queryParams
}

func (r GeneralLedgerTransactionGetRequest) NewPathParams() *GeneralLedgerTransactionGetPathParams {
	return &GeneralLedgerTransactionGetPathParams{
		Page:  1,
		Count: 100,
	}
}

type GeneralLedgerTransactionGetPathParams struct {
	Page  int
	Count int
}

func (p *GeneralLedgerTransactionGetPathParams) Params() map[string]string {
	return map[string]string{
		"page":  strconv.Itoa(p.Page),
		"count": strconv.Itoa(p.Count),
	}
}

func (r *GeneralLedgerTransactionGetRequest) PathParams() *GeneralLedgerTransactionGetPathParams {
	return r.pathParams
}

func (r *GeneralLedgerTransactionGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GeneralLedgerTransactionGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GeneralLedgerTransactionGetRequest) Method() string {
	return r.method
}

func (r GeneralLedgerTransactionGetRequest) NewRequestBody() GeneralLedgerTransactionGetRequestBody {
	return GeneralLedgerTransactionGetRequestBody{}
}

type GeneralLedgerTransactionGetRequestBody struct {
}

func (r *GeneralLedgerTransactionGetRequest) RequestBody() *GeneralLedgerTransactionGetRequestBody {
	return &r.requestBody
}

func (r *GeneralLedgerTransactionGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GeneralLedgerTransactionGetRequest) SetRequestBody(body GeneralLedgerTransactionGetRequestBody) {
	r.requestBody = body
}

func (r *GeneralLedgerTransactionGetRequest) NewResponseBody() *GeneralLedgerTransactionGetResponseBody {
	return &GeneralLedgerTransactionGetResponseBody{}
}

type GeneralLedgerTransactionGetResponseBody []struct {
	ID                  int       `json:"ID"`
	InvoiceNumber       string    `json:"InvoiceNumber,omitempty"`
	Account             string    `json:"Account"`
	Created             time.Time `json:"Created"`
	CreatedBy           string    `json:"CreatedBy"`
	Modified            time.Time `json:"Modified"`
	DueDate             string    `json:"DueDate"`
	Text                string    `json:"Text"`
	Reference           string    `json:"Reference,omitempty"`
	JournalDate         string    `json:"JournalDate"`
	Origin              int       `json:"Origin"`
	Voucher             string    `json:"Voucher"`
	JournalType         int       `json:"JournalType"`
	Code                int       `json:"Code"`
	Currency            string    `json:"Currency"`
	Exchange            float64   `json:"Exchange"`
	Amount              float64   `json:"Amount"`
	InputAmount         float64   `json:"InputAmount"`
	TaxCode             string    `json:"TaxCode,omitempty"`
	TaxGroup            int       `json:"TaxGroup"`
	TaxPercent          float64   `json:"TaxPercent"`
	HCode               string    `json:"HCode,omitempty"`
	SubAccount          string    `json:"SubAccount"`
	HType               int       `json:"HType"`
	IsCredit            bool      `json:"IsCredit"`
	PeriodID            int       `json:"PeriodId"`
	Quantity            float64   `json:"Quantity"`
	NumberOfQuantity    float64   `json:"NumberOfQuantity"`
	PercentageOfVATUsed float64   `json:"PercentageOfVATUsed"`
	VATReportID         int       `json:"VATReportID"`
	Dim1                string    `json:"Dim1,omitempty"`
	Dim2                string    `json:"Dim2,omitempty"`
}

func (r *GeneralLedgerTransactionGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("generalledger/transaction/page/{{.page}}/{{.count}}", r.PathParams())
	return &u
}

func (r *GeneralLedgerTransactionGetRequest) Do() (GeneralLedgerTransactionGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
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
