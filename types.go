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
