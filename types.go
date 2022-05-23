package vismaonline

import "time"

type Meta struct {
	CurrentPage          int       `json:"CurrentPage"`
	PageSize             int       `json:"PageSize"`
	TotalNumberOfPages   int       `json:"TotalNumberOfPages"`
	TotalNumberOfResults int       `json:"TotalNumberOfResults"`
	ServerTimeUtc        time.Time `json:"ServerTimeUtc"`
}

type Accounts []Account

type Account struct {
	Name                      string    `json:"Name"`
	Number                    string    `json:"Number"`
	VatCodeID                 int       `json:"VatCodeId"`
	VatCodeDescription        string    `json:"VatCodeDescription"`
	FiscalYearID              string    `json:"FiscalYearId"`
	ReferenceCode             string    `json:"ReferenceCode"`
	Type                      int       `json:"Type"`
	TypeDescription           string    `json:"TypeDescription"`
	ModifiedUtc               time.Time `json:"ModifiedUtc"`
	IsActive                  bool      `json:"IsActive"`
	IsProjectAllowed          bool      `json:"IsProjectAllowed"`
	IsCostCenterAllowed       bool      `json:"IsCostCenterAllowed"`
	IsBlockedForManualBooking bool      `json:"IsBlockedForManualBooking"`
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
	VatCode                 string      `json:"VatCode"`
	BankGiro                string      `json:"BankGiro"`
	PlusGiro                string      `json:"PlusGiro"`
	BankAccount             interface{} `json:"BankAccount"`
	Iban                    string      `json:"Iban"`
	AccountNumberDigits     int         `json:"AccountNumberDigits"`
	AccountingLockedTo      time.Time   `json:"AccountingLockedTo"`
	AccountingLockInterval  int         `json:"AccountingLockInterval"`
	TaxDeclarationDate      interface{} `json:"TaxDeclarationDate"`
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
	KeepOriginalDraftDate              interface{} `json:"KeepOriginalDraftDate"`
	UsePaymentFilesForOutgoingPayments bool        `json:"UsePaymentFilesForOutgoingPayments"`
	UseAutomaticVatCalculation         bool        `json:"UseAutomaticVatCalculation"`
	ShowCostCenterReminder             bool        `json:"ShowCostCenterReminder"`
	ShowProjectReminder                bool        `json:"ShowProjectReminder"`
}

type Vatcodes []Vatcode

type Vatcode struct {
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
