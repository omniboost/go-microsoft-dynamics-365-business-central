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
