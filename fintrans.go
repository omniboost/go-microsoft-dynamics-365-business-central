package multivers

// https://api.online.unit4.nl/V16/Help/Api/GET-api-database-ProductGroup-productGroupId

import (
	"context"
	"fmt"
)

const (
	FinTransPath = "/api/%s/FinTrans"
)

func NewFinTransService(client *Client) *FinTransService {
	return &FinTransService{Client: client}
}

type FinTransService struct {
	Client *Client
}

func (s *FinTransService) Post(database string, finTrans FinTrans, ctx context.Context) (*FinTransGetResponse, error) {
	method := "POST"
	responseBody := s.NewGetResponse()
	path := fmt.Sprintf(FinTransPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, finTrans)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *FinTransService) NewGetResponse() *FinTransGetResponse {
	return &FinTransGetResponse{}
}

type FinTransGetResponse FinTrans

type FinTrans struct {
	JournalID          string          `json:"journalId"`
	FiscalYear         int             `json:"fiscalYear"`
	JournalTransaction int             `json:"journalTransaction,omitempty"`
	PeriodNumber       int             `json:"periodNumber"`
	TransactionDate    string          `json:"transactionDate"`
	Description        string          `json:"description"`
	Document           string          `json:"document"`
	FinTransEntries    FinTransEntries `json:"finTransEntries"`
}

type FinTransEntries []FinTransEntry

type FinTransEntry struct {
	Type                       string  `json:"$type,omitempty"`                      // "UNIT4.Multivers.API.BL.Financial.Edit.BasicEntryProxy, UNIT4.Multivers.API.Web.WebApi.Model",
	AccountID                  string  `json:"accountId,omitempty"`                  // "4500",
	CostCentreID               string  `json:"costCentreId,omitempty"`               //"",
	CostCentreIDRequired       bool    `json:"costCentreIdRequired,omitempty"`       //false
	CostUnitID                 string  `json:"costUnitId,omitempty"`                 //""
	CostUnitIDRequired         bool    `json:"costUnitIdRequired,omitempty"`         //false,
	CreditAmount               float64 `json:"creditAmount"`                         //100.00,
	DebitAmount                float64 `json:"debitAmount"`                          //0.0,
	Description                string  `json:"description,omitempty"`                //"Test Basicentry",
	Document                   string  `json:"document,omitempty"`                   //"",
	NonDeductibleVatPercentage float64 `json:"nondeductibleVatPercentage,omitempty"` // 0.0,
	Quantity                   float64 `json:"quantity,omitempty"`                   // 0.0,
	TransactionDate            string  `json:"transactionDate,omitempty"`            // "28-2-2018",
	VatAmount                  float64 `json:"vatAmount,omitempty"`                  // 17.36,
	VatChanged                 bool    `json:"vatChanged,omitempty"`                 // false,
	VatCodeID                  int     `json:"vatCodeId,omitempty"`                  // 2,
	VatIncluded                bool    `json:"vatIncluded,omitempty"`                // true,
	VatScenarioID              int     `json:"vatScenarioId,omitempty"`              // 6,
	VatType                    int     `json:"vatType,omitempty"`                    // 0
}
