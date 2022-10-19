package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/odata"
	"github.com/omniboost/go-poweroffice/utils"
)

// Manual Journal Voucher service used to query and create and post
// ManualJournalVoucher. (Norwegian: Manuelt bilag) Querying this service will
// only return the vouchers created by the integration itself, not all vouchers
// of type One.Domain.Entities.Accounting.Vouchers.VoucherType.ManualJournal.

func (c *Client) NewCustomersGet() CustomersGet {
	r := CustomersGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomersGet struct {
	client      *Client
	queryParams *CustomersGetQueryParams
	pathParams  *CustomersGetPathParams
	method      string
	headers     http.Header
	requestBody CustomersGetBody
}

func (r CustomersGet) NewQueryParams() *CustomersGetQueryParams {
	selectFields, _ := utils.Fields(&Customer{})
	expandFields := []string{}
	return &CustomersGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CustomersGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomersGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomersGet) QueryParams() *CustomersGetQueryParams {
	return r.queryParams
}

func (r CustomersGet) NewPathParams() *CustomersGetPathParams {
	return &CustomersGetPathParams{}
}

type CustomersGetPathParams struct {
}

func (p *CustomersGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomersGet) PathParams() *CustomersGetPathParams {
	return r.pathParams
}

func (r *CustomersGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomersGet) SetMethod(method string) {
	r.method = method
}

func (r *CustomersGet) Method() string {
	return r.method
}

func (r CustomersGet) NewRequestBody() CustomersGetBody {
	return CustomersGetBody{}
}

type CustomersGetBody struct {
}

func (r *CustomersGet) RequestBody() *CustomersGetBody {
	return nil
}

func (r *CustomersGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomersGet) SetRequestBody(body CustomersGetBody) {
	r.requestBody = body
}

func (r *CustomersGet) NewResponseBody() *CustomersGetResponseBody {
	return &CustomersGetResponseBody{}
}

type CustomersGetResponseBody struct {
	Data    Customers `json:"data"`
	Success bool      `json:"success"`
	Count   int       `json:"count"`
}

func (r *CustomersGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/Customer", r.PathParams())
	return &u
}

func (r *CustomersGet) Do() (CustomersGetResponseBody, error) {
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
