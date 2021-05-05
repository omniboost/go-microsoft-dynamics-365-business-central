package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewJournalTransactionV2Post() JournalTransactionV2Post {
	r := JournalTransactionV2Post{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalTransactionV2Post struct {
	client      *Client
	queryParams *JournalTransactionV2PostQueryParams
	pathParams  *JournalTransactionV2PostPathParams
	method      string
	headers     http.Header
	requestBody JournalTransactionV2PostBody
}

func (r JournalTransactionV2Post) NewQueryParams() *JournalTransactionV2PostQueryParams {
	return &JournalTransactionV2PostQueryParams{}
}

type JournalTransactionV2PostQueryParams struct{}

func (p JournalTransactionV2PostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalTransactionV2Post) QueryParams() *JournalTransactionV2PostQueryParams {
	return r.queryParams
}

func (r JournalTransactionV2Post) NewPathParams() *JournalTransactionV2PostPathParams {
	return &JournalTransactionV2PostPathParams{}
}

type JournalTransactionV2PostPathParams struct {
}

func (p *JournalTransactionV2PostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalTransactionV2Post) PathParams() *JournalTransactionV2PostPathParams {
	return r.pathParams
}

func (r *JournalTransactionV2Post) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalTransactionV2Post) SetMethod(method string) {
	r.method = method
}

func (r *JournalTransactionV2Post) Method() string {
	return r.method
}

func (r JournalTransactionV2Post) NewRequestBody() JournalTransactionV2PostBody {
	return JournalTransactionV2PostBody{}
}

type JournalTransactionV2PostBody struct {
	BatchNumber             ValueString `json:"batchNumber"`
	Hold                    ValueBool   `json:"hold"`
	TransactionDate         ValueTime   `json:"transactionDate"`
	PostPeriod              ValueString `json:"postPeriod"`
	FinancialPeriod         ValueString `json:"financialPeriod"`
	Ledger                  ValueString `json:"ledger"`
	CurrencyID              ValueString `json:"currencyId"`
	ExchangeRate            ValueInt    `json:"exchangeRate,omitempty"`
	AutoReversing           ValueBool   `json:"autoReversing"`
	Description             ValueString `json:"description"`
	ControlTotalInCurrency  ValueInt    `json:"controlTotalInCurrency"`
	CreateVATTransaction    ValueBool   `json:"createVatTransaction"`
	SkipVATAmountValidation ValueBool   `json:"skipVatAMountValidation"`
	TransactionCode         ValueString `json:"transactionCode,omitempty"`
	Branch                  ValueString `json:"branch,omitempty"`
	OverrideNumberSeries    ValueBool   `json:"overrideNumberSeries"`
	JournalTransactionLines []struct {
		Operation     string      `json:"operation"`
		LineNumber    ValueInt    `json:"lineNumber"`
		AccountNumber ValueString `json:"accountNumber"`
		SubAccount    []struct {
			SegmentID    int    `json:"segmentId"`
			SegmentValue string `json:"segmentValue"`
		} `json:"subaccount"`
		Project                ValueString     `json:"project,omitempty"`
		PorjectTask            ValueString     `json:"projectTask,omitempty"`
		ReferenceNumber        ValueString     `json:"referenceNumber"`
		TransactionDescription ValueString     `json:"transactionDescription"`
		DebitAmountInCurrency  ValueNumber     `json:"debitAmountInCurrency"`
		CreditAmountInCurrency ValueNumber     `json:"creditAmountInCurrency"`
		VATCodeID              ValueNullString `json:"vatCodeId"`
		VATID                  ValueNullString `json:"vatId,omitempty"`
		Branch                 ValueString     `json:"branch,omitempty"`
		UOM                    ValueString     `json:"UOM,omitempty"`
		Quantity               ValueInt        `json:"quantity"`
	} `json:"journalTransactionLines"`
}

// func (r *JournalTransactionV2Post) MarshalJSON() ([]byte, error) {
// 	return omitempty.MarshalJSON(r)
// }

func (r *JournalTransactionV2Post) RequestBody() *JournalTransactionV2PostBody {
	return &r.requestBody
}

func (r *JournalTransactionV2Post) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *JournalTransactionV2Post) SetRequestBody(body JournalTransactionV2PostBody) {
	r.requestBody = body
}

func (r *JournalTransactionV2Post) NewResponseBody() *JournalTransactionV2PostResponseBody {
	return &JournalTransactionV2PostResponseBody{}
}

type JournalTransactionV2PostResponseBody JournalTransactions

func (r *JournalTransactionV2Post) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v2/journaltransaction", r.PathParams())
	return &u
}

func (r *JournalTransactionV2Post) Do() (JournalTransactionV2PostResponseBody, error) {
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
