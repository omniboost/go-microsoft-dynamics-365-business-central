package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCustomerInvoiceGetRequest() CustomerInvoiceGetRequest {
	r := CustomerInvoiceGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerInvoiceGetQueryParams()
	r.pathParams = r.NewCustomerInvoiceGetPathParams()
	r.requestBody = r.NewCustomerInvoiceGetRequestBody()
	return r
}

type CustomerInvoiceGetRequest struct {
	client      *Client
	queryParams *CustomerInvoiceGetQueryParams
	pathParams  *CustomerInvoiceGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoiceGetRequestBody
}

func (r CustomerInvoiceGetRequest) NewCustomerInvoiceGetQueryParams() *CustomerInvoiceGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &CustomerInvoiceGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type CustomerInvoiceGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerInvoiceGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInvoiceGetRequest) QueryParams() *CustomerInvoiceGetQueryParams {
	return r.queryParams
}

func (r CustomerInvoiceGetRequest) NewCustomerInvoiceGetPathParams() *CustomerInvoiceGetPathParams {
	return &CustomerInvoiceGetPathParams{}
}

type CustomerInvoiceGetPathParams struct {
	InvoiceID int
}

func (p *CustomerInvoiceGetPathParams) Params() map[string]string {
	return map[string]string{
		"invoice_id": strconv.Itoa(p.InvoiceID),
	}
}

func (r *CustomerInvoiceGetRequest) PathParams() *CustomerInvoiceGetPathParams {
	return r.pathParams
}

func (r *CustomerInvoiceGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoiceGetRequest) Method() string {
	return r.method
}

func (r CustomerInvoiceGetRequest) NewCustomerInvoiceGetRequestBody() CustomerInvoiceGetRequestBody {
	return CustomerInvoiceGetRequestBody{}
}

type CustomerInvoiceGetRequestBody struct{}

func (r *CustomerInvoiceGetRequest) RequestBody() *CustomerInvoiceGetRequestBody {
	return &r.requestBody
}

func (r *CustomerInvoiceGetRequest) SetRequestBody(body CustomerInvoiceGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerInvoiceGetRequest) NewResponseBody() *CustomerInvoiceGetResponseBody {
	return &CustomerInvoiceGetResponseBody{}
}

type CustomerInvoiceGetResponseBody struct {
	BackupDate          string        `json:"backupDate"`
	Version             int           `json:"version"`
	GroupID             int           `json:"groupId"`
	Code                string        `json:"code"`
	ShortName           string        `json:"shortName"`
	PreviousOnlineState int           `json:"previousOnlineState"`
	OnlineState         int           `json:"onlineState"`
	Description         string        `json:"description"`
	AdministrationID    string        `json:"administrationId"`
	Users               []string      `json:"users"`
	ReportPath          string        `json:"reportPath"`
	Messages            []interface{} `json:"messages"`
	CanChange           bool          `json:"canChange"`
	CannotChangeReason  string        `json:"cannotChangeReason"`
}

func (r *CustomerInvoiceGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CustomerInvoice/{{.invoice_id}}", r.PathParams())
}

func (r *CustomerInvoiceGetRequest) Do() (CustomerInvoiceGetResponseBody, error) {
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
