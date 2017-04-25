package multivers

type ProductInfo struct {
	VatCodeID          int      `json:"vatCodeId"`
	EanCode            string   `json:"eanCode"`
	ProductState       int      `json:"productState"`
	DateCreated        DateNLNL `json:"dateCreated"`
	QuantityScale      int      `json:"quantityScale"`
	ShortName          string   `json:"shortName"`
	PriceExclVat       float64  `json:"priceExclVat"`
	Description        string   `json:"description"`
	ProductGroupID     string   `json:"productGroupId"`
	ProjectSurcharge   float64  `json:"projectSurcharge"`
	PriceInclVat       float64  `json:"priceInclVat"`
	ProductID          string   `json:"productId"`
	TechnicalStock     float64  `json:"technicalStock"`
	LastUpdate         DateNLNL `json:"lastUpdate"`
	StockTransferPrice float64  `json:"stockTransferPrice"`
	DiscountAccountID  string   `json:"discountAccountId"`
	Unit               string   `json:"unit"`
	ProjectEntryType   string   `json:"projectEntryType"`
	ProductType        int      `json:"productType"`
	AccountID          string   `json:"accountId"`
	PricePer           float64  `json:"pricePer"`
}
