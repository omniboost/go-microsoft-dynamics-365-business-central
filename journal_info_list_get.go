package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/odata"
	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewJournalInfoListGetRequest() JournalInfoListGetRequest {
	r := JournalInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewJournalInfoListGetQueryParams()
	r.pathParams = r.NewJournalInfoListGetPathParams()
	r.requestBody = r.NewJournalInfoListGetRequestBody()
	return r
}

type JournalInfoListGetRequest struct {
	client      *Client
	queryParams *JournalInfoListGetQueryParams
	pathParams  *JournalInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody JournalInfoListGetRequestBody
}

func (r JournalInfoListGetRequest) NewJournalInfoListGetQueryParams() *JournalInfoListGetQueryParams {
	fields, _ := utils.Fields(&JournalInfoListGetRequest{})
	return &JournalInfoListGetQueryParams{
		Select:  odata.NewSelect(fields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(fields),
	}
}

type JournalInfoListGetQueryParams struct {
	Select  *odata.Select  `schema:"$select,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderBy,omitempty"`
}

func (p JournalInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalInfoListGetRequest) QueryParams() *JournalInfoListGetQueryParams {
	return r.queryParams
}

func (r JournalInfoListGetRequest) NewJournalInfoListGetPathParams() *JournalInfoListGetPathParams {
	return &JournalInfoListGetPathParams{}
}

type JournalInfoListGetPathParams struct {
}

func (p *JournalInfoListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalInfoListGetRequest) PathParams() *JournalInfoListGetPathParams {
	return r.pathParams
}

func (r *JournalInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalInfoListGetRequest) Method() string {
	return r.method
}

func (r JournalInfoListGetRequest) NewJournalInfoListGetRequestBody() JournalInfoListGetRequestBody {
	return JournalInfoListGetRequestBody{}
}

type JournalInfoListGetRequestBody struct{}

func (r *JournalInfoListGetRequest) RequestBody() *JournalInfoListGetRequestBody {
	return &r.requestBody
}

func (r *JournalInfoListGetRequest) SetRequestBody(body JournalInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *JournalInfoListGetRequest) NewResponseBody() *JournalInfoListGetResponseBody {
	return &JournalInfoListGetResponseBody{}
}

type JournalInfoListGetResponseBody []struct {
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

func (r *JournalInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/JournalInfoList", r.PathParams())
}

func (r *JournalInfoListGetRequest) Do() (JournalInfoListGetResponseBody, error) {
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
