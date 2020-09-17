package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCustomerInvoiceInfoListByJournalGetRequest() CustomerInvoiceInfoListByJournalGetRequest {
	r := CustomerInvoiceInfoListByJournalGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerInvoiceInfoListByJournalGetQueryParams()
	r.pathParams = r.NewCustomerInvoiceInfoListByJournalGetPathParams()
	r.requestBody = r.NewCustomerInvoiceInfoListByJournalGetRequestBody()
	return r
}

type CustomerInvoiceInfoListByJournalGetRequest struct {
	client      *Client
	queryParams *CustomerInvoiceInfoListByJournalGetQueryParams
	pathParams  *CustomerInvoiceInfoListByJournalGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoiceInfoListByJournalGetRequestBody
}

func (r CustomerInvoiceInfoListByJournalGetRequest) NewCustomerInvoiceInfoListByJournalGetQueryParams() *CustomerInvoiceInfoListByJournalGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &CustomerInvoiceInfoListByJournalGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type CustomerInvoiceInfoListByJournalGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerInvoiceInfoListByJournalGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) QueryParams() *CustomerInvoiceInfoListByJournalGetQueryParams {
	return r.queryParams
}

func (r CustomerInvoiceInfoListByJournalGetRequest) NewCustomerInvoiceInfoListByJournalGetPathParams() *CustomerInvoiceInfoListByJournalGetPathParams {
	return &CustomerInvoiceInfoListByJournalGetPathParams{}
}

type CustomerInvoiceInfoListByJournalGetPathParams struct {
	FiscalYear   int
	JournalID    string
	InvoiceState int
}

func (p *CustomerInvoiceInfoListByJournalGetPathParams) Params() map[string]string {
	return map[string]string{
		"fiscal_year":   strconv.Itoa(p.FiscalYear),
		"journal_id":    p.JournalID,
		"invoice_state": strconv.Itoa(p.InvoiceState),
	}
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) PathParams() *CustomerInvoiceInfoListByJournalGetPathParams {
	return r.pathParams
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) Method() string {
	return r.method
}

func (r CustomerInvoiceInfoListByJournalGetRequest) NewCustomerInvoiceInfoListByJournalGetRequestBody() CustomerInvoiceInfoListByJournalGetRequestBody {
	return CustomerInvoiceInfoListByJournalGetRequestBody{}
}

type CustomerInvoiceInfoListByJournalGetRequestBody struct{}

func (r *CustomerInvoiceInfoListByJournalGetRequest) RequestBody() *CustomerInvoiceInfoListByJournalGetRequestBody {
	return &r.requestBody
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) SetRequestBody(body CustomerInvoiceInfoListByJournalGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) NewResponseBody() *CustomerInvoiceInfoListByJournalGetResponseBody {
	return &CustomerInvoiceInfoListByJournalGetResponseBody{}
}

type CustomerInvoiceInfoListByJournalGetResponseBody []struct {
	City                      string  `json:"city"`
	ContactPerson             string  `json:"contactPerson"`
	CreditSqueezeRemaining    float64 `json:"creditSqueezeRemaining"`
	CreditSqueezeRemainingCur float64 `json:"creditSqueezeRemainingCur"`
	CurrencyDescription       string  `json:"currencyDescription"`
	CurrencyID                string  `json:"currencyId"`
	CurrentExchangeRate       float64 `json:"currentExchangeRate"`
	CustomerID                string  `json:"customerId"`
	CustomerInvoiceLines      []struct {
		Amount        float64 `json:"amount"`
		AmountCur     float64 `json:"amountCur"`
		Description   string  `json:"description"`
		Permanence    bool    `json:"permanence"`
		VatCodeID     int     `json:"vatCodeId"`
		VatPercentage float64 `json:"vatPercentage"`
	} `json:"customerInvoiceLines"`
	CustomerName          string        `json:"customerName"`
	DaysOld               int           `json:"daysOld"`
	DunForPayment         bool          `json:"dunForPayment"`
	ExchangeRate          float64       `json:"exchangeRate"`
	FiscalYear            int           `json:"fiscalYear"`
	InvoiceAmount         float64       `json:"invoiceAmount"`
	InvoiceAmountCur      float64       `json:"invoiceAmountCur"`
	InvoiceBalance        float64       `json:"invoiceBalance"`
	InvoiceBalanceCur     float64       `json:"invoiceBalanceCur"`
	InvoiceDate           string        `json:"invoiceDate"`
	InvoiceExpirationDate string        `json:"invoiceExpirationDate"`
	InvoiceID             string        `json:"invoiceId"`
	InvoiceReference      string        `json:"invoiceReference"`
	IsAdvance             bool          `json:"isAdvance"`
	JournalID             string        `json:"journalId"`
	Name                  string        `json:"name"`
	PaymentConditionID    string        `json:"paymentConditionId"`
	PaymentDate           string        `json:"paymentDate"`
	PaymentReference      string        `json:"paymentReference"`
	Payments              []interface{} `json:"payments"`
	PhoneNumber           string        `json:"phoneNumber"`
	RebateAmount          float64       `json:"rebateAmount"`
	RebateAmountCur       float64       `json:"rebateAmountCur"`
	RebateExpirationDate  string        `json:"rebateExpirationDate"`
	RebateRemaining       float64       `json:"rebateRemaining"`
	RebateRemainingCur    float64       `json:"rebateRemainingCur"`
	ReminderCount         int           `json:"reminderCount"`
	SettledAmount         float64       `json:"settledAmount"`
	SettledAmountCur      float64       `json:"settledAmountCur"`
	ShortName             string        `json:"shortName"`
	State                 int           `json:"state"`
	Street                string        `json:"street"`
	TurnoverAmount        float64       `json:"turnoverAmount"`
	VatAmount             float64       `json:"vatAmount"`
	VatAmountCur          float64       `json:"vatAmountCur"`
	ZipCode               string        `json:"zipCode"`
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CustomerInvoiceInfoList/ByJournal/{{.fiscal_year}}/{{.journal_id}}//{{.invoice_state}}", r.PathParams())
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) Do() (CustomerInvoiceInfoListByJournalGetResponseBody, error) {
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
