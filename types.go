package vismanet

import (
	"encoding/json"
	"time"

	"github.com/cydev/zero"
	"github.com/omniboost/go-visma.net/omitempty"
)

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

type Currencies []Currency

type Currency struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ValueString string

func (v ValueString) MarshalJSON() ([]byte, error) {
	type alias ValueString
	v2 := Value{
		Value: alias(v),
	}
	return json.Marshal(v2)
}

type ValueBool bool

func (v ValueBool) MarshalJSON() ([]byte, error) {
	type alias ValueBool
	v2 := Value{
		Value: alias(v),
	}
	return json.Marshal(v2)
}

type ValueTime Time

// func (v *ValueTime) UnmarshalJSON(data []byte) error {
// 	var value string
// 	err := json.Unmarshal(data, &value)
// 	if err != nil {
// 		return err
// 	}

// 	if value == "" {
// 		return nil
// 	}

// 	// first try standard date
// 	t, err := time.Parse(time.RFC3339, value)
// 	if err == nil {
// 		vt := ValueTime(Time{t})
// 		*v = vt
// 		return nil
// 	}

// 	// try untill date format
// 	t, err = time.Parse("2006-01-02T15:04:05-07:00", value)
// 	if err == nil {
// 		vt := ValueTime(t)
// 		*v = vt
// 		return nil
// 	}

// 	return err
// }

type ValueInt int

func (v ValueInt) MarshalJSON() ([]byte, error) {
	type alias ValueInt
	v2 := Value{
		Value: alias(v),
	}
	return json.Marshal(v2)
}

type ValueNumber float64

func (v ValueNumber) MarshalJSON() ([]byte, error) {
	type alias ValueNumber
	v2 := Value{
		Value: alias(v),
	}
	return json.Marshal(v2)
}

type Value struct {
	Value interface{} `json:"value"`
}

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(t.Time.Format("2006-01-02T15:04:05-07:00"))
}

func (t *Time) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	t.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try untill date format
	t.Time, err = time.Parse("2006-01-02T15:04:05-07:00", value)
	return
}

type SalesOrderTypes []SalesOrderType

type SalesOrderType struct {
	OrderType            string `json:"orderType"`
	Active               bool   `json:"active"`
	Description          string `json:"description"`
	Behavior             string `json:"behavior"`
	DefaultOperation     string `json:"defaultOperation"`
	CustomerDocumentType string `json:"customerDocumentType"`
	Metadata             struct {
		TotalCount  int `json:"totalCount"`
		MaxPageSize int `json:"maxPageSize"`
	} `json:"metadata"`
}

type Invoice struct {
	PaymentMethodID                    ValueString    `json:"paymentMethodId,omitempty"`
	CreditTermsID                      ValueString    `json:"creditTermsId,omitempty"`
	CurrencyID                         ValueString    `json:"currencyId"`
	CustomerRefNumber                  ValueString    `json:"customerRefNumber"`
	CashDiscountDate                   ValueTime      `json:"cashDiscountDate"`
	DocumentDueDate                    ValueTime      `json:"documentDueDate"`
	ExternalReference                  ValueString    `json:"externalReference"`
	CustomerProject                    ValueString    `json:"customerProject"`
	ExchangeRate                       ValueInt       `json:"exchangeRate"`
	DomesticServicesDeductibleDocument ValueBool      `json:"domesticServicesDeductibleDocument"`
	RotRutDetails                      RotRutDetails  `json:"rotRutDetails,omitempty"`
	PaymentReference                   ValueString    `json:"paymentReference,omitempty"`
	Contact                            ValueInt       `json:"contact,omitempty"`
	Project                            ValueString    `json:"project,omitempty"`
	TaxDetailLines                     TaxDetailLines `json:"taxDetailLines,omitempty"`
	InvoiceLines                       InvoiceLines   `json:"invoiceLines"`
	SendToAutoInvoice                  ValueBool      `json:"sendToAutoInvoice"`
	CustomerVatZoneID                  ValueString    `json:"customerVatZoneId,omitempty"`
	BillingAddress                     BillingAddress `json:"billingAddress"`
	InvoiceContact                     InvoiceContact `json:"invoiceContact,omitempty"`
	StartDate                          ValueTime      `json:"startDate,omitempty"`
	EndDate                            ValueTime      `json:"endDate,omitempty"`
	AccountingCostRef                  ValueString    `json:"accountingCostRef"`
	OriginatorDocRef                   ValueString    `json:"originatorDocRef"`
	ContractDocRef                     ValueString    `json:"contractDocRef"`
	OverrideNumberSeries               ValueBool      `json:"overrideNumberSeries"`
	ReferenceNumber                    ValueString    `json:"referenceNumber"`
	CustomerNumber                     ValueString    `json:"customerNumber"`
	DocumentDate                       ValueTime      `json:"documentDate"`
	OrigInvoiceDate                    ValueTime      `json:"origInvoiceDate"`
	Hold                               ValueBool      `json:"hold"`
	PostPeriod                         ValueString    `json:"postPeriod"`
	FinancialPeriod                    ValueString    `json:"financialPeriod"`
	InvoiceText                        ValueString    `json:"invoiceText"`
	LocationID                         ValueString    `json:"locationId,omitempty"`
	SalesPersonID                      ValueString    `json:"salesPersonID,omitempty"`
	Salesperson                        ValueString    `json:"salesperson,omitempty"`
	Note                               ValueString    `json:"note"`
	BranchNumber                       ValueString    `json:"branchNumber,omitempty"`
	CashAccount                        ValueString    `json:"cashAccount,omitempty"`
	DontPrint                          ValueBool      `json:"dontPrint"`
	DontEmail                          ValueBool      `json:"dontEmail"`
}

type InvoiceContact struct {
	OverrideContact ValueBool   `json:"overrideContact"`
	Name            ValueString `json:"name"`
	Attention       ValueString `json:"attention"`
	Email           ValueString `json:"email"`
	Phone1          ValueString `json:"phone1"`
}

type BillingAddress struct {
	OverrideAddress ValueBool   `json:"overrideAddress"`
	AddressLine1    ValueString `json:"addressLine1"`
	AddressLine2    ValueString `json:"addressLine2"`
	AddressLine3    ValueString `json:"addressLine3"`
	PostalCode      ValueString `json:"postalCode"`
	City            ValueString `json:"city"`
	CountryID       ValueString `json:"countryId"`
	County          ValueString `json:"county"`
}

type InvoiceLines []InvoiceLine

type InvoiceLine struct {
	DiscountCode               ValueString `json:"discountCode,omitempty"`
	DomesticServicesDeductible ValueBool   `json:"domesticServicesDeductible"`
	ItemType                   ValueString `json:"itemType"`
	TypeOfWork                 ValueString `json:"typeOfWork"`
	TaskID                     ValueString `json:"taskId,omitempty"`
	StartDate                  ValueTime   `json:"startDate"`
	EndDate                    ValueTime   `json:"endDate"`
	Operation                  string      `json:"operation"`
	InventoryNumber            ValueString `json:"inventoryNumber,omitempty"`
	LineNumber                 ValueInt    `json:"lineNumber"`
	Description                ValueString `json:"description"`
	Quantity                   ValueInt    `json:"quantity"`
	UnitPriceInCurrency        ValueNumber `json:"unitPriceInCurrency"`
	ManualAmountInCurrency     ValueNumber `json:"manualAmountInCurrency"`
	AccountNumber              ValueString `json:"accountNumber"`
	VATCodeID                  ValueString `json:"vatCodeId,omitempty"`
	UOM                        ValueString `json:"uom,omitempty"`
	DiscountPercent            ValueNumber `json:"discountPercent"`
	DiscountAmountInCurrency   ValueNumber `json:"discountAmountInCurrency"`
	ManualDiscount             ValueBool   `json:"manualDiscount"`
	Subaccount                 []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"subaccount"`
	Salesperson      ValueString `json:"salesperson,omitempty"`
	DeferralSchedule ValueInt    `json:"deferralSchedule"`
	DeferralCode     ValueString `json:"deferralCode,omitempty"`
	TermStartDate    ValueTime   `json:"termStartDate"`
	TermEndDate      ValueTime   `json:"termEndDate"`
	Note             ValueString `json:"note"`
	BranchNumber     ValueString `json:"branchNumber,omitempty"`
}

type RotRutDetails struct {
	DistributedAutomaticaly ValueBool   `json:"distributedAutomaticaly"`
	Type                    ValueString `json:"type"`
	Appartment              ValueString `json:"appartment"`
	Estate                  ValueString `json:"estate"`
	OrganizationNbr         ValueString `json:"organizationNbr"`
	Distribution            []struct {
		Operation  string      `json:"operation"`
		LineNumber ValueInt    `json:"lineNbr"`
		PersonalID ValueString `json:"personalId"`
		Amount     ValueNumber `json:"amount"`
		Extra      ValueBool   `json:"extra"`
	} `json:"distribution"`
}

type TaxDetailLines []TaxDetailLine

func (ll TaxDetailLines) IsEmpty() bool {
	return len(ll) == 0
}

type TaxDetailLine struct {
	TaxID         ValueString `json:"taxId"`
	TaxableAmount ValueNumber `json:"taxableAmount"`
	VatAmount     ValueNumber `json:"vatAmount"`
}

type Customers []Customer

type Customer struct {
	InternalID  int    `json:"internalId"`
	Number      string `json:"number"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	MainAddress struct {
		AddressID    int    `json:"addressId"`
		AddressLine1 string `json:"addressLine1"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		Country      struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"mainAddress,omitempty"`
	MainContact struct {
		ContactID int    `json:"contactId"`
		Name      string `json:"name"`
		Email     string `json:"email"`
	} `json:"mainContact,omitempty"`
	CustomerClass struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"customerClass"`
	CreditTerms struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"creditTerms"`
	CurrencyID         string  `json:"currencyId"`
	CreditVerification string  `json:"creditVerification"`
	CreditLimit        float64 `json:"creditLimit"`
	CreditDaysPastDue  int     `json:"creditDaysPastDue"`
	InvoiceAddress     struct {
		AddressID    int    `json:"addressId"`
		AddressLine1 string `json:"addressLine1"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		Country      struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"invoiceAddress,omitempty"`
	InvoiceContact struct {
		ContactID int    `json:"contactId"`
		Name      string `json:"name"`
		Email     string `json:"email"`
	} `json:"invoiceContact,omitempty"`
	PrintInvoices                bool   `json:"printInvoices"`
	AcceptAutoInvoices           bool   `json:"acceptAutoInvoices"`
	SendInvoicesByEmail          bool   `json:"sendInvoicesByEmail"`
	PrintStatements              bool   `json:"printStatements"`
	SendStatementsByEmail        bool   `json:"sendStatementsByEmail"`
	PrintMultiCurrencyStatements bool   `json:"printMultiCurrencyStatements"`
	StatementType                string `json:"statementType"`
	DeliveryAddress              struct {
		AddressID    int    `json:"addressId"`
		AddressLine1 string `json:"addressLine1"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		Country      struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"deliveryAddress,omitempty"`
	DeliveryContact struct {
		ContactID int    `json:"contactId"`
		Name      string `json:"name"`
		Email     string `json:"email"`
	} `json:"deliveryContact,omitempty"`
	VatZone struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"vatZone"`
	Location struct {
		CountryID string `json:"countryId"`
		ID        string `json:"id"`
		Name      string `json:"name"`
	} `json:"location"`
	Attributes               []interface{} `json:"attributes"`
	LastModifiedDateTime     string        `json:"lastModifiedDateTime"`
	CreatedDateTime          string        `json:"createdDateTime"`
	InvoiceToDefaultLocation bool          `json:"invoiceToDefaultLocation"`
	PaymentMethods           []interface{} `json:"paymentMethods"`
	DefaultPaymentMethodID   string        `json:"defaultPaymentMethodId"`
	Metadata                 struct {
		TotalCount  int `json:"totalCount"`
		MaxPageSize int `json:"maxPageSize"`
	} `json:"metadata"`
}

type CustomerPostBody struct {
	Number                       ValueString `json:"number"`
	Name                         ValueString `json:"name"`
	Status                       ValueString `json:"status"`
	AccountReference             ValueString `json:"accountReference"`
	ParentRecordNumber           ValueString `json:"parentRecordNumber,omitempty"`
	CurrencyID                   ValueString `json:"currencyId,omitempty"`
	CreditLimit                  ValueInt    `json:"creditLimit,omitempty"`
	CreditDaysPastDue            ValueInt    `json:"creditDaysPastDue,omitempty"`
	OverrideWithClassValues      bool        `json:"overrideWithClassValues,omitempty"`
	CustomerClassID              ValueString `json:"customerClassId,omitempty"`
	CreditTermsID                ValueString `json:"creditTermsId,omitempty"`
	PrintInvoices                ValueBool   `json:"printInvoices,omitempty"`
	AcceptAutoInvoices           ValueBool   `json:"acceptAutoInvoices,omitempty"`
	SendInvoicesByEmail          ValueBool   `json:"sendInvoicesByEmail"`
	PrintStatements              ValueBool   `json:"printStatements,omitempty"`
	SendStatementsByEmail        ValueBool   `json:"sendStatementsByEmail"`
	PrintMultiCurrencyStatements ValueBool   `json:"printMultiCurrencyStatements,omitempty"`
	InvoiceToDefaultLocation     ValueBool   `json:"invoiceToDefaultLocation,omitempty"`
	VATRegistrationID            ValueString `json:"vatRegistrationId"`
	CorporateID                  ValueString `json:"corporateId,omitempty"`
	VATZoneID                    ValueString `json:"vatZoneId,omitempty"`
	GLN                          ValueString `json:"gln,omitempty"`
	Note                         ValueString `json:"note"`
	MainAddress                  Address     `json:"mainAddress,omitempty"`
	MainContact                  Contact     `json:"mainContact,omitempty"`
	CreditVerification           ValueString `json:"creditVerification"`
	InvoiceAddress               Address     `json:"invoiceAddress,omitempty"`
	InvoiceContact               Contact     `json:"invoiceContact,omitempty"`
	StatementType                ValueString `json:"statementType,omitempty"`
	DeliveryAddress              Address     `json:"deliveryAddress,omitempty"`
	DeliveryContact              Contact     `json:"deliveryContact,omitempty"`
	PriceClassID                 ValueString `json:"priceClassId,omitempty"`
	// EInvoiceContract struct {
	// 	Value struct {
	// 		FInvoiceContractID      ValueString `json:"fInvoiceContractID"`
	// 		FInvoiceIntermediatorID ValueString `json:"fInvoiceIntermediatorID"`
	// 	} `json:"value"`
	// } `json:"eInvoiceContract"`
	// DefaultPaymentMethod struct {
	// 	Value     ValueString `json:"paymentMethodId"`
	// 	IsDefault ValueBool   `json:"isDefault"`
	// } `json:"defaultPaymentMethod"`
	GlAccounts CustomerGLAccounts `json:"glAccounts,omitempty"`
	// DirectDebitLines DirectDebitLines   `json:"directDebitLines,omitempty"`
	// AttributeLines   []struct {
	// 	AttributeID    ValueString `json:"attributeId"`
	// 	AttributeValue ValueString `json:"attributeValue"`
	// } `json:"attributeLines"`
	// OverrideNumberSeries ValueBool `json:"overrideNumberSeries"`
}

func (r CustomerPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type Address struct {
	AddressLine1 ValueString `json:"addressLine1"`
	AddressLine2 ValueString `json:"addressLine2"`
	AddressLine3 ValueString `json:"addressLine3"`
	PostalCode   ValueString `json:"postalCode"`
	City         ValueString `json:"city"`
	CountryID    ValueString `json:"countryId"`
	County       ValueString `json:"county"`
}

func (a Address) MarshalJSON() ([]byte, error) {
	type alias Address
	a2 := Value{
		Value: alias(a),
	}
	return json.Marshal(a2)
}

func (a Address) IsEmpty() bool {
	return zero.IsZero(a)
}

type Contact struct {
	Name      ValueString `json:"name"`
	Attention ValueString `json:"attention"`
	Email     ValueString `json:"email"`
	Web       ValueString `json:"web"`
	Phone1    ValueString `json:"phone1"`
	Phone2    ValueString `json:"phone2"`
	Fax       ValueString `json:"fax"`
}

func (c Contact) MarshalJSON() ([]byte, error) {
	type alias Contact
	c2 := Value{
		Value: alias(c),
	}
	return json.Marshal(c2)
}

func (c Contact) IsEmpty() bool {
	return zero.IsZero(c)
}

type CustomerGLAccounts struct {
	CustomerLedgerAccount    ValueString `json:"customerLedgerAccount"`
	CustomerLedgerSubaccount []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"customerLedgerSubaccount"`
	SalesAccount           ValueString `json:"salesAccount"`
	SalesNonTaxableAccount ValueString `json:"salesNonTaxableAccount"`
	SalesEuAccount         ValueString `json:"salesEuAccount"`
	SalesExportAccount     ValueString `json:"salesExportAccount"`
	SalesSubaccount        []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"salesSubaccount"`
	DiscountAccount    ValueString `json:"discountAccount"`
	DiscountSubaccount []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"discountSubaccount"`
	FreightAccount    ValueString `json:"freightAccount"`
	FreightSubaccount []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"freightSubaccount"`
	CashDiscountAccount    ValueString `json:"cashDiscountAccount"`
	CashDiscountSubaccount []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"cashDiscountSubaccount"`
	PrepaymentAccount    ValueString `json:"prepaymentAccount"`
	PrepaymentSubaccount []struct {
		SegmentID    int    `json:"segmentId"`
		SegmentValue string `json:"segmentValue"`
	} `json:"prepaymentSubaccount"`
}

func (c CustomerGLAccounts) MarshalJSON() ([]byte, error) {
	type alias CustomerGLAccounts
	v := Value{
		Value: alias(c),
	}
	return json.Marshal(v)
}

func (c CustomerGLAccounts) IsEmpty() bool {
	return zero.IsZero(c)
}

type DirectDebitLines []DirectDebitLine

func (dl DirectDebitLines) IsEmpty() bool {
	return len(dl) == 0
}

type DirectDebitLine struct {
	Operation          string      `json:"operation"`
	ID                 string      `json:"id"`
	MandateID          ValueString `json:"mandateId"`
	MandateDescription ValueString `json:"mandateDescription"`
	DateOfSignature    ValueTime   `json:"dateOfSignature"`
	IsDefault          ValueBool   `json:"isDefault"`
	OneTime            ValueBool   `json:"oneTime"`
	Bic                ValueString `json:"bic"`
	Iban               ValueString `json:"iban"`
	LastCollectionDate ValueTime   `json:"lastCollectionDate"`
	MaxAmount          ValueInt    `json:"maxAmount"`
	ExpirationDate     ValueTime   `json:"expirationDate"`
}
