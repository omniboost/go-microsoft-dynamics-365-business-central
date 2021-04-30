package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/omitempty"
	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerCreditNoteV2Post() CustomerCreditNoteV2Post {
	r := CustomerCreditNoteV2Post{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerCreditNoteV2Post struct {
	client      *Client
	queryParams *CustomerCreditNoteV2PostQueryParams
	pathParams  *CustomerCreditNoteV2PostPathParams
	method      string
	headers     http.Header
	requestBody CustomerCreditNoteV2PostBody
}

func (r CustomerCreditNoteV2Post) NewQueryParams() *CustomerCreditNoteV2PostQueryParams {
	return &CustomerCreditNoteV2PostQueryParams{}
}

type CustomerCreditNoteV2PostQueryParams struct{}

func (p CustomerCreditNoteV2PostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerCreditNoteV2Post) QueryParams() *CustomerCreditNoteV2PostQueryParams {
	return r.queryParams
}

func (r CustomerCreditNoteV2Post) NewPathParams() *CustomerCreditNoteV2PostPathParams {
	return &CustomerCreditNoteV2PostPathParams{}
}

type CustomerCreditNoteV2PostPathParams struct {
}

func (p *CustomerCreditNoteV2PostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerCreditNoteV2Post) PathParams() *CustomerCreditNoteV2PostPathParams {
	return r.pathParams
}

func (r *CustomerCreditNoteV2Post) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerCreditNoteV2Post) SetMethod(method string) {
	r.method = method
}

func (r *CustomerCreditNoteV2Post) Method() string {
	return r.method
}

func (r CustomerCreditNoteV2Post) NewRequestBody() CustomerCreditNoteV2PostBody {
	return CustomerCreditNoteV2PostBody{}
}

type CustomerCreditNoteV2PostBody struct {
	CurrencyID        ValueString `json:"currencyId"`
	CustomerRefNumber ValueString `json:"customerRefNumber"`
	ExternalReference ValueString `json:"externalReference"`
	Contact           ValueInt    `json:"contact,omitempty"`
	Project           ValueString `json:"project,omitempty"`
	Lines             []struct {
		DiscountCode             ValueString `json:"discountCode,omitempty"`
		TaskID                   ValueString `json:"taskId,omitempty"`
		Operation                string      `json:"operation"`
		InventoryNumber          ValueString `json:"inventoryNumber,omitempty"`
		LineNumber               ValueInt    `json:"lineNumber"`
		Description              ValueString `json:"description"`
		Quantity                 ValueInt    `json:"quantity"`
		UnitPriceInCurrency      ValueNumber `json:"unitPriceInCurrency"`
		ManualAmountInCurrency   ValueNumber `json:"manualAmountInCurrency"`
		AccountNumber            ValueString `json:"accountNumber"`
		VATCodeID                ValueString `json:"vatCodeId"`
		Uom                      ValueString `json:"uom,omitempty"`
		DiscountPercent          ValueNumber `json:"discountPercent"`
		DiscountAmountInCurrency ValueNumber `json:"discountAmountInCurrency"`
		ManualDiscount           ValueBool   `json:"manualDiscount"`
		Subaccount               []struct {
			SegmentID    int    `json:"segmentId"`
			SegmentValue string `json:"segmentValue"`
		} `json:"subaccount"`
		Salesperson      ValueString `json:"salesperson,omitempty"`
		DeferralSchedule ValueInt    `json:"deferralSchedule,omitempty"`
		DeferralCode     ValueString `json:"deferralCode,omitempty"`
		TermStartDate    ValueTime   `json:"termStartDate,omitempty"`
		TermEndDate      ValueTime   `json:"termEndDate,omitempty"`
		Note             ValueString `json:"note,omitempty"`
		BranchNumber     ValueString `json:"branchNumber,omitempty"`
	} `json:"lines"`
	TaxDetails []struct {
		TaxID         ValueString `json:"taxId"`
		TaxableAmount ValueNumber `json:"taxableAmount"`
		VATAmount     ValueNumber `json:"vatAmount"`
	} `json:"taxDetails"`
	ApplicationLines []struct {
		Operation    string      `json:"operation"`
		DocumentType ValueString `json:"documentType"`
		RefNbr       ValueString `json:"refNbr"`
		AmountPaid   ValueNumber `json:"amountPaid"`
	} `json:"applicationLines"`
	CustomerVATZoneID    ValueString  `json:"customerVatZoneId,omitempty"`
	InvoiceAddress       ValueAddress `json:"invoiceAddress"`
	InvoiceContact       Contact      `json:"invoiceContact,omitempty"`
	OverrideNumberSeries ValueBool    `json:"overrideNumberSeries"`
	SendToAutoInvoice    ValueBool    `json:"sendToAutoInvoice"`
	ReferenceNumber      ValueString  `json:"referenceNumber"`
	CustomerNumber       ValueString  `json:"customerNumber"`
	DocumentDate         ValueTime    `json:"documentDate"`
	OrigInvoiceDate      ValueTime    `json:"origInvoiceDate,omitempty"`
	Hold                 ValueBool    `json:"hold"`
	PostPeriod           ValueString  `json:"postPeriod"`
	FinancialPeriod      ValueString  `json:"financialPeriod"`
	InvoiceText          ValueString  `json:"invoiceText"`
	LocationID           ValueString  `json:"locationId,omitempty"`
	SalesPersonID        ValueInt     `json:"salesPersonID,omitempty"`
	Salesperson          ValueString  `json:"salesperson,omitempty"`
	Note                 ValueString  `json:"note"`
	BranchNumber         ValueString  `json:"branchNumber,omitempty"`
	CashAccount          ValueString  `json:"cashAccount,omitempty"`
	DontPrint            ValueBool    `json:"dontPrint"`
	DontEmail            ValueBool    `json:"dontEmail"`
}

func (r *CustomerCreditNoteV2Post) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *CustomerCreditNoteV2Post) RequestBody() *CustomerCreditNoteV2PostBody {
	return &r.requestBody
}

func (r *CustomerCreditNoteV2Post) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerCreditNoteV2Post) SetRequestBody(body CustomerCreditNoteV2PostBody) {
	r.requestBody = body
}

func (r *CustomerCreditNoteV2Post) NewResponseBody() *CustomerCreditNoteV2PostResponseBody {
	return &CustomerCreditNoteV2PostResponseBody{}
}

type CustomerCreditNoteV2PostResponseBody JournalTransactions

func (r *CustomerCreditNoteV2Post) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v2/customerCreditNote", r.PathParams())
	return &u
}

func (r *CustomerCreditNoteV2Post) Do() (CustomerCreditNoteV2PostResponseBody, error) {
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
