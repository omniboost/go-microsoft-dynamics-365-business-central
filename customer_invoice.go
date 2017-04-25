package multivers

// https://api.online.unit4.nl/V16/Help/Api/GET-api-database-ProductGroup-productGroupId

import (
	"context"
	"fmt"
)

const (
	CustomerInvoicePath = "/api/%s/CustomerInvoice/%s.json"
)

func NewCustomerInvoiceService(client *Client) *CustomerInvoiceService {
	return &CustomerInvoiceService{Client: client}
}

type CustomerInvoiceService struct {
	Client *Client
}

func (s *CustomerInvoiceService) Get(database string, invoiceID string, ctx context.Context) (*CustomerInvoiceGetResponse, error) {
	method := "GET"
	responseBody := s.NewGetResponse()
	path := fmt.Sprintf(CustomerInvoicePath, database, invoiceID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *CustomerInvoiceService) NewGetResponse() *CustomerInvoiceGetResponse {
	return &CustomerInvoiceGetResponse{}
}

type CustomerInvoiceGetResponse CustomerInvoice

type CustomerInvoice struct {
	AccountManager         string                `json:"accountManager"`
	AccountManagerID       string                `json:"accountManagerId"`
	Messages               []Message             `json:"messages"`
	AmountCreditSqueeze    float64               `json:"amountCreditSqueeze"`
	AmountCreditSqueezeCur float64               `json:"amountCreditSqueezeCur"`
	AmountRebate           float64               `json:"amountRebate"`
	AmountRebateCur        float64               `json:"amountRebateCur"`
	AmountTotal            float64               `json:"amountTotal"`
	AmountTotalCur         float64               `json:"amountTotalCur"`
	BordereauxNumber       string                `json:"bordereauxNumber"`
	CanChange              bool                  `json:"canChange"`
	ContactPerson          string                `json:"contactPerson"`
	ContactPersonID        string                `json:"contactPersonId"`
	CurrencyID             string                `json:"currencyId"`
	CustomerID             string                `json:"customerId"`
	CustomerInvoiceLinces  []CustomerInvoiceLine `json:"customerInvoiceLines"`
	DocumentNumber         int                   `json:"documentNumber"`
	DunForPayment          bool                  `json:"dunForPayment"`
	ExchangeRate           float64               `json:"exchangeRate"`
	FiscalYear             int                   `json:"fiscalYear"`
	InvoiceDate            DateNLNL              `json:"invoiceDate"`
	InvoiceExpirationDate  DateNLNL              `json:"invoiceExpirationDate"`
	InvoiceID              string                `json:"invoiceId"`
	InvoiceType            InvoiceType           `json:"invoiceType"`
	JournalID              string                `json:"journalId"`
	JournalSection         string                `json:"journalSection"`
	JournalTransaction     int                   `json:"journalTransaction"`
	KvcPaymentCondition    bool                  `json:"kvcPaymentCondition"`
	MandateID              string                `json:"mandateId"`
	NumberOfReminder       int                   `json:"numberOfReminders"`
	OpeningBalance         bool                  `json:"openingBalance"`
	OrderID                string                `json:"orderId"`
	PaymentConditionID     string                `json:"paymentConditionId"`
	PaymentReference       string                `json:"paymentReference"`
	PeriodNumber           int                   `json:"periodNumber"`
	ProcessedBy            string                `json:"processedBy"`
	ProcessedByID          string                `json:"processedById"`
	RebateExpirationDate   string                `json:"rebateExpirationDate"`
	Reference              string                `json:"reference"`
	SystemInvoice          bool                  `json:"systemInvoice"`
	TotalAmountVatExcl     float64               `json:"totalAmountVatExcl"`
	TotalAmountVatExclCur  float64               `json:"totalAmountVatExclCur"`
	VatAdjusted            bool                  `json:"vatAdjusted"`
	VatAmount              float64               `json:"vatAmount"`
	VatAmountCur           float64               `json:"vatAmountCur"`
	VatOnInvoice           bool                  `json:"vatOnInvoice"`
	VatScenarioID          int                   `json:"vatScenarioId"`
	VatTransactionLines    []VatTransactionLine  `json:"vatTransactionLines"`
}

type CustomerInvoiceLine struct {
	AccountID                       string        `json:"accountId"`
	Messages                        []Message     `json:"messages"`
	CanChange                       bool          `json:"canChange"`
	CostCentreID                    string        `json:"costCentreId"`
	CostCentreIDRequired            bool          `json:"costCentreIdRequired"`
	CostUnitID                      string        `json:"costUnitId"`
	CostUnitIDRequired              bool          `json:"costUnitIdRequired"`
	CreditAmount                    float64       `json:"creditAmount"`
	CreditAmountCur                 float64       `json:"creditAmountCur"`
	CurrencyID                      string        `json:"currencyId"`
	DebitAmount                     float64       `json:"debitAmount"`
	DebitAmountCur                  float64       `json:"debitAmountCur"`
	Description                     string        `json:"description"`
	DocumentNumber                  int           `json:"documentNumber"`
	ExchangeRate                    float64       `json:"exchangeRate"`
	IsSubAdminSpecificationRequired bool          `json:"isSubAdminSpecificationRequired"`
	JournalSection                  int           `json:"journalSection"`
	LineNumbers                     []int         `json:"lineNumbers"`
	Quantity                        float64       `json:"quantity"`
	SubAdminSpecifications          []interface{} `json:"subAdminSpecifications"`
	TransactionDate                 string        `json:"transactionDate"`
	VatCodeID                       int           `json:"vatCodeId"`
	VatType                         int           `json:"vatType"`
}

type VatTransactionLine struct {
	Messages               []Message `json:"messages"`
	AmountNotDeductibleCur float64   `json:"amountNotDeductibleCur"`
	AmountTurnoverCur      float64   `json:"amountTurnoverCur"`
	CanChange              bool      `json:"canChange"`
	CurrencyID             string    `json:"currencyId"`
	FiscalYear             int       `json:"fiscalYear"`
	VatAmountCur           float64   `json:"vatAmountCur"`
	VatCodeID              int       `json:"vatCodeId"`
	VatScenarioID          int       `json:"vatScenarioId"`
	VatType                int       `json:"vatType"`
}

type InvoiceType int
