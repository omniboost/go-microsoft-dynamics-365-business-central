package central

import "time"

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
