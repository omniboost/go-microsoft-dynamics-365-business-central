package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewJournalEntryLineInfoListByPeriodRequest() JournalEntryLineInfoListByPeriodRequest {
	r := JournalEntryLineInfoListByPeriodRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewJournalEntryLineInfoListByPeriodQueryParams()
	r.pathParams = r.NewJournalEntryLineInfoListByPeriodPathParams()
	r.requestBody = r.NewJournalEntryLineInfoListByPeriodRequestBody()
	return r
}

type JournalEntryLineInfoListByPeriodRequest struct {
	client      *Client
	queryParams *JournalEntryLineInfoListByPeriodQueryParams
	pathParams  *JournalEntryLineInfoListByPeriodPathParams
	method      string
	headers     http.Header
	requestBody JournalEntryLineInfoListByPeriodRequestBody
}

func (r JournalEntryLineInfoListByPeriodRequest) NewJournalEntryLineInfoListByPeriodQueryParams() *JournalEntryLineInfoListByPeriodQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &JournalEntryLineInfoListByPeriodQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type JournalEntryLineInfoListByPeriodQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p JournalEntryLineInfoListByPeriodQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalEntryLineInfoListByPeriodRequest) QueryParams() *JournalEntryLineInfoListByPeriodQueryParams {
	return r.queryParams
}

func (r JournalEntryLineInfoListByPeriodRequest) NewJournalEntryLineInfoListByPeriodPathParams() *JournalEntryLineInfoListByPeriodPathParams {
	return &JournalEntryLineInfoListByPeriodPathParams{}
}

type JournalEntryLineInfoListByPeriodPathParams struct {
	Year        int `schema:"year"`
	StartPeriod int `schema:"start_period"`
	EndPeriod   int `schema:"end_period"`
}

func (p *JournalEntryLineInfoListByPeriodPathParams) Params() map[string]string {
	return map[string]string{
		"year":         strconv.Itoa(p.Year),
		"start_period": strconv.Itoa(p.StartPeriod),
		"end_period":   strconv.Itoa(p.EndPeriod),
	}
}

func (r *JournalEntryLineInfoListByPeriodRequest) PathParams() *JournalEntryLineInfoListByPeriodPathParams {
	return r.pathParams
}

func (r *JournalEntryLineInfoListByPeriodRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntryLineInfoListByPeriodRequest) Method() string {
	return r.method
}

func (r JournalEntryLineInfoListByPeriodRequest) NewJournalEntryLineInfoListByPeriodRequestBody() JournalEntryLineInfoListByPeriodRequestBody {
	return JournalEntryLineInfoListByPeriodRequestBody{}
}

type JournalEntryLineInfoListByPeriodRequestBody struct{}

func (r *JournalEntryLineInfoListByPeriodRequest) RequestBody() *JournalEntryLineInfoListByPeriodRequestBody {
	return &r.requestBody
}

func (r *JournalEntryLineInfoListByPeriodRequest) SetRequestBody(body JournalEntryLineInfoListByPeriodRequestBody) {
	r.requestBody = body
}

func (r *JournalEntryLineInfoListByPeriodRequest) NewResponseBody() *JournalEntryLineInfoListByPeriodResponseBody {
	return &JournalEntryLineInfoListByPeriodResponseBody{}
}

type JournalEntryLineInfoListByPeriodResponseBody []struct {
	AccountID                 string      `json:"accountId"`
	CostCentreID              string      `json:"costCentreId"`
	CostUnitID                string      `json:"costUnitId"`
	CreditAmount              float64     `json:"creditAmount"`
	CreditAmountCur           float64     `json:"creditAmountCur"`
	CurrencyID                string      `json:"currencyId"`
	CustomerID                string      `json:"customerId"`
	DebitAmount               float64     `json:"debitAmount"`
	DebitAmountCur            float64     `json:"debitAmountCur"`
	Description               string      `json:"description"`
	Document                  string      `json:"document"`
	FiscalYear                int         `json:"fiscalYear"`
	FixedAssetID              string      `json:"fixedAssetId"`
	JournalID                 string      `json:"journalId"`
	JournalTransaction        int         `json:"journalTransaction"`
	OpeningBalanceTransaction bool        `json:"openingBalanceTransaction"`
	PeriodNumber              int         `json:"periodNumber"`
	Quantity                  interface{} `json:"quantity"`
	SupplierID                string      `json:"supplierId"`
	TransactionDate           string      `json:"transactionDate"`
	VatAdjusted               bool        `json:"vatAdjusted"`
	VatAmount                 float64     `json:"vatAmount"`
	VatCodeID                 interface{} `json:"vatCodeId"`
	VatType                   int         `json:"vatType"`
}

func (r *JournalEntryLineInfoListByPeriodRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/JournalEntryLineInfoList/ByPeriod/{{.year}}/{{.start_period}}/{{.end_period}}", r.PathParams())
}

func (r *JournalEntryLineInfoListByPeriodRequest) Do() (JournalEntryLineInfoListByPeriodResponseBody, error) {
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
