package multivers

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
)

const (
	CustomerInvoiceInfoListByFiscalYearPath              = "api/%s/CustomerInvoiceInfoList/ByFiscalYear/%d/%d.json"
	CustomerInvoiceInfoListByIDPath                      = "api/%s/CustomerInvoiceInfoList/%s.json"
	InvoiceStateAny                         InvoiceState = 0
	InvoiceStateOpen                        InvoiceState = 1
	InvoiceStatePayedInFull                 InvoiceState = 2
)

func NewCustomerInvoiceInfoListService(client *Client) *CustomerInvoiceInfoListService {
	return &CustomerInvoiceInfoListService{Client: client}
}

type CustomerInvoiceInfoListService struct {
	Client *Client
}

func (s *CustomerInvoiceInfoListService) GetByFiscalYear(database string, fiscalYear int, invoiceState InvoiceState, requestParams *CustomerInvoiceInfoListByFiscalYearGetParams, ctx context.Context) (*CustomerInvoiceInfoListByFiscalYearGetResponse, error) {
	method := "GET"
	responseBody := s.NewByFiscalYearGetResponse()
	path := fmt.Sprintf(CustomerInvoiceInfoListByFiscalYearPath, database, fiscalYear, invoiceState)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *CustomerInvoiceInfoListService) NewByFiscalYearGetParams() *CustomerInvoiceInfoListByFiscalYearGetParams {
	return &CustomerInvoiceInfoListByFiscalYearGetParams{}
}

type CustomerInvoiceInfoListByFiscalYearGetParams struct {
	CustomerID string `schema:"id,omitempty"`
}

func (p *CustomerInvoiceInfoListByFiscalYearGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}

func (s *CustomerInvoiceInfoListService) NewByFiscalYearGetResponse() *CustomerInvoiceInfoListByFiscalYearGetResponse {
	return &CustomerInvoiceInfoListByFiscalYearGetResponse{}
}

type CustomerInvoiceInfoListByFiscalYearGetResponse []CustomerInvoiceInfo

func (s *CustomerInvoiceInfoListService) GetByID(database string, ID string, requestParams *CustomerInvoiceInfoListByIDGetParams, ctx context.Context) (*CustomerInvoiceInfoListByIDGetResponse, error) {
	method := "GET"
	responseBody := s.NewByIDGetResponse()
	path := fmt.Sprintf(CustomerInvoiceInfoListByIDPath, database, ID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *CustomerInvoiceInfoListService) NewByIDGetParams() *CustomerInvoiceInfoListByIDGetParams {
	return &CustomerInvoiceInfoListByIDGetParams{}
}

type CustomerInvoiceInfoListByIDGetParams struct {
	StartDate    DateNLNL     `schema:"startDate,omitempty"`
	EndDate      DateNLNL     `schema:"endDate,omitempty"`
	InvoiceState InvoiceState `schema:"invoiceState,omitempty"`
}

func (p *CustomerInvoiceInfoListByIDGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}

func (s *CustomerInvoiceInfoListService) NewByIDGetResponse() *CustomerInvoiceInfoListByIDGetResponse {
	return &CustomerInvoiceInfoListByIDGetResponse{}
}

type CustomerInvoiceInfoListByIDGetResponse []CustomerInvoiceInfo

type CustomerInvoiceInfo struct {
	City                      string                    `json:"city"`
	ContactPerson             string                    `json:"contactPerson"`
	CreditSqueezeRemaining    float64                   `json:"creditSqueezeRemaining"`
	CreditSqueezeRemainingCur float64                   `json:"creditSqueezeRemainingCur"`
	CurrencyDescription       string                    `json:"currencyDescription"`
	CurrencyID                string                    `json:"currencyId"`
	CurrentExchangeRate       float64                   `json:"currentExchangeRate"`
	CustomerID                string                    `json:"customerId"`
	CustomerInvoiceLines      []CustomerInvoiceInfoLine `json:"customerInvoiceLines"`
	CustomerName              string                    `json:"customerName"`
	DaysOld                   int                       `json:"daysOld"`
	DunForPayment             bool                      `json:"dunForPayment"`
	ExchangeRate              float64                   `json:"exchangeRate"`
	FiscalYear                int                       `json:"fiscalYear"`
	InvoiceAmount             float64                   `json:"invoiceAmount"`
	InvoicieAmountCur         float64                   `json:"invoiceAmountCur"`
	InvoiceBalance            float64                   `json:"invoiceBalance"`
	InvoiceBalanceCur         float64                   `json:"invoiceBalanceCur"`
	InvoiceDate               DateNLNL                  `json:"invoiceDate"`
	InvoiceExpirationDate     DateNLNL                  `json:"invoiceExpirationDate"`
	InvoiceID                 string                    `json:"invoiceId"`
	InvoiceReference          string                    `json:"invoiceReference"`
	IsAdvance                 bool                      `json:"isAdvance"`
	Name                      string                    `json:"name"`
	PaymentConditionID        string                    `json:"paymentConditionId"`
	PaymentDate               DateNLNL                  `json:"paymentDate"`
	PaymentReference          string                    `json:"paymentReference"`
	PhoneNumber               string                    `json:"phoneNumber"`
	RebateExpirationDate      DateNLNL                  `json:"rebateExpirationDate"`
	RebateRemaining           float64                   `json:"rebateRemaining"`
	RebateRemainingCur        float64                   `json:"rebateRemainingCur"`
	ReminderCount             float64                   `json:"reminderCount"`
	SettledAmount             float64                   `json:"settledAmount"`
	SettledAmountCur          float64                   `json:"settledAmountCur"`
	Shortname                 string                    `json:"shortName"`
	State                     InvoiceState              `json:"state"`
	Street                    string                    `json:"street"`
	TurnoverAmount            float64                   `json:"turnoverAmount"`
	VatAmount                 float64                   `json:"vatAmount"`
	VatAmountCur              float64                   `json:"vatAmountCur"`
	ZipCode                   string                    `json:"zipCode"`
}

type InvoiceState int

type CustomerInvoiceInfoLine struct {
	Amount        float64 `json:"amount"`
	AmountCur     float64 `json:"amountCur"`
	Description   string  `json:"description"`
	Permanence    bool    `json:"permanence"`
	VatCodeID     int     `json:"vatCodeId"`
	VatPercentage float64 `json:"vatPercentage"`
}
