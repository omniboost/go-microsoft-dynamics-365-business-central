package central

import (
	"github.com/cydev/zero"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/omitempty"
)

type Accounts []Account

type Account struct {
	OdataEtag            string  `json:"@odata.etag"`
	ID                   string  `json:"id"`
	Number               string  `json:"number"`
	DisplayName          string  `json:"displayName"`
	Category             string  `json:"category"`
	SubCategory          string  `json:"subCategory"`
	Blocked              bool    `json:"blocked"`
	AccountType          string  `json:"accountType"`
	DirectPosting        bool    `json:"directPosting"`
	NetChange            float64 `json:"netChange"`
	LastModifiedDateTime Time    `json:"lastModifiedDateTime"`
}

type Journals []Journal

type Journal struct {
	OdataEtag              string `json:"@odata.etag"`
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	DisplayName            string `json:"displayName"`
	TemplateDisplayName    string `json:"templateDisplayName"`
	LastModifiedDateTime   Time   `json:"lastModifiedDateTime"`
	BalancingAccountID     string `json:"balancingAccountId"`
	BalancingAccountNumber string `json:"balancingAccountNumber"`
}

type JournalLines []JournalLine

type JournalLine struct {
	OdataEtag              string  `json:"@odata.etag,omitempty"`
	ID                     string  `json:"id,omitempty"`
	JournalID              string  `json:"journalId,omitempty"`
	JournalDisplayName     string  `json:"journalDisplayName,omitempty"`
	LineNumber             int     `json:"lineNumber"`
	AccountType            string  `json:"accountType,omitempty"`
	AccountID              string  `json:"accountId,omitempty"`
	AccountNumber          string  `json:"accountNumber,omitempty"`
	PostingDate            string  `json:"postingDate"`
	DocumentNumber         string  `json:"documentNumber"`
	ExternalDocumentNumber string  `json:"externalDocumentNumber"`
	Amount                 float64 `json:"amount"`
	Description            string  `json:"description"`
	Comment                string  `json:"comment"`
	TaxCode                string  `json:"taxCode"`
	BalanceAccountType     string  `json:"balanceAccountType,omitempty"`
	BalancingAccountID     string  `json:"balancingAccountId,omitempty"`
	BalancingAccountNumber string  `json:"balancingAccountNumber,omitempty"`
	LastModifiedDateTime   Time    `json:"lastModifiedDateTime,omitempty"`
}

func (l JournalLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}

func (l JournalLine) IsEmpty() bool {
	return zero.IsZero(l)
}

type Dimensions []Dimension

type Dimension struct {
	OdataEtag            string `json:"@odata.etag"`
	ID                   string `json:"id"`
	Code                 string `json:"code"`
	DisplayName          string `json:"displayName"`
	LastModifiedDateTime Time   `json:"lastModifiedDateTime"`
}

func (d Dimension) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

func (d Dimension) IsEmpty() bool {
	return zero.IsZero(d)
}

type Companies []Company

type Company struct {
	ID                string `json:"id"`
	SystemVersion     string `json:"systemVersion"`
	Timestamp         int    `json:"timestamp"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
	BusinessProfileID string `json:"businessProfileId"`
	SystemCreatedAt   Time   `json:"systemCreatedAt"`
	SystemCreatedBy   string `json:"systemCreatedBy"`
	SystemModifiedAt  Time   `json:"systemModifiedAt"`
	SystemModifiedBy  string `json:"systemModifiedBy"`
}

func (c Company) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

func (c Company) IsEmpty() bool {
	return zero.IsZero(c)
}

type DimensionSetLines []DimensionSetLine

type DimensionSetLine struct {
	OdataEtag        string `json:"@odata.etag,omitempty"`
	ID               string `json:"id,omitempty"`
	Code             string `json:"code"`
	ParentID         string `json:"parentId"`
	ParentType       string `json:"parentType,omitempty"`
	DisplayName      string `json:"displayName,omitempty"`
	ValueID          string `json:"valueId,omitempty"`
	ValueCode        string `json:"valueCode,omitempty"`
	ValueDisplayName string `json:"valueDisplayName,omitempty"`
}

type DimensionValues []DimensionValue

type DimensionValue struct {
	OdataEtag            string `json:"@odata.etag"`
	ID                   string `json:"id"`
	Code                 string `json:"code"`
	DimensionID          string `json:"dimensionId"`
	DisplayName          string `json:"displayName"`
	LastModifiedDateTime Time   `json:"lastModifiedDateTime"`
}

type SalesInvoices []SalesInvoice

type SalesInvoice struct {
	OdataEtag                      string  `json:"@odata.etag,omitempty"`
	ID                             string  `json:"id,omitempty"`
	Number                         string  `json:"number,omitempty"`
	ExternalDocumentNumber         string  `json:"externalDocumentNumber,omitempty"`
	InvoiceDate                    string  `json:"invoiceDate,omitempty"`
	PostingDate                    string  `json:"postingDate,omitempty"`
	DueDate                        string  `json:"dueDate,omitempty"`
	CustomerPurchaseOrderReference string  `json:"customerPurchaseOrderReference,omitempty"`
	CustomerID                     string  `json:"customerId,omitempty"`
	CustomerNumber                 string  `json:"customerNumber,omitempty"`
	CustomerName                   string  `json:"customerName,omitempty"`
	BillToName                     string  `json:"billToName,omitempty"`
	BillToCustomerID               string  `json:"billToCustomerId,omitempty"`
	BillToCustomerNumber           string  `json:"billToCustomerNumber,omitempty"`
	ShipToName                     string  `json:"shipToName,omitempty"`
	ShipToContact                  string  `json:"shipToContact,omitempty"`
	SellToAddressLine1             string  `json:"sellToAddressLine1,omitempty"`
	SellToAddressLine2             string  `json:"sellToAddressLine2,omitempty"`
	SellToCity                     string  `json:"sellToCity,omitempty"`
	SellToCountry                  string  `json:"sellToCountry,omitempty"`
	SellToState                    string  `json:"sellToState,omitempty"`
	SellToPostCode                 string  `json:"sellToPostCode,omitempty"`
	BillToAddressLine1             string  `json:"billToAddressLine1,omitempty"`
	BillToAddressLine2             string  `json:"billToAddressLine2,omitempty"`
	BillToCity                     string  `json:"billToCity,omitempty"`
	BillToCountry                  string  `json:"billToCountry,omitempty"`
	BillToState                    string  `json:"billToState,omitempty"`
	BillToPostCode                 string  `json:"billToPostCode,omitempty"`
	ShipToAddressLine1             string  `json:"shipToAddressLine1,omitempty"`
	ShipToAddressLine2             string  `json:"shipToAddressLine2,omitempty"`
	ShipToCity                     string  `json:"shipToCity,omitempty"`
	ShipToCountry                  string  `json:"shipToCountry,omitempty"`
	ShipToState                    string  `json:"shipToState,omitempty"`
	ShipToPostCode                 string  `json:"shipToPostCode,omitempty"`
	CurrencyID                     string  `json:"currencyId,omitempty"`
	ShortcutDimension1Code         string  `json:"shortcutDimension1Code,omitempty"`
	ShortcutDimension2Code         string  `json:"shortcutDimension2Code,omitempty"`
	CurrencyCode                   string  `json:"currencyCode,omitempty"`
	OrderID                        string  `json:"orderId,omitempty"`
	OrderNumber                    string  `json:"orderNumber,omitempty"`
	PaymentTermsID                 string  `json:"paymentTermsId,omitempty"`
	ShipmentMethodID               string  `json:"shipmentMethodId,omitempty"`
	Salesperson                    string  `json:"salesperson,omitempty"`
	PricesIncludeTax               bool    `json:"pricesIncludeTax,omitempty"`
	RemainingAmount                float64 `json:"remainingAmount,omitempty"`
	DiscountAmount                 float64 `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax       bool    `json:"discountAppliedBeforeTax,omitempty"`
	TotalAmountExcludingTax        float64 `json:"totalAmountExcludingTax,omitempty"`
	TotalTaxAmount                 float64 `json:"totalTaxAmount,omitempty"`
	TotalAmountIncludingTax        float64 `json:"totalAmountIncludingTax,omitempty"`
	Status                         string  `json:"status,omitempty"`
	LastModifiedDateTime           Time    `json:"lastModifiedDateTime,omitempty"`
	PhoneNumber                    string  `json:"phoneNumber,omitempty"`
	Email                          string  `json:"email,omitempty"`
}

func (i SalesInvoice) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type SalesInvoiceLines []SalesInvoiceLine

type SalesInvoiceLine struct {
	OdataEtag                 string  `json:"@odata.etag,omitempty"`
	ID                        string  `json:"id,omitempty"`
	DocumentID                string  `json:"documentId,omitempty"`
	Sequence                  int     `json:"sequence,omitempty"`
	ItemID                    string  `json:"itemId,omitempty"`
	AccountID                 string  `json:"accountId,omitempty"`
	LineType                  string  `json:"lineType,omitempty"`
	LineObjectNumber          string  `json:"lineObjectNumber,omitempty"`
	Description               string  `json:"description,omitempty"`
	Description2              string  `json:"description2,omitempty"`
	UnitOfMeasureID           string  `json:"unitOfMeasureId,omitempty"`
	UnitOfMeasureCode         string  `json:"unitOfMeasureCode,omitempty"`
	Quantity                  float64 `json:"quantity,omitempty"`
	UnitPrice                 float64 `json:"unitPrice,omitempty"`
	DiscountAmount            float64 `json:"discountAmount,omitempty"`
	DiscountPercent           float64 `json:"discountPercent,omitempty"`
	DiscountAppliedBeforeTax  bool    `json:"discountAppliedBeforeTax,omitempty"`
	AmountExcludingTax        float64 `json:"amountExcludingTax,omitempty"`
	TaxCode                   string  `json:"taxCode,omitempty"`
	TaxPercent                float64 `json:"taxPercent,omitempty"`
	TotalTaxAmount            float64 `json:"totalTaxAmount,omitempty"`
	AmountIncludingTax        float64 `json:"amountIncludingTax,omitempty"`
	InvoiceDiscountAllocation int     `json:"invoiceDiscountAllocation,omitempty"`
	NetAmount                 float64 `json:"netAmount,omitempty"`
	NetTaxAmount              float64 `json:"netTaxAmount,omitempty"`
	NetAmountIncludingTax     float64 `json:"netAmountIncludingTax,omitempty"`
	ShipmentDate              string  `json:"shipmentDate,omitempty"`
	ItemVariantID             string  `json:"itemVariantId,omitempty"`
	LocationID                string  `json:"locationId,omitempty"`
}

func (i SalesInvoiceLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type Customers []Customer

type Customer struct {
	OdataEtag             string  `json:"@odata.etag,omitempty"`
	ID                    string  `json:"id,omitempty"`
	Number                string  `json:"number,omitempty"`
	DisplayName           string  `json:"displayName,omitempty"`
	Type                  string  `json:"type,omitempty"`
	AddressLine1          string  `json:"addressLine1,omitempty"`
	AddressLine2          string  `json:"addressLine2,omitempty"`
	City                  string  `json:"city,omitempty"`
	State                 string  `json:"state,omitempty"`
	Country               string  `json:"country,omitempty"`
	PostalCode            string  `json:"postalCode,omitempty"`
	PhoneNumber           string  `json:"phoneNumber,omitempty"`
	Email                 string  `json:"email,omitempty"`
	Website               string  `json:"website,omitempty"`
	SalespersonCode       string  `json:"salespersonCode,omitempty"`
	BalanceDue            float64 `json:"balanceDue,omitempty"`
	CreditLimit           float64 `json:"creditLimit,omitempty"`
	TaxLiable             bool    `json:"taxLiable,omitempty"`
	TaxAreaID             string  `json:"taxAreaId,omitempty"`
	TaxAreaDisplayName    string  `json:"taxAreaDisplayName,omitempty"`
	TaxRegistrationNumber string  `json:"taxRegistrationNumber,omitempty"`
	CurrencyID            string  `json:"currencyId,omitempty"`
	CurrencyCode          string  `json:"currencyCode,omitempty"`
	PaymentTermsID        string  `json:"paymentTermsId,omitempty"`
	ShipmentMethodID      string  `json:"shipmentMethodId,omitempty"`
	PaymentMethodID       string  `json:"paymentMethodId,omitempty"`
	Blocked               string  `json:"blocked,omitempty"`
	LastModifiedDateTime  Time    `json:"lastModifiedDateTime,omitempty"`
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

type SalesCreditMemos []SalesCreditMemo

type SalesCreditMemo struct {
	OdataEtag                string  `json:"@odata.etag,omitempty"`
	ID                       string  `json:"id,omitempty"`
	Number                   string  `json:"number,omitempty"`
	ExternalDocumentNumber   string  `json:"externalDocumentNumber,omitempty"`
	CreditMemoDate           string  `json:"creditMemoDate,omitempty"`
	PostingDate              string  `json:"postingDate,omitempty"`
	DueDate                  string  `json:"dueDate,omitempty"`
	CustomerID               string  `json:"customerId,omitempty"`
	CustomerNumber           string  `json:"customerNumber,omitempty"`
	CustomerName             string  `json:"customerName,omitempty"`
	BillToName               string  `json:"billToName,omitempty"`
	BillToCustomerID         string  `json:"billToCustomerId,omitempty"`
	BillToCustomerNumber     string  `json:"billToCustomerNumber,omitempty"`
	SellToAddressLine1       string  `json:"sellToAddressLine1,omitempty"`
	SellToAddressLine2       string  `json:"sellToAddressLine2,omitempty"`
	SellToCity               string  `json:"sellToCity,omitempty"`
	SellToCountry            string  `json:"sellToCountry,omitempty"`
	SellToState              string  `json:"sellToState,omitempty"`
	SellToPostCode           string  `json:"sellToPostCode,omitempty"`
	BillToAddressLine1       string  `json:"billToAddressLine1,omitempty"`
	BillToAddressLine2       string  `json:"billToAddressLine2,omitempty"`
	BillToCity               string  `json:"billToCity,omitempty"`
	BillToCountry            string  `json:"billToCountry,omitempty"`
	BillToState              string  `json:"billToState,omitempty"`
	BillToPostCode           string  `json:"billToPostCode,omitempty"`
	ShortcutDimension1Code   string  `json:"shortcutDimension1Code,omitempty"`
	ShortcutDimension2Code   string  `json:"shortcutDimension2Code,omitempty"`
	CurrencyID               string  `json:"currencyId,omitempty"`
	CurrencyCode             string  `json:"currencyCode,omitempty"`
	PaymentTermsID           string  `json:"paymentTermsId,omitempty"`
	ShipmentMethodID         string  `json:"shipmentMethodId,omitempty"`
	Salesperson              string  `json:"salesperson,omitempty"`
	PricesIncludeTax         bool    `json:"pricesIncludeTax,omitempty"`
	DiscountAmount           float64 `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax bool    `json:"discountAppliedBeforeTax,omitempty"`
	TotalAmountExcludingTax  float64 `json:"totalAmountExcludingTax,omitempty"`
	TotalTaxAmount           float64 `json:"totalTaxAmount,omitempty"`
	TotalAmountIncludingTax  float64 `json:"totalAmountIncludingTax,omitempty"`
	Status                   string  `json:"status,omitempty"`
	LastModifiedDateTime     Time    `json:"lastModifiedDateTime,omitempty"`
	InvoiceID                string  `json:"invoiceId,omitempty"`
	InvoiceNumber            string  `json:"invoiceNumber,omitempty"`
	PhoneNumber              string  `json:"phoneNumber,omitempty"`
	Email                    string  `json:"email,omitempty"`
	CustomerReturnReasonID   string  `json:"customerReturnReasonId,omitempty"`
}

func (i SalesCreditMemo) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type SalesCreditMemoLines []SalesCreditMemoLine

type SalesCreditMemoLine struct {
	OdataEtag                 string  `json:"@odata.etag,omitempty"`
	ID                        string  `json:"id,omitempty"`
	DocumentID                string  `json:"documentId,omitempty"`
	Sequence                  int     `json:"sequence,omitempty"`
	ItemID                    string  `json:"itemId,omitempty"`
	AccountID                 string  `json:"accountId,omitempty"`
	LineType                  string  `json:"lineType,omitempty"`
	LineObjectNumber          string  `json:"lineObjectNumber,omitempty"`
	Description               string  `json:"description,omitempty"`
	Description2              string  `json:"description2,omitempty"`
	UnitOfMeasureID           string  `json:"unitOfMeasureId,omitempty"`
	UnitOfMeasureCode         string  `json:"unitOfMeasureCode,omitempty"`
	UnitPrice                 float64 `json:"unitPrice,omitempty"`
	Quantity                  float64 `json:"quantity,omitempty"`
	DiscountAmount            float64 `json:"discountAmount,omitempty"`
	DiscountPercent           float64 `json:"discountPercent,omitempty"`
	DiscountAppliedBeforeTax  bool    `json:"discountAppliedBeforeTax,omitempty"`
	AmountExcludingTax        float64 `json:"amountExcludingTax,omitempty"`
	TaxCode                   string  `json:"taxCode,omitempty"`
	TaxPercent                float64 `json:"taxPercent,omitempty"`
	TotalTaxAmount            float64 `json:"totalTaxAmount,omitempty"`
	AmountIncludingTax        float64 `json:"amountIncludingTax,omitempty"`
	InvoiceDiscountAllocation float64 `json:"invoiceDiscountAllocation,omitempty"`
	NetAmount                 float64 `json:"netAmount,omitempty"`
	NetTaxAmount              float64 `json:"netTaxAmount,omitempty"`
	NetAmountIncludingTax     float64 `json:"netAmountIncludingTax,omitempty"`
	ShipmentDate              string  `json:"shipmentDate,omitempty"`
	ItemVariantID             string  `json:"itemVariantId,omitempty"`
	LocationID                string  `json:"locationId,omitempty"`
}

func (i SalesCreditMemoLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}
