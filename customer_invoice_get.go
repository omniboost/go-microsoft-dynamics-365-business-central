package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerInvoiceGet() CustomerInvoiceGet {
	r := CustomerInvoiceGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerInvoiceGet struct {
	client      *Client
	queryParams *CustomerInvoiceGetQueryParams
	pathParams  *CustomerInvoiceGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoiceGetBody
}

func (r CustomerInvoiceGet) NewQueryParams() *CustomerInvoiceGetQueryParams {
	return &CustomerInvoiceGetQueryParams{}
}

type CustomerInvoiceGetQueryParams struct{}

func (p CustomerInvoiceGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInvoiceGet) QueryParams() *CustomerInvoiceGetQueryParams {
	return r.queryParams
}

func (r CustomerInvoiceGet) NewPathParams() *CustomerInvoiceGetPathParams {
	return &CustomerInvoiceGetPathParams{}
}

type CustomerInvoiceGetPathParams struct {
	InvoiceNumber string `schema:"invoice_number"`
}

func (p *CustomerInvoiceGetPathParams) Params() map[string]string {
	return map[string]string{
		"invoice_number": p.InvoiceNumber,
	}
}

func (r *CustomerInvoiceGet) PathParams() *CustomerInvoiceGetPathParams {
	return r.pathParams
}

func (r *CustomerInvoiceGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerInvoiceGet) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoiceGet) Method() string {
	return r.method
}

func (r CustomerInvoiceGet) NewRequestBody() CustomerInvoiceGetBody {
	return CustomerInvoiceGetBody{}
}

type CustomerInvoiceGetBody struct {
}

func (r *CustomerInvoiceGet) RequestBody() *CustomerInvoiceGetBody {
	return nil
}

func (r *CustomerInvoiceGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerInvoiceGet) SetRequestBody(body CustomerInvoiceGetBody) {
	r.requestBody = body
}

func (r *CustomerInvoiceGet) NewResponseBody() *CustomerInvoiceGetResponseBody {
	return &CustomerInvoiceGetResponseBody{}
}

type CustomerInvoiceGetResponseBody struct {
	CreditTerms struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"creditTerms"`
	DocumentDueDate   string        `json:"documentDueDate"`
	CashDiscountDate  string        `json:"cashDiscountDate"`
	ExternalReference string        `json:"externalReference"`
	ExchangeRate      float64       `json:"exchangeRate"`
	Attachments       []interface{} `json:"attachments"`
	TaxDetails        []struct {
		TaxID         string  `json:"taxId"`
		VatRate       float64 `json:"vatRate"`
		TaxableAmount float64 `json:"taxableAmount"`
		VatAmount     float64 `json:"vatAmount"`
	} `json:"taxDetails"`
	InvoiceLines []struct {
		LineNumber             int     `json:"lineNumber"`
		Description            string  `json:"description"`
		Quantity               float64 `json:"quantity"`
		UnitPrice              float64 `json:"unitPrice"`
		UnitPriceInCurrency    float64 `json:"unitPriceInCurrency"`
		ManualAmount           float64 `json:"manualAmount"`
		ManualAmountInCurrency float64 `json:"manualAmountInCurrency"`
		Amount                 float64 `json:"amount"`
		AmountInCurrency       float64 `json:"amountInCurrency"`
		Account                struct {
			Type        string `json:"type"`
			Number      string `json:"number"`
			Description string `json:"description"`
		} `json:"account"`
		VatCode struct {
			ID          string `json:"id"`
			Description string `json:"description"`
		} `json:"vatCode"`
		DiscountPercent          float64 `json:"discountPercent"`
		DiscountAmount           float64 `json:"discountAmount"`
		DiscountAmountInCurrency float64 `json:"discountAmountInCurrency"`
		ManualDiscount           bool    `json:"manualDiscount"`
		Subaccount               struct {
			SubaccountNumber     string `json:"subaccountNumber"`
			SubaccountID         int    `json:"subaccountId"`
			Description          string `json:"description"`
			LastModifiedDateTime string `json:"lastModifiedDateTime"`
			Active               bool   `json:"active"`
			Segments             []struct {
				SegmentID               int    `json:"segmentId"`
				SegmentDescription      string `json:"segmentDescription"`
				SegmentValue            string `json:"segmentValue"`
				SegmentValueDescription string `json:"segmentValueDescription"`
			} `json:"segments"`
		} `json:"subaccount"`
		DeferralSchedule int `json:"deferralSchedule"`
		BranchNumber     struct {
			Number string `json:"number"`
			Name   string `json:"name"`
		} `json:"branchNumber"`
	} `json:"invoiceLines"`
	SendToAutoInvoice bool    `json:"sendToAutoInvoice"`
	RoundingDiff      float64 `json:"roundingDiff"`
	CustomerVatZone   struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"customerVatZone"`
	AccountingCostRef         string  `json:"accountingCostRef"`
	OriginatorDocRef          string  `json:"originatorDocRef"`
	ContractDocRef            string  `json:"contractDocRef"`
	Hold                      bool    `json:"hold"`
	DiscountTotal             float64 `json:"discountTotal"`
	DiscountTotalInCurrency   float64 `json:"discountTotalInCurrency"`
	DetailTotal               float64 `json:"detailTotal"`
	DetailTotalInCurrency     float64 `json:"detailTotalInCurrency"`
	VatTaxableTotal           float64 `json:"vatTaxableTotal"`
	VatTaxableTotalInCurrency float64 `json:"vatTaxableTotalInCurrency"`
	VatExemptTotal            float64 `json:"vatExemptTotal"`
	VatExemptTotalInCurrency  float64 `json:"vatExemptTotalInCurrency"`
	PaymentReference          string  `json:"paymentReference"`
	InvoiceAddress            struct {
		AddressID    int    `json:"addressId"`
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		AddressLine3 string `json:"addressLine3"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		Country      struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		OverrideAddress bool `json:"overrideAddress"`
	} `json:"invoiceAddress"`
	InvoiceContact struct {
		ContactID       int    `json:"contactId"`
		BusinessName    string `json:"businessName"`
		Attention       string `json:"attention"`
		Email           string `json:"email"`
		Phone1          string `json:"phone1"`
		OverrideContact bool   `json:"overrideContact"`
	} `json:"invoiceContact"`
	Applications []interface{} `json:"applications"`
	DontPrint    bool          `json:"dontPrint"`
	DontEmail    bool          `json:"dontEmail"`
	Customer     struct {
		Number string `json:"number"`
		Name   string `json:"name"`
	} `json:"customer"`
	DocumentType           string  `json:"documentType"`
	ReferenceNumber        string  `json:"referenceNumber"`
	PostPeriod             string  `json:"postPeriod"`
	FinancialPeriod        string  `json:"financialPeriod"`
	DocumentDate           string  `json:"documentDate"`
	OrigInvoiceDate        string  `json:"origInvoiceDate"`
	Status                 string  `json:"status"`
	CurrencyID             string  `json:"currencyId"`
	Amount                 float64 `json:"amount"`
	AmountInCurrency       float64 `json:"amountInCurrency"`
	Balance                float64 `json:"balance"`
	BalanceInCurrency      float64 `json:"balanceInCurrency"`
	CashDiscount           float64 `json:"cashDiscount"`
	CashDiscountInCurrency float64 `json:"cashDiscountInCurrency"`
	PaymentMethod          struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"paymentMethod"`
	CustomerRefNumber    string  `json:"customerRefNumber"`
	InvoiceText          string  `json:"invoiceText"`
	LastModifiedDateTime string  `json:"lastModifiedDateTime"`
	CreatedDateTime      string  `json:"createdDateTime"`
	Note                 string  `json:"note"`
	VatTotal             float64 `json:"vatTotal"`
	VatTotalInCurrency   float64 `json:"vatTotalInCurrency"`
	Location             struct {
		CountryID string `json:"countryId"`
		ID        string `json:"id"`
		Name      string `json:"name"`
	} `json:"location"`
	BranchNumber struct {
		Number string `json:"number"`
		Name   string `json:"name"`
	} `json:"branchNumber"`
	CashAccount string `json:"cashAccount"`
	Project     struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"project"`
	Account struct {
		Type        string `json:"type"`
		Number      string `json:"number"`
		Description string `json:"description"`
	} `json:"account"`
	Subaccount struct {
		SubaccountNumber     string `json:"subaccountNumber"`
		SubaccountID         int    `json:"subaccountId"`
		Description          string `json:"description"`
		LastModifiedDateTime string `json:"lastModifiedDateTime"`
		Active               bool   `json:"active"`
		Segments             []struct {
			SegmentID               int    `json:"segmentId"`
			SegmentDescription      string `json:"segmentDescription"`
			SegmentValue            string `json:"segmentValue"`
			SegmentValueDescription string `json:"segmentValueDescription"`
		} `json:"segments"`
	} `json:"subaccount"`
	CustomerProject string `json:"customerProject"`
	Metadata        struct {
		TotalCount  int `json:"totalCount"`
		MaxPageSize int `json:"maxPageSize"`
	} `json:"metadata"`
}

func (r *CustomerInvoiceGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customerinvoice/{{.invoice_number}}", r.PathParams())
	return &u
}

func (r *CustomerInvoiceGet) Do() (CustomerInvoiceGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
