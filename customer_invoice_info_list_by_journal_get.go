package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/odata"
	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCustomerInvoiceInfoListByJournalGetRequest() CustomerInvoiceInfoListByJournalGetRequest {
	r := CustomerInvoiceInfoListByJournalGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerInvoiceInfoListByJournalGetQueryParams()
	r.pathParams = r.NewCustomerInvoiceInfoListByJournalGetPathParams()
	r.requestBody = r.NewCustomerInvoiceInfoListByJournalGetRequestBody()
	return r
}

type CustomerInvoiceInfoListByJournalGetRequest struct {
	client      *Client
	queryParams *CustomerInvoiceInfoListByJournalGetQueryParams
	pathParams  *CustomerInvoiceInfoListByJournalGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoiceInfoListByJournalGetRequestBody
}

func (r CustomerInvoiceInfoListByJournalGetRequest) NewCustomerInvoiceInfoListByJournalGetQueryParams() *CustomerInvoiceInfoListByJournalGetQueryParams {
	body := CustomerInvoiceInfoListByJournalGetResponseBody{{}}
	fields, _ := utils.Fields(body[0])
	return &CustomerInvoiceInfoListByJournalGetQueryParams{
		Select: odata.NewSelect(fields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CustomerInvoiceInfoListByJournalGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerInvoiceInfoListByJournalGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) QueryParams() *CustomerInvoiceInfoListByJournalGetQueryParams {
	return r.queryParams
}

func (r CustomerInvoiceInfoListByJournalGetRequest) NewCustomerInvoiceInfoListByJournalGetPathParams() *CustomerInvoiceInfoListByJournalGetPathParams {
	return &CustomerInvoiceInfoListByJournalGetPathParams{}
}

type CustomerInvoiceInfoListByJournalGetPathParams struct {
	FiscalYear   int
	JournalID    string
	InvoiceState int
}

func (p *CustomerInvoiceInfoListByJournalGetPathParams) Params() map[string]string {
	return map[string]string{
		"fiscal_year":   strconv.Itoa(p.FiscalYear),
		"journal_id":    p.JournalID,
		"invoice_state": strconv.Itoa(p.InvoiceState),
	}
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) PathParams() *CustomerInvoiceInfoListByJournalGetPathParams {
	return r.pathParams
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) Method() string {
	return r.method
}

func (r CustomerInvoiceInfoListByJournalGetRequest) NewCustomerInvoiceInfoListByJournalGetRequestBody() CustomerInvoiceInfoListByJournalGetRequestBody {
	return CustomerInvoiceInfoListByJournalGetRequestBody{}
}

type CustomerInvoiceInfoListByJournalGetRequestBody struct{}

func (r *CustomerInvoiceInfoListByJournalGetRequest) RequestBody() *CustomerInvoiceInfoListByJournalGetRequestBody {
	return &r.requestBody
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) SetRequestBody(body CustomerInvoiceInfoListByJournalGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) NewResponseBody() *CustomerInvoiceInfoListByJournalGetResponseBody {
	return &CustomerInvoiceInfoListByJournalGetResponseBody{}
}

type CustomerInvoiceInfoListByJournalGetResponseBody []CustomerInvoice

func (r *CustomerInvoiceInfoListByJournalGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CustomerInvoiceInfoList/ByJournal/{{.fiscal_year}}/{{.journal_id}}//{{.invoice_state}}", r.PathParams())
}

func (r *CustomerInvoiceInfoListByJournalGetRequest) Do() (CustomerInvoiceInfoListByJournalGetResponseBody, error) {
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
