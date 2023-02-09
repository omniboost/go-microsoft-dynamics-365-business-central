package central

import (
	"time"

	"github.com/cydev/zero"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/omitempty"
)

type Accounts []Account

type Account struct {
	OdataEtag            string    `json:"@odata.etag"`
	ID                   string    `json:"id"`
	Number               string    `json:"number"`
	DisplayName          string    `json:"displayName"`
	Category             string    `json:"category"`
	SubCategory          string    `json:"subCategory"`
	Blocked              bool      `json:"blocked"`
	AccountType          string    `json:"accountType"`
	DirectPosting        bool      `json:"directPosting"`
	NetChange            float64   `json:"netChange"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}

type Journals []Journal

type Journal struct {
	OdataEtag              string    `json:"@odata.etag"`
	ID                     string    `json:"id"`
	Code                   string    `json:"code"`
	DisplayName            string    `json:"displayName"`
	TemplateDisplayName    string    `json:"templateDisplayName"`
	LastModifiedDateTime   time.Time `json:"lastModifiedDateTime"`
	BalancingAccountID     string    `json:"balancingAccountId"`
	BalancingAccountNumber string    `json:"balancingAccountNumber"`
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
	OdataEtag            string    `json:"@odata.etag"`
	ID                   string    `json:"id"`
	Code                 string    `json:"code"`
	DisplayName          string    `json:"displayName"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}

func (d Dimension) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

func (d Dimension) IsEmpty() bool {
	return zero.IsZero(d)
}

type Companies []Company

type Company struct {
	ID                string    `json:"id"`
	SystemVersion     string    `json:"systemVersion"`
	Timestamp         int       `json:"timestamp"`
	Name              string    `json:"name"`
	DisplayName       string    `json:"displayName"`
	BusinessProfileID string    `json:"businessProfileId"`
	SystemCreatedAt   time.Time `json:"systemCreatedAt"`
	SystemCreatedBy   string    `json:"systemCreatedBy"`
	SystemModifiedAt  time.Time `json:"systemModifiedAt"`
	SystemModifiedBy  string    `json:"systemModifiedBy"`
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
	OdataEtag            string    `json:"@odata.etag"`
	ID                   string    `json:"id"`
	Code                 string    `json:"code"`
	DimensionID          string    `json:"dimensionId"`
	DisplayName          string    `json:"displayName"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}
