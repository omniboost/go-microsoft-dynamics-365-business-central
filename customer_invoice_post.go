package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCustomerInvoicePostRequest() CustomerInvoicePostRequest {
	r := CustomerInvoicePostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerInvoicePostQueryParams()
	r.pathParams = r.NewCustomerInvoicePostPathParams()
	r.requestBody = r.NewCustomerInvoicePostRequestBody()
	return r
}

type CustomerInvoicePostRequest struct {
	client      *Client
	queryParams *CustomerInvoicePostQueryParams
	pathParams  *CustomerInvoicePostPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoicePostRequestBody
}

func (r CustomerInvoicePostRequest) NewCustomerInvoicePostQueryParams() *CustomerInvoicePostQueryParams {
	// selectFields, _ := utils.Fields(&CustomerInvoice{})
	return &CustomerInvoicePostQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type CustomerInvoicePostQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerInvoicePostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInvoicePostRequest) QueryParams() *CustomerInvoicePostQueryParams {
	return r.queryParams
}

func (r CustomerInvoicePostRequest) NewCustomerInvoicePostPathParams() *CustomerInvoicePostPathParams {
	return &CustomerInvoicePostPathParams{}
}

type CustomerInvoicePostPathParams struct {
}

func (p *CustomerInvoicePostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerInvoicePostRequest) PathParams() *CustomerInvoicePostPathParams {
	return r.pathParams
}

func (r *CustomerInvoicePostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoicePostRequest) Method() string {
	return r.method
}

func (r CustomerInvoicePostRequest) NewCustomerInvoicePostRequestBody() CustomerInvoicePostRequestBody {
	return CustomerInvoicePostRequestBody{}
}

type CustomerInvoicePostRequestBody CustomerInvoice

func (r *CustomerInvoicePostRequest) RequestBody() *CustomerInvoicePostRequestBody {
	return &r.requestBody
}

func (r *CustomerInvoicePostRequest) SetRequestBody(body CustomerInvoicePostRequestBody) {
	r.requestBody = body
}

func (r *CustomerInvoicePostRequest) NewResponseBody() *CustomerInvoicePostResponseBody {
	return &CustomerInvoicePostResponseBody{}
}

type CustomerInvoicePostResponseBody CustomerInvoice

func (r *CustomerInvoicePostRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CustomerInvoice", r.PathParams())
}

func (r *CustomerInvoicePostRequest) Do() (CustomerInvoicePostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
