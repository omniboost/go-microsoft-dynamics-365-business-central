package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/odata"
	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewJournalTransactionInfoListGetRequest() JournalTransactionInfoListGetRequest {
	r := JournalTransactionInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewJournalTransactionInfoListGetQueryParams()
	r.pathParams = r.NewJournalTransactionInfoListGetPathParams()
	r.requestBody = r.NewJournalTransactionInfoListGetRequestBody()
	return r
}

type JournalTransactionInfoListGetRequest struct {
	client      *Client
	queryParams *JournalTransactionInfoListGetQueryParams
	pathParams  *JournalTransactionInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody JournalTransactionInfoListGetRequestBody
}

func (r JournalTransactionInfoListGetRequest) NewJournalTransactionInfoListGetQueryParams() *JournalTransactionInfoListGetQueryParams {
	body := JournalTransactionInfoListGetResponseBody{{}}
	fields, _ := utils.Fields(body[0])
	return &JournalTransactionInfoListGetQueryParams{
		Select:  odata.NewSelect(fields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(fields),
	}
}

type JournalTransactionInfoListGetQueryParams struct {
	Select  *odata.Select  `schema:"$select,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

func (p JournalTransactionInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalTransactionInfoListGetRequest) QueryParams() *JournalTransactionInfoListGetQueryParams {
	return r.queryParams
}

func (r JournalTransactionInfoListGetRequest) NewJournalTransactionInfoListGetPathParams() *JournalTransactionInfoListGetPathParams {
	return &JournalTransactionInfoListGetPathParams{}
}

type JournalTransactionInfoListGetPathParams struct {
	JournalID  string
	FiscalYear int
}

func (p *JournalTransactionInfoListGetPathParams) Params() map[string]string {
	return map[string]string{
		"journal_id":  p.JournalID,
		"fiscal_year": strconv.Itoa(p.FiscalYear),
	}
}

func (r *JournalTransactionInfoListGetRequest) PathParams() *JournalTransactionInfoListGetPathParams {
	return r.pathParams
}

func (r *JournalTransactionInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalTransactionInfoListGetRequest) Method() string {
	return r.method
}

func (r JournalTransactionInfoListGetRequest) NewJournalTransactionInfoListGetRequestBody() JournalTransactionInfoListGetRequestBody {
	return JournalTransactionInfoListGetRequestBody{}
}

type JournalTransactionInfoListGetRequestBody struct{}

func (r *JournalTransactionInfoListGetRequest) RequestBody() *JournalTransactionInfoListGetRequestBody {
	return &r.requestBody
}

func (r *JournalTransactionInfoListGetRequest) SetRequestBody(body JournalTransactionInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *JournalTransactionInfoListGetRequest) NewResponseBody() *JournalTransactionInfoListGetResponseBody {
	return &JournalTransactionInfoListGetResponseBody{}
}

type JournalTransactionInfoListGetResponseBody []struct {
	Description     string  `json:"description"`
	NewBalance      float64 `json:"newBalance"`
	PreviousBalance float64 `json:"previousBalance"`
	TransactionDate string  `json:"transactionDate"`
	TransactionID   int     `json:"transactionId"`
}

func (r *JournalTransactionInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/JournalTransactionInfoList/{{.journal_id}}/{{.fiscal_year}}", r.PathParams())
}

func (r *JournalTransactionInfoListGetRequest) Do() (JournalTransactionInfoListGetResponseBody, error) {
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
