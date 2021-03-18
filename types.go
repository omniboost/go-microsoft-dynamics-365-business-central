package vismanet

type Accounts []Account

type Account struct {
	AccountID            int    `json:"accountID"`
	AccountCD            string `json:"accountCD"`
	AccountClass         string `json:"accountClass"`
	Type                 string `json:"type"`
	Active               bool   `json:"active"`
	Description          string `json:"description"`
	UseDefaultSub        bool   `json:"useDefaultSub"`
	PostOption           string `json:"postOption"`
	CashAccount          bool   `json:"cashAccount"`
	LastModifiedDateTime string `json:"lastModifiedDateTime"`
	Currency             string `json:"currency,omitempty"`
}

type VATCategories []VATCategory

type VATCategory struct {
	VATCategoryID        string           `json:"vatCategoryId"`
	Description          string           `json:"description"`
	Active               bool             `json:"active"`
	ExcludeListedTaxes   bool             `json:"excludeListedTaxes"`
	LastModifiedDateTime string           `json:"lastModifiedDateTime"`
	VATCategoryLines     VATCategoryLines `json:"vatCategoryLines"`
}

type VATCategoryLines []VATCategoryLine

type VATCategoryLine struct {
	VATID        string   `json:"vatId"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	CalculateOn  string   `json:"calculateOn"`
	CashDiscount string   `json:"cashDiscount"`
	VATRates     VATRates `json:"vatRates"`
}

type VATRates []VATRate

type VATRate struct {
	RevisionID string  `json:"revisionID"`
	VATRate    float64 `json:"vatRate"`
	StartDate  string  `json:"startDate"`
}

type JournalTransactions []JournalTransaction

type JournalTransaction struct {
	Module                  string                  `json:"module"`
	BatchNumber             string                  `json:"batchNumber"`
	Status                  string                  `json:"status"`
	Hold                    bool                    `json:"hold"`
	TransactionDate         string                  `json:"transactionDate"`
	PostPeriod              string                  `json:"postPeriod"`
	FinancialPeriod         string                  `json:"financialPeriod"`
	Ledger                  string                  `json:"ledger"`
	LedgerDescription       string                  `json:"ledgerDescription"`
	CurrencyID              string                  `json:"currencyId"`
	ExchangeRate            float64                 `json:"exchangeRate"`
	AutoReversing           bool                    `json:"autoReversing"`
	ReversingEntry          bool                    `json:"reversingEntry"`
	Description             string                  `json:"description"`
	DebitTotal              float64                 `json:"debitTotal"`
	DebitTotalInCurrency    float64                 `json:"debitTotalInCurrency"`
	CreditTotal             float64                 `json:"creditTotal"`
	CreditTotalInCurrency   float64                 `json:"creditTotalInCurrency"`
	ControlTotal            float64                 `json:"controlTotal"`
	ControlTotalInCurrency  float64                 `json:"controlTotalInCurrency"`
	CreateVatTransaction    bool                    `json:"createVatTransaction"`
	SkipVatAmountValidation bool                    `json:"skipVatAmountValidation"`
	LastModifiedDateTime    string                  `json:"lastModifiedDateTime"`
	Branch                  string                  `json:"branch"`
	JournalTransactionLines JournalTransactionLines `json:"journalTransactionLines"`
	Metadata                Metadata                `json:"metadata"`
}

type JournalTransactionLines []JournalTransactionLine

type JournalTransactionLine struct {
	LineNumber             int        `json:"lineNumber"`
	AccountNumber          string     `json:"accountNumber"`
	Description            string     `json:"description"`
	Subaccount             Subaccount `json:"subaccount"`
	ReferenceNumber        string     `json:"referenceNumber"`
	DebitAmount            float64    `json:"debitAmount"`
	DebitAmountInCurrency  float64    `json:"debitAmountInCurrency"`
	CreditAmount           float64    `json:"creditAmount"`
	CreditAmountInCurrency float64    `json:"creditAmountInCurrency"`
	TransactionDescription string     `json:"transactionDescription"`
	Branch                 string     `json:"branch"`
	TransactionType        string     `json:"transactionType"`
	Module                 string     `json:"module"`
	Project                struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"project"`
	Quantity float64 `json:"quantity"`
}

type Metadata struct {
	TotalCount  int `json:"totalCount"`
	MaxPageSize int `json:"maxPageSize"`
}

type Subaccount struct {
	SubaccountNumber     string   `json:"subaccountNumber"`
	SubaccountID         int      `json:"subaccountId"`
	Description          string   `json:"description"`
	LastModifiedDateTime string   `json:"lastModifiedDateTime"`
	Active               bool     `json:"active"`
	Segments             Segments `json:"segments"`
}

type Segments []Segment

type Segment struct {
	SegmentID               int    `json:"segmentId"`
	SegmentDescription      string `json:"segmentDescription"`
	SegmentValue            string `json:"segmentValue"`
	SegmentValueDescription string `json:"segmentValueDescription"`
}

type Ledgers []Ledger

type Ledger struct {
	InternalID           int    `json:"internalId"`
	Number               string `json:"number"`
	Description          string `json:"description"`
	BalanceType          string `json:"balanceType"`
	CurrencyID           string `json:"currencyId"`
	ConsolidationSource  bool   `json:"consolidationSource"`
	BranchAccounting     bool   `json:"branchAccounting"`
	LastModifiedDateTime string `json:"lastModifiedDateTime"`
	PostInterCompany     bool   `json:"postInterCompany"`
}
