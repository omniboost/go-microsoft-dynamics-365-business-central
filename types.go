package vismaonline

import (
	"github.com/cydev/zero"
	"github.com/omniboost/go-vismaonline/omitempty"
)

type Meta struct {
	CurrentPage          int      `json:"CurrentPage"`
	PageSize             int      `json:"PageSize"`
	TotalNumberOfPages   int      `json:"TotalNumberOfPages"`
	TotalNumberOfResults int      `json:"TotalNumberOfResults"`
	ServerTimeUTC        DateTime `json:"ServerTimeUTC"`
}

type Accounts []Account

type Account struct {
	Name                      string   `json:"Name"`
	Number                    string   `json:"Number"`
	VATCodeID                 int      `json:"VatCodeId"`
	VATCodeDescription        string   `json:"VatCodeDescription"`
	FiscalYearID              string   `json:"FiscalYearId"`
	ReferenceCode             string   `json:"ReferenceCode"`
	Type                      int      `json:"Type"`
	TypeDescription           string   `json:"TypeDescription"`
	ModifiedUTC               DateTime `json:"ModifiedUTC"`
	IsActive                  bool     `json:"IsActive"`
	IsProjectAllowed          bool     `json:"IsProjectAllowed"`
	IsCostCenterAllowed       bool     `json:"IsCostCenterAllowed"`
	IsBlockedForManualBooking bool     `json:"IsBlockedForManualBooking"`
}

type Companysettings struct {
	Name                    string      `json:"Name"`
	Email                   string      `json:"Email"`
	Phone                   string      `json:"Phone"`
	MobilePhone             string      `json:"MobilePhone"`
	Address1                string      `json:"Address1"`
	Address2                string      `json:"Address2"`
	CountryCode             string      `json:"CountryCode"`
	PostalCode              string      `json:"PostalCode"`
	City                    string      `json:"City"`
	Website                 interface{} `json:"Website"`
	OurReference            interface{} `json:"OurReference"`
	CurrencyCode            string      `json:"CurrencyCode"`
	TermsOfPaymentID        string      `json:"TermsOfPaymentId"`
	CorporateIdentityNumber string      `json:"CorporateIdentityNumber"`
	VATCode                 string      `json:"VatCode"`
	BankGiro                string      `json:"BankGiro"`
	PlusGiro                string      `json:"PlusGiro"`
	BankAccount             interface{} `json:"BankAccount"`
	Iban                    string      `json:"Iban"`
	AccountNumberDigits     int         `json:"AccountNumberDigits"`
	AccountingLockedTo      DateTime    `json:"AccountingLockedTo"`
	AccountingLockInterval  int         `json:"AccountingLockInterval"`
	TaxDeclarationDate      Date        `json:"TaxDeclarationDate"`
	Gln                     interface{} `json:"Gln"`
	ProductVariant          int         `json:"ProductVariant"`
	TypeOfBusiness          int         `json:"TypeOfBusiness"`
	VatPeriod               int         `json:"VatPeriod"`
	ActivatedModules        []string    `json:"ActivatedModules"`
	CompanyText             struct {
		CustomerInvoiceTextDomestic interface{} `json:"CustomerInvoiceTextDomestic"`
		CustomerInvoiceTextForeign  interface{} `json:"CustomerInvoiceTextForeign"`
		OrderTextDomestic           interface{} `json:"OrderTextDomestic"`
		OrderTextForeign            interface{} `json:"OrderTextForeign"`
		OverDueTextDomestic         string      `json:"OverDueTextDomestic"`
		OverDueTextForeign          string      `json:"OverDueTextForeign"`
	} `json:"CompanyText"`
	NextCustomerNumber         int  `json:"NextCustomerNumber"`
	NextSupplierNumber         int  `json:"NextSupplierNumber"`
	NextCustomerInvoiceNumber  int  `json:"NextCustomerInvoiceNumber"`
	NextQuoteNumber            int  `json:"NextQuoteNumber"`
	ShowPricesExclVatPC        bool `json:"ShowPricesExclVatPC"`
	IsPayslipActivated         bool `json:"IsPayslipActivated"`
	UsesMoss                   bool `json:"UsesMoss"`
	UsesPaymentReferenceNumber bool `json:"UsesPaymentReferenceNumber"`
	DomesticCurrencyRounding   int  `json:"DomesticCurrencyRounding"`
	AutoInvoice                struct {
		AutoInvoiceActivationStatus int `json:"AutoInvoiceActivationStatus"`
		AutoInvoiceB2CStatus        int `json:"AutoInvoiceB2CStatus"`
		AutoInvoiceInboundStatus    int `json:"AutoInvoiceInboundStatus"`
	} `json:"AutoInvoice"`
	ApprovalSettings struct {
		UsesSupplierInvoiceApproval bool `json:"UsesSupplierInvoiceApproval"`
		UsesVatReportApproval       bool `json:"UsesVatReportApproval"`
	} `json:"ApprovalSettings"`
	UsesReverseConstructionVat bool `json:"UsesReverseConstructionVat"`
	UsesRotReducedInvoicing    bool `json:"UsesRotReducedInvoicing"`
	CompanyRotRutSettings      struct {
		RutMaxAmountForPersBelow65Year float64 `json:"RutMaxAmountForPersBelow65Year"`
		RutMaxAmountForPersOver65Year  float64 `json:"RutMaxAmountForPersOver65Year"`
		RutReducedInvoicingPercent     float64 `json:"RutReducedInvoicingPercent"`
		RotReducedInvoicingMaxAmount   float64 `json:"RotReducedInvoicingMaxAmount"`
		RotReducedInvoicingPercent     float64 `json:"RotReducedInvoicingPercent"`
	} `json:"CompanyRotRutSettings"`
	CompanyStatus                      int         `json:"CompanyStatus"`
	CompanyIdentifier                  string      `json:"CompanyIdentifier"`
	BankgiroNumberPrint                interface{} `json:"BankgiroNumberPrint"`
	KeepOriginalDraftDate              Date        `json:"KeepOriginalDraftDate"`
	UsePaymentFilesForOutgoingPayments bool        `json:"UsePaymentFilesForOutgoingPayments"`
	UseAutomaticVatCalculation         bool        `json:"UseAutomaticVatCalculation"`
	ShowCostCenterReminder             bool        `json:"ShowCostCenterReminder"`
	ShowProjectReminder                bool        `json:"ShowProjectReminder"`
}

type VATCodes []VATCode

type VATCode struct {
	ID              string  `json:"Id"`
	Code            string  `json:"Code"`
	Description     string  `json:"Description"`
	VatRate         float64 `json:"VatRate"`
	OssCodeType     int     `json:"OssCodeType"`
	RelatedAccounts struct {
		AccountNumber1 int         `json:"AccountNumber1"`
		AccountNumber2 interface{} `json:"AccountNumber2"`
		AccountNumber3 interface{} `json:"AccountNumber3"`
	} `json:"RelatedAccounts"`
}

type Vouchers []Voucher

type Voucher struct {
	ID                    string      `json:"Id,omitempty,omitempty"`
	VoucherDate           Date        `json:"VoucherDate,omitempty"`
	VoucherText           string      `json:"VoucherText,omitempty"`
	Rows                  VoucherRows `json:"Rows,omitempty"`
	NumberAndNumberSeries string      `json:"NumberAndNumberSeries,omitempty"`
	NumberSeries          string      `json:"NumberSeries,omitempty"`
	Attachments           interface{} `json:"Attachments,omitempty"`
	ModifiedUTC           DateTime    `json:"ModifiedUTC,omitempty"`
	VoucherType           int         `json:"VoucherType,omitempty"`
	SourceID              string      `json:"SourceId,omitempty"`
	CreatedUTC            DateTime    `json:"CreatedUTC,omitempty"`
}

func (v Voucher) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

type VoucherRows []VoucherRow

type VoucherRow struct {
	AccountNumber      int         `json:"AccountNumber,omitempty"`
	AccountDescription string      `json:"AccountDescription,omitempty"`
	DebitAmount        float64     `json:"DebitAmount,omitempty"`
	CreditAmount       float64     `json:"CreditAmount,omitempty"`
	TransactionText    string      `json:"TransactionText,omitempty"`
	CostCenterItemID1  interface{} `json:"CostCenterItemId1,omitempty"`
	CostCenterItemID2  interface{} `json:"CostCenterItemId2,omitempty"`
	CostCenterItemID3  interface{} `json:"CostCenterItemId3,omitempty"`
	VATCodeID          string      `json:"VatCodeId,omitempty"`
	VATCodeAndPercent  string      `json:"VatCodeAndPercent,omitempty"`
	VATAmount          float64     `json:"VatAmount,omitempty"`
	Quantity           int         `json:"Quantity,omitempty"`
	Weight             float64     `json:"Weight,omitempty"`
	DeliveryDate       Date        `json:"DeliveryDate,omitempty"`
	HarvestYear        interface{} `json:"HarvestYear,omitempty"`
	ProjectID          interface{} `json:"ProjectId,omitempty"`
}

func (v VoucherRow) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

type Customers []Customer

type Customer struct {
	ID                                     string         `json:"Id,omitempty"`
	CustomerNumber                         string         `json:"CustomerNumber,omitempty"`
	CorporateIdentityNumber                string         `json:"CorporateIdentityNumber,omitempty"`
	ContactPersonEmail                     string         `json:"ContactPersonEmail,omitempty"`
	ContactPersonMobile                    string         `json:"ContactPersonMobile,omitempty"`
	ContactPersonName                      string         `json:"ContactPersonName,omitempty"`
	ContactPersonPhone                     string         `json:"ContactPersonPhone,omitempty"`
	CurrencyCode                           string         `json:"CurrencyCode,omitempty"`
	GLN                                    interface{}    `json:"GLN,omitempty"`
	EmailAddress                           string         `json:"EmailAddress,omitempty"`
	InvoiceAddress1                        string         `json:"InvoiceAddress1,omitempty"`
	InvoiceAddress2                        string         `json:"InvoiceAddress2,omitempty"`
	InvoiceCity                            string         `json:"InvoiceCity,omitempty"`
	InvoiceCountryCode                     string         `json:"InvoiceCountryCode,omitempty"`
	InvoicePostalCode                      string         `json:"InvoicePostalCode,omitempty"`
	DeliveryCustomerName                   interface{}    `json:"DeliveryCustomerName,omitempty"`
	DeliveryAddress1                       interface{}    `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2                       interface{}    `json:"DeliveryAddress2,omitempty"`
	DeliveryCity                           interface{}    `json:"DeliveryCity,omitempty"`
	DeliveryCountryCode                    interface{}    `json:"DeliveryCountryCode,omitempty"`
	DeliveryPostalCode                     interface{}    `json:"DeliveryPostalCode,omitempty"`
	DeliveryMethodID                       interface{}    `json:"DeliveryMethodId,omitempty"`
	DeliveryTermID                         interface{}    `json:"DeliveryTermId,omitempty"`
	PayToAccountID                         string         `json:"PayToAccountId,omitempty"`
	Name                                   string         `json:"Name,omitempty"`
	Note                                   string         `json:"Note,omitempty"`
	ReverseChargeOnConstructionServices    bool           `json:"ReverseChargeOnConstructionServices,omitempty"`
	WebshopCustomerNumber                  interface{}    `json:"WebshopCustomerNumber,omitempty"`
	MobilePhone                            string         `json:"MobilePhone,omitempty"`
	Telephone                              string         `json:"Telephone,omitempty"`
	TermsOfPaymentID                       string         `json:"TermsOfPaymentId,omitempty"`
	TermsOfPayment                         TermsOfPayment `json:"TermsOfPayment,omitempty"`
	VatNumber                              string         `json:"VatNumber,omitempty"`
	WwwAddress                             string         `json:"WwwAddress,omitempty"`
	LastInvoiceDate                        Date           `json:"LastInvoiceDate,omitempty"`
	IsPrivatePerson                        bool           `json:"IsPrivatePerson"`
	IsNorthernIreland                      bool           `json:"IsNorthernIreland,omitempty"`
	DiscountPercentage                     float64        `json:"DiscountPercentage,omitempty"`
	ChangedUtc                             DateTime       `json:"ChangedUtc,omitempty"`
	IsActive                               bool           `json:"IsActive,omitempty"`
	ForceBookkeepVat                       bool           `json:"ForceBookkeepVat,omitempty"`
	EdiGlnNumber                           string         `json:"EdiGlnNumber,omitempty"`
	SalesDocumentLanguage                  string         `json:"SalesDocumentLanguage,omitempty"`
	ElectronicAddress                      interface{}    `json:"ElectronicAddress,omitempty"`
	ElectronicReference                    interface{}    `json:"ElectronicReference,omitempty"`
	EdiServiceDelivererID                  interface{}    `json:"EdiServiceDelivererId,omitempty"`
	AutoInvoiceActivationEmailSentDate     interface{}    `json:"AutoInvoiceActivationEmailSentDate,omitempty"`
	AutoInvoiceRegistrationRequestSentDate interface{}    `json:"AutoInvoiceRegistrationRequestSentDate,omitempty"`
	EmailAddresses                         []interface{}  `json:"EmailAddresses,omitempty"`
	CustomerLabels                         interface{}    `json:"CustomerLabels,omitempty"`
	IsFutureInvoiceDateAllowed             bool           `json:"IsFutureInvoiceDateAllowed,omitempty"`
	DeliveryBasedVat                       bool           `json:"DeliveryBasedVat,omitempty"`
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

func (c Customer) IsEmpty() bool {
	return zero.IsZero(c)
}

type TermsOfPayments []TermsOfPayment

type TermsOfPayment struct {
	ID                     string      `json:"Id"`
	Name                   string      `json:"Name"`
	NameEnglish            string      `json:"NameEnglish"`
	NumberOfDays           int         `json:"NumberOfDays"`
	TermsOfPaymentTypeID   int         `json:"TermsOfPaymentTypeId"`
	TermsOfPaymentTypeText interface{} `json:"TermsOfPaymentTypeText"`
	AvailableForSales      bool        `json:"AvailableForSales"`
	AvailableForPurchase   bool        `json:"AvailableForPurchase"`
}

func (t TermsOfPayment) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(t)
}

func (t TermsOfPayment) IsEmpty() bool {
	return zero.IsZero(t)
}

type Articles []Article

type Article struct {
	ID                              string        `json:"Id"`
	IsActive                        bool          `json:"IsActive"`
	Number                          string        `json:"Number"`
	Name                            string        `json:"Name"`
	NameEnglish                     interface{}   `json:"NameEnglish"`
	NetPrice                        float64       `json:"NetPrice"`
	GrossPrice                      float64       `json:"GrossPrice"`
	CodingID                        string        `json:"CodingId"`
	CodingName                      string        `json:"CodingName"`
	UnitID                          string        `json:"UnitId"`
	UnitName                        string        `json:"UnitName"`
	UnitAbbreviation                string        `json:"UnitAbbreviation"`
	StockBalance                    float64       `json:"StockBalance"`
	StockBalanceManuallyChangedUtc  interface{}   `json:"StockBalanceManuallyChangedUtc"`
	StockBalanceReserved            float64       `json:"StockBalanceReserved"`
	StockBalanceAvailable           float64       `json:"StockBalanceAvailable"`
	CreatedUtc                      DateTime      `json:"CreatedUtc"`
	ChangedUtc                      DateTime      `json:"ChangedUtc"`
	HouseWorkType                   int           `json:"HouseWorkType"`
	PurchasePrice                   float64       `json:"PurchasePrice"`
	PurchasePriceManuallyChangedUtc interface{}   `json:"PurchasePriceManuallyChangedUtc"`
	SendToWebshop                   bool          `json:"SendToWebshop"`
	UsesMoss                        bool          `json:"UsesMoss"`
	ArticleLabels                   []interface{} `json:"ArticleLabels"`
	IsStock                         bool          `json:"IsStock"`
	IsServiceArticle                bool          `json:"IsServiceArticle"`
	StockLocationReference          string        `json:"StockLocationReference"`
	FreightCosts                    float64       `json:"FreightCosts"`
	FreightCostsManuallyChangedUtc  interface{}   `json:"FreightCostsManuallyChangedUtc"`
	UpdateStockPrices               bool          `json:"UpdateStockPrices"`
	Barcodes                        interface{}   `json:"Barcodes"`
	StockValue                      float64       `json:"StockValue"`
	GreenTechnologyType             int           `json:"GreenTechnologyType"`
}

func (a Article) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(a)
}

func (a Article) IsEmpty() bool {
	return zero.IsZero(a)
}

type CustomerInvoices []CustomerInvoice

type CustomerInvoice struct {
	ID                            string              `json:"Id,omitempty"`
	EuThirdParty                  bool                `json:"EuThirdParty"`
	IsCreditInvoice               bool                `json:"IsCreditInvoice,omitempty"`
	CurrencyCode                  string              `json:"CurrencyCode,omitempty"`
	CurrencyRate                  float64             `json:"CurrencyRate,omitempty"`
	CreatedByUserID               string              `json:"CreatedByUserId,omitempty"`
	TotalAmount                   float64             `json:"TotalAmount,omitempty"`
	TotalVatAmount                float64             `json:"TotalVatAmount,omitempty"`
	TotalRoundings                float64             `json:"TotalRoundings,omitempty"`
	TotalAmountInvoiceCurrency    float64             `json:"TotalAmountInvoiceCurrency,omitempty"`
	TotalVatAmountInvoiceCurrency float64             `json:"TotalVatAmountInvoiceCurrency,omitempty"`
	SetOffAmountInvoiceCurrency   float64             `json:"SetOffAmountInvoiceCurrency,omitempty"`
	CustomerID                    string              `json:"CustomerId,omitempty"`
	Rows                          CustomerInvoiceRows `json:"Rows,omitempty"`
	VatSpecification              []struct {
		AmountInvoiceCurrency    float64 `json:"AmountInvoiceCurrency,omitempty"`
		VatAmountInvoiceCurrency float64 `json:"VatAmountInvoiceCurrency,omitempty"`
		VatPercent               float64 `json:"VatPercent,omitempty"`
	} `json:"VatSpecification,omitempty"`
	InvoiceDate                              Date          `json:"InvoiceDate,omitempty"`
	DueDate                                  Date          `json:"DueDate,omitempty"`
	DeliveryDate                             Date          `json:"DeliveryDate,omitempty"`
	RotReducedInvoicingType                  int           `json:"RotReducedInvoicingType"`
	RotReducedInvoicingAmount                float64       `json:"RotReducedInvoicingAmount,omitempty"`
	RotReducedInvoicingPercent               float64       `json:"RotReducedInvoicingPercent,omitempty"`
	RotReducedInvoicingPropertyName          interface{}   `json:"RotReducedInvoicingPropertyName,omitempty"`
	RotReducedInvoicingOrgNumber             interface{}   `json:"RotReducedInvoicingOrgNumber,omitempty"`
	Persons                                  []interface{} `json:"Persons,omitempty"`
	RotReducedInvoicingAutomaticDistribution bool          `json:"RotReducedInvoicingAutomaticDistribution,omitempty"`
	ElectronicReference                      interface{}   `json:"ElectronicReference,omitempty"`
	ElectronicAddress                        interface{}   `json:"ElectronicAddress,omitempty"`
	EdiServiceDelivererID                    interface{}   `json:"EdiServiceDelivererId,omitempty"`
	OurReference                             interface{}   `json:"OurReference,omitempty"`
	YourReference                            string        `json:"YourReference,omitempty"`
	BuyersOrderReference                     string        `json:"BuyersOrderReference,omitempty"`
	InvoiceCustomerName                      string        `json:"InvoiceCustomerName,omitempty"`
	InvoiceAddress1                          string        `json:"InvoiceAddress1,omitempty"`
	InvoiceAddress2                          string        `json:"InvoiceAddress2,omitempty"`
	InvoicePostalCode                        string        `json:"InvoicePostalCode,omitempty"`
	InvoiceCity                              string        `json:"InvoiceCity,omitempty"`
	InvoiceCountryCode                       string        `json:"InvoiceCountryCode,omitempty"`
	DeliveryCustomerName                     interface{}   `json:"DeliveryCustomerName,omitempty"`
	DeliveryAddress1                         interface{}   `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2                         interface{}   `json:"DeliveryAddress2,omitempty"`
	DeliveryPostalCode                       interface{}   `json:"DeliveryPostalCode,omitempty"`
	DeliveryCity                             interface{}   `json:"DeliveryCity,omitempty"`
	DeliveryCountryCode                      interface{}   `json:"DeliveryCountryCode,omitempty"`
	DeliveryMethodName                       interface{}   `json:"DeliveryMethodName,omitempty"`
	DeliveryTermName                         interface{}   `json:"DeliveryTermName,omitempty"`
	DeliveryMethodCode                       interface{}   `json:"DeliveryMethodCode,omitempty"`
	DeliveryTermCode                         interface{}   `json:"DeliveryTermCode,omitempty"`
	CustomerIsPrivatePerson                  bool          `json:"CustomerIsPrivatePerson,omitempty"`
	TermsOfPaymentID                         string        `json:"TermsOfPaymentId,omitempty"`
	CustomerEmail                            string        `json:"CustomerEmail,omitempty"`
	InvoiceNumber                            int           `json:"InvoiceNumber,omitempty"`
	CustomerNumber                           string        `json:"CustomerNumber,omitempty"`
	PaymentReferenceNumber                   interface{}   `json:"PaymentReferenceNumber,omitempty"`
	RotPropertyType                          interface{}   `json:"RotPropertyType,omitempty"`
	SalesDocumentAttachments                 []interface{} `json:"SalesDocumentAttachments,omitempty"`
	MessageThreads                           []interface{} `json:"MessageThreads,omitempty"`
	Notes                                    []interface{} `json:"Notes,omitempty"`
	HasAutoInvoiceError                      bool          `json:"HasAutoInvoiceError,omitempty"`
	IsNotDelivered                           bool          `json:"IsNotDelivered,omitempty"`
	ReverseChargeOnConstructionServices      bool          `json:"ReverseChargeOnConstructionServices,omitempty"`
	WorkHouseOtherCosts                      interface{}   `json:"WorkHouseOtherCosts,omitempty"`
	RemainingAmount                          float64       `json:"RemainingAmount,omitempty"`
	RemainingAmountInvoiceCurrency           float64       `json:"RemainingAmountInvoiceCurrency,omitempty"`
	ReferringInvoiceID                       interface{}   `json:"ReferringInvoiceId,omitempty"`
	CreatedFromOrderID                       interface{}   `json:"CreatedFromOrderId,omitempty"`
	CreatedFromDraftID                       interface{}   `json:"CreatedFromDraftId,omitempty"`
	VoucherNumber                            string        `json:"VoucherNumber,omitempty"`
	VoucherID                                string        `json:"VoucherId,omitempty"`
	CreatedUtc                               DateTime      `json:"CreatedUtc,omitempty"`
	ModifiedUtc                              DateTime      `json:"ModifiedUtc,omitempty"`
	ReversedConstructionVatInvoicing         bool          `json:"ReversedConstructionVatInvoicing,omitempty"`
	IncludesVat                              bool          `json:"IncludesVat,omitempty"`
	SendType                                 interface{}   `json:"SendType,omitempty"`
	PaymentReminderIssued                    bool          `json:"PaymentReminderIssued,omitempty"`
	UsesGreenTechnology                      bool          `json:"UsesGreenTechnology,omitempty"`
	IsSold                                   bool          `json:"IsSold,omitempty"`
	FactoringInvoiceStatus                   int           `json:"FactoringInvoiceStatus,omitempty"`
	PaymentDate                              interface{}   `json:"PaymentDate,omitempty"`
	PaymentStatus                            int           `json:"PaymentStatus,omitempty"`
	CreditedBy                               []interface{} `json:"CreditedBy,omitempty"`
}

func (c CustomerInvoice) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

func (c CustomerInvoice) IsEmpty() bool {
	return zero.IsZero(c)
}

type CustomerInvoiceRows []CustomerInvoiceRow

type CustomerInvoiceRow struct {
	ID                      string      `json:"Id,omitempty"`
	ArticleNumber           string      `json:"ArticleNumber,omitempty"`
	ArticleID               string      `json:"ArticleId,omitempty"`
	IsServiceArticle        bool        `json:"IsServiceArticle,omitempty"`
	AmountNoVat             float64     `json:"AmountNoVat,omitempty"`
	PercentVat              float64     `json:"PercentVat,omitempty"`
	LineNumber              int         `json:"LineNumber,omitempty"`
	IsTextRow               bool        `json:"IsTextRow"`
	Text                    string      `json:"Text,omitempty"`
	UnitPrice               float64     `json:"UnitPrice,omitempty"`
	UnitAbbreviation        string      `json:"UnitAbbreviation,omitempty"`
	UnitAbbreviationEnglish string      `json:"UnitAbbreviationEnglish,omitempty"`
	DiscountPercentage      float64     `json:"DiscountPercentage,omitempty"`
	Quantity                float64     `json:"Quantity,omitempty"`
	IsWorkCost              bool        `json:"IsWorkCost,omitempty"`
	IsVatFree               bool        `json:"IsVatFree,omitempty"`
	CostCenterItemID1       interface{} `json:"CostCenterItemId1,omitempty"`
	CostCenterItemID2       interface{} `json:"CostCenterItemId2,omitempty"`
	CostCenterItemID3       interface{} `json:"CostCenterItemId3,omitempty"`
	UnitID                  string      `json:"UnitId,omitempty"`
	ProjectID               interface{} `json:"ProjectId,omitempty"`
	WorkCostType            int         `json:"WorkCostType,omitempty"`
	WorkHours               float64     `json:"WorkHours,omitempty"`
	MaterialCosts           float64     `json:"MaterialCosts,omitempty"`
	GreenTechnologyType     int         `json:"GreenTechnologyType,omitempty"`
}

func (r CustomerInvoiceRow) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r CustomerInvoiceRow) IsEmpty() bool {
	return zero.IsZero(r)
}

type CustomerLedgerItems []CustomerLedgerItem

type CustomerLedgerItem struct {
	CurrencyCode                   string   `json:"CurrencyCode"`
	CurrencyRate                   float64  `json:"CurrencyRate"`
	CurrencyRateUnit               float64  `json:"CurrencyRateUnit"`
	CustomerID                     string   `json:"CustomerId"`
	DueDate                        Date     `json:"DueDate"`
	ID                             string   `json:"Id,omitempty"`
	InvoiceDate                    Date     `json:"InvoiceDate"`
	InvoiceNumber                  int      `json:"InvoiceNumber"`
	IsCreditInvoice                bool     `json:"IsCreditInvoice"`
	PaymentReferenceNumber         string   `json:"PaymentReferenceNumber"`
	RemainingAmountInvoiceCurrency float64  `json:"RemainingAmountInvoiceCurrency"`
	RoundingsAmountInvoiceCurrency float64  `json:"RoundingsAmountInvoiceCurrency"`
	TotalAmountInvoiceCurrency     float64  `json:"TotalAmountInvoiceCurrency"`
	VATAmountInvoiceCurrency       float64  `json:"VATAmountInvoiceCurrency"`
	VoucherID                      string   `json:"VoucherId,omitempty"`
	ModifiedUTC                    DateTime `json:"ModifiedUtc,omitempty"`
	Voucher                        Voucher  `json:"Voucher,omitempty"`
}

func (i CustomerLedgerItem) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i CustomerLedgerItem) IsEmpty() bool {
	return zero.IsZero(i)
}
