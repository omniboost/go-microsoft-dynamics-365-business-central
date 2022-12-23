package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/utils"
)

// Manual Journal Voucher service used to query and create and post
// ManualJournalVoucher. (Norwegian: Manuelt bilag) Querying this service will
// only return the vouchers created by the integration itself, not all vouchers
// of type One.Domain.Entities.Accounting.Vouchers.VoucherType.ManualJournal.

func (c *Client) NewOutgoingInvoiceIDGet() OutgoingInvoiceIDGet {
	r := OutgoingInvoiceIDGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OutgoingInvoiceIDGet struct {
	client      *Client
	queryParams *OutgoingInvoiceIDGetQueryParams
	pathParams  *OutgoingInvoiceIDGetPathParams
	method      string
	headers     http.Header
	requestBody OutgoingInvoiceIDGetBody
}

func (r OutgoingInvoiceIDGet) NewQueryParams() *OutgoingInvoiceIDGetQueryParams {
	return &OutgoingInvoiceIDGetQueryParams{}
}

type OutgoingInvoiceIDGetQueryParams struct {
}

func (p OutgoingInvoiceIDGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OutgoingInvoiceIDGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r OutgoingInvoiceIDGet) NewPathParams() *OutgoingInvoiceIDGetPathParams {
	return &OutgoingInvoiceIDGetPathParams{}
}

type OutgoingInvoiceIDGetPathParams struct {
	ID string `schema:"id"`
}

func (p *OutgoingInvoiceIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *OutgoingInvoiceIDGet) PathParams() *OutgoingInvoiceIDGetPathParams {
	return r.pathParams
}

func (r *OutgoingInvoiceIDGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OutgoingInvoiceIDGet) SetMethod(method string) {
	r.method = method
}

func (r *OutgoingInvoiceIDGet) Method() string {
	return r.method
}

func (r OutgoingInvoiceIDGet) NewRequestBody() OutgoingInvoiceIDGetBody {
	return OutgoingInvoiceIDGetBody{}
}

type OutgoingInvoiceIDGetBody struct {
}

func (r *OutgoingInvoiceIDGet) RequestBody() *OutgoingInvoiceIDGetBody {
	return nil
}

func (r *OutgoingInvoiceIDGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *OutgoingInvoiceIDGet) SetRequestBody(body OutgoingInvoiceIDGetBody) {
	r.requestBody = body
}

func (r *OutgoingInvoiceIDGet) NewResponseBody() *OutgoingInvoiceIDGetResponseBody {
	return &OutgoingInvoiceIDGetResponseBody{}
}

type OutgoingInvoiceIDGetResponseBody struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Count   int         `json:"count"`
}

func (r *OutgoingInvoiceIDGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/OutgoingInvoice/{{.id}}", r.PathParams())
	return &u
}

func (r *OutgoingInvoiceIDGet) Do() (OutgoingInvoiceIDGetResponseBody, error) {
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
