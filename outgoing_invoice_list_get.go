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

func (c *Client) NewOutgoingInvoiceListGet() OutgoingInvoiceListGet {
	r := OutgoingInvoiceListGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OutgoingInvoiceListGet struct {
	client      *Client
	queryParams *OutgoingInvoiceListGetQueryParams
	pathParams  *OutgoingInvoiceListGetPathParams
	method      string
	headers     http.Header
	requestBody OutgoingInvoiceListGetBody
}

func (r OutgoingInvoiceListGet) NewQueryParams() *OutgoingInvoiceListGetQueryParams {
	return &OutgoingInvoiceListGetQueryParams{}
}

type OutgoingInvoiceListGetQueryParams struct {
}

func (p OutgoingInvoiceListGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OutgoingInvoiceListGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r OutgoingInvoiceListGet) NewPathParams() *OutgoingInvoiceListGetPathParams {
	return &OutgoingInvoiceListGetPathParams{}
}

type OutgoingInvoiceListGetPathParams struct {
}

func (p *OutgoingInvoiceListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OutgoingInvoiceListGet) PathParams() *OutgoingInvoiceListGetPathParams {
	return r.pathParams
}

func (r *OutgoingInvoiceListGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OutgoingInvoiceListGet) SetMethod(method string) {
	r.method = method
}

func (r *OutgoingInvoiceListGet) Method() string {
	return r.method
}

func (r OutgoingInvoiceListGet) NewRequestBody() OutgoingInvoiceListGetBody {
	return OutgoingInvoiceListGetBody{}
}

type OutgoingInvoiceListGetBody struct {
}

func (r *OutgoingInvoiceListGet) RequestBody() *OutgoingInvoiceListGetBody {
	return nil
}

func (r *OutgoingInvoiceListGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *OutgoingInvoiceListGet) SetRequestBody(body OutgoingInvoiceListGetBody) {
	r.requestBody = body
}

func (r *OutgoingInvoiceListGet) NewResponseBody() *OutgoingInvoiceListGetResponseBody {
	return &OutgoingInvoiceListGetResponseBody{}
}

type OutgoingInvoiceListGetResponseBody struct {
	Data    OutgoingInvoiceListItems `json:"data"`
	Success bool                     `json:"success"`
	Count   int                      `json:"count"`
}

func (r *OutgoingInvoiceListGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/OutgoingInvoice/List", r.PathParams())
	return &u
}

func (r *OutgoingInvoiceListGet) Do() (OutgoingInvoiceListGetResponseBody, error) {
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
