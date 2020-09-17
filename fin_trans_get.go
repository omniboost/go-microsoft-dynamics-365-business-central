package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewFinTransGetRequest() FinTransGetRequest {
	r := FinTransGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewFinTransGetQueryParams()
	r.pathParams = r.NewFinTransGetPathParams()
	r.requestBody = r.NewFinTransGetRequestBody()
	return r
}

type FinTransGetRequest struct {
	client      *Client
	queryParams *FinTransGetQueryParams
	pathParams  *FinTransGetPathParams
	method      string
	headers     http.Header
	requestBody FinTransGetRequestBody
}

func (r FinTransGetRequest) NewFinTransGetQueryParams() *FinTransGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &FinTransGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type FinTransGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p FinTransGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinTransGetRequest) QueryParams() *FinTransGetQueryParams {
	return r.queryParams
}

func (r FinTransGetRequest) NewFinTransGetPathParams() *FinTransGetPathParams {
	return &FinTransGetPathParams{}
}

type FinTransGetPathParams struct {
	Year      int    `schema:"year"`
	JournalID string `schema:"journal_id"`
}

func (p *FinTransGetPathParams) Params() map[string]string {
	return map[string]string{
		"year":       strconv.Itoa(p.Year),
		"journal_id": p.JournalID,
	}
}

func (r *FinTransGetRequest) PathParams() *FinTransGetPathParams {
	return r.pathParams
}

func (r *FinTransGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinTransGetRequest) Method() string {
	return r.method
}

func (r FinTransGetRequest) NewFinTransGetRequestBody() FinTransGetRequestBody {
	return FinTransGetRequestBody{}
}

type FinTransGetRequestBody struct{}

func (r *FinTransGetRequest) RequestBody() *FinTransGetRequestBody {
	return &r.requestBody
}

func (r *FinTransGetRequest) SetRequestBody(body FinTransGetRequestBody) {
	r.requestBody = body
}

func (r *FinTransGetRequest) NewResponseBody() *FinTransGetResponseBody {
	return &FinTransGetResponseBody{}
}

type FinTransGetResponseBody struct {
	Messages           []interface{} `json:"messages"`
	CanChange          bool          `json:"canChange"`
	CannotChangeReason string        `json:"cannotChangeReason"`
	CurrencyID         string        `json:"currencyId"`
	Description        string        `json:"description"`
	Document           string        `json:"document"`
	ExchangeRate       float64       `json:"exchangeRate"`
	FinTransEntries    []struct {
		Type                            string        `json:"$type"`
		AccountID                       string        `json:"accountId"`
		Messages                        []interface{} `json:"messages"`
		CanChange                       bool          `json:"canChange"`
		CannotChangeReason              string        `json:"cannotChangeReason"`
		CostCentreID                    string        `json:"costCentreId"`
		CostCentreIDRequired            bool          `json:"costCentreIdRequired"`
		CostUnitID                      string        `json:"costUnitId"`
		CostUnitIDRequired              bool          `json:"costUnitIdRequired"`
		CreditAmount                    float64       `json:"creditAmount"`
		DebitAmount                     float64       `json:"debitAmount"`
		Description                     string        `json:"description"`
		Document                        string        `json:"document"`
		EntryType                       int           `json:"entryType"`
		InvoiceType                     int           `json:"invoiceType"`
		IsJournalEntrySplit             bool          `json:"isJournalEntrySplit"`
		IsSubAdminSpecificationRequired bool          `json:"isSubAdminSpecificationRequired"`
		JournalSection                  int           `json:"journalSection"`
		LineNumbers                     []int         `json:"lineNumbers"`
		NondeductibleVatPercentage      float64       `json:"nondeductibleVatPercentage"`
		PermCode                        interface{}   `json:"permCode"`
		Quantity                        float64       `json:"quantity"`
		SplitDescription                string        `json:"splitDescription"`
		SplitJournalEntry               bool          `json:"splitJournalEntry"`
		StartDateOfSplit                string        `json:"startDateOfSplit"`
		SubAdminSpecifications          []interface{} `json:"subAdminSpecifications"`
		SubTransaction                  string        `json:"subTransaction"`
		TeleBankGUID                    string        `json:"teleBankGuid"`
		TotalPeriodsToSplit             interface{}   `json:"totalPeriodsToSplit"`
		TransactionDate                 string        `json:"transactionDate"`
		VatAmount                       float64       `json:"vatAmount"`
		VatChanged                      bool          `json:"vatChanged"`
		VatCodeID                       int           `json:"vatCodeId"`
		VatIncluded                     bool          `json:"vatIncluded"`
		VatOnInvoice                    bool          `json:"vatOnInvoice"`
		VatScenarioID                   int           `json:"vatScenarioId"`
		VatType                         int           `json:"vatType"`
	} `json:"finTransEntries"`
	FiscalYear                int     `json:"fiscalYear"`
	JournalID                 string  `json:"journalId"`
	JournalTransaction        int     `json:"journalTransaction"`
	OpeningBalanceTransaction bool    `json:"openingBalanceTransaction"`
	PeriodNumber              int     `json:"periodNumber"`
	PreviousBalance           float64 `json:"previousBalance"`
	TotalCreditAmount         float64 `json:"totalCreditAmount"`
	TotalDebitAmount          float64 `json:"totalDebitAmount"`
	TransactionDate           string  `json:"transactionDate"`
}

func (r *FinTransGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/FinTrans/{{.year}}/{{.journal_id}}/2", r.PathParams())
}

func (r *FinTransGetRequest) Do() (FinTransGetResponseBody, error) {
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
