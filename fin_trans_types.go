package multivers

import "github.com/shopspring/decimal"

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
	Type                       string          `json:"$type,omitempty"`                      // "UNIT4.Multivers.API.BL.Financial.Edit.BasicEntryProxy, UNIT4.Multivers.API.Web.WebApi.Model",
	AccountID                  string          `json:"accountId,omitempty"`                  // "4500",
	CostCentreID               string          `json:"costCentreId,omitempty"`               //"",
	CostCentreIDRequired       bool            `json:"costCentreIdRequired,omitempty"`       //false
	CostUnitID                 string          `json:"costUnitId,omitempty"`                 //""
	CostUnitIDRequired         bool            `json:"costUnitIdRequired,omitempty"`         //false,
	CreditAmount               decimal.Decimal `json:"creditAmount"`                         //100.00,
	DebitAmount                decimal.Decimal `json:"debitAmount"`                          //0.0,
	Description                string          `json:"description,omitempty"`                //"Test Basicentry",
	Document                   string          `json:"document,omitempty"`                   //"",
	NonDeductibleVATPercentage float64         `json:"nondeductibleVATPercentage,omitempty"` // 0.0,
	Quantity                   float64         `json:"quantity,omitempty"`                   // 0.0,
	TransactionDate            string          `json:"transactionDate,omitempty"`            // "28-2-2018",
	VATAmount                  decimal.Decimal `json:"vatAmount,omitempty"`                  // 17.36,
	VATChanged                 bool            `json:"vatChanged,omitempty"`                 // false,
	VATCodeID                  int             `json:"vatCodeId,omitempty"`                  // 2,
	VATIncluded                bool            `json:"vatIncluded,omitempty"`                // true,
	VATScenarioID              int             `json:"vatScenarioId,omitempty"`              // 6,
	VATType                    int             `json:"vatType,omitempty"`                    // 0
}
