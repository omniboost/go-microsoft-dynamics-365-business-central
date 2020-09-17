package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/odata"
	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewJournalInfoGetRequest() JournalInfoGetRequest {
	r := JournalInfoGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewJournalInfoGetQueryParams()
	r.pathParams = r.NewJournalInfoGetPathParams()
	r.requestBody = r.NewJournalInfoGetRequestBody()
	return r
}

type JournalInfoGetRequest struct {
	client      *Client
	queryParams *JournalInfoGetQueryParams
	pathParams  *JournalInfoGetPathParams
	method      string
	headers     http.Header
	requestBody JournalInfoGetRequestBody
}

func (r JournalInfoGetRequest) NewJournalInfoGetQueryParams() *JournalInfoGetQueryParams {
	body := JournalInfoGetResponseBody{{}}
	fields, _ := utils.Fields(body[0])
	return &JournalInfoGetQueryParams{
		Select:  odata.NewSelect(fields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(fields),
	}
}

type JournalInfoGetQueryParams struct {
	Select  *odata.Select  `schema:"$select,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`

	FiscalYear int `schema:"fiscalYear,omitempty"`
}

func (p JournalInfoGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalInfoGetRequest) QueryParams() *JournalInfoGetQueryParams {
	return r.queryParams
}

func (r JournalInfoGetRequest) NewJournalInfoGetPathParams() *JournalInfoGetPathParams {
	return &JournalInfoGetPathParams{}
}

type JournalInfoGetPathParams struct {
	JournalID string
}

func (p *JournalInfoGetPathParams) Params() map[string]string {
	return map[string]string{
		"journal_id": p.JournalID,
	}
}

func (r *JournalInfoGetRequest) PathParams() *JournalInfoGetPathParams {
	return r.pathParams
}

func (r *JournalInfoGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalInfoGetRequest) Method() string {
	return r.method
}

func (r JournalInfoGetRequest) NewJournalInfoGetRequestBody() JournalInfoGetRequestBody {
	return JournalInfoGetRequestBody{}
}

type JournalInfoGetRequestBody struct{}

func (r *JournalInfoGetRequest) RequestBody() *JournalInfoGetRequestBody {
	return &r.requestBody
}

func (r *JournalInfoGetRequest) SetRequestBody(body JournalInfoGetRequestBody) {
	r.requestBody = body
}

func (r *JournalInfoGetRequest) NewResponseBody() *JournalInfoGetResponseBody {
	return &JournalInfoGetResponseBody{}
}

type JournalInfoGetResponseBody []struct {
	AccountID                     string      `json:"accountId"`
	AmendmentBalance              float64     `json:"amendmentBalance"`
	BankAccountID                 string      `json:"bankAccountId"`
	CurrencyID                    string      `json:"currencyId"`
	Description                   string      `json:"description"`
	EndBalance                    float64     `json:"endBalance"`
	EntriesToMatchCount           interface{} `json:"entriesToMatchCount"`
	FiscalYear                    int         `json:"fiscalYear"`
	JournalID                     string      `json:"journalId"`
	JournalKind                   int         `json:"journalKind"`
	JournalTransaction            int         `json:"journalTransaction"`
	JournalType                   int         `json:"journalType"`
	LastElectronicTransactionDate string      `json:"lastElectronicTransactionDate"`
	LastTransactionDate           string      `json:"lastTransactionDate"`
	OpeningBalanceJournal         bool        `json:"openingBalanceJournal"`
	RangeID                       string      `json:"rangeId"`
	SequenceNumber                int         `json:"sequenceNumber"`
	SheetNumberDay                int         `json:"sheetNumberDay"`
	SheetNumberPeriod             int         `json:"sheetNumberPeriod"`
	StartBalance                  float64     `json:"startBalance"`
}

func (r *JournalInfoGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/JournalInfo/{{.journal_id}}", r.PathParams())
}

func (r *JournalInfoGetRequest) Do() (JournalInfoGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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
