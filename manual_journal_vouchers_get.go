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

func (c *Client) NewManualJournalVouchersGet() ManualJournalVouchersGet {
	r := ManualJournalVouchersGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ManualJournalVouchersGet struct {
	client      *Client
	queryParams *ManualJournalVouchersGetQueryParams
	pathParams  *ManualJournalVouchersGetPathParams
	method      string
	headers     http.Header
	requestBody ManualJournalVouchersGetBody
}

func (r ManualJournalVouchersGet) NewQueryParams() *ManualJournalVouchersGetQueryParams {
	return &ManualJournalVouchersGetQueryParams{}
}

type ManualJournalVouchersGetQueryParams struct {
}

func (p ManualJournalVouchersGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ManualJournalVouchersGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r ManualJournalVouchersGet) NewPathParams() *ManualJournalVouchersGetPathParams {
	return &ManualJournalVouchersGetPathParams{}
}

type ManualJournalVouchersGetPathParams struct {
}

func (p *ManualJournalVouchersGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ManualJournalVouchersGet) PathParams() *ManualJournalVouchersGetPathParams {
	return r.pathParams
}

func (r *ManualJournalVouchersGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ManualJournalVouchersGet) SetMethod(method string) {
	r.method = method
}

func (r *ManualJournalVouchersGet) Method() string {
	return r.method
}

func (r ManualJournalVouchersGet) NewRequestBody() ManualJournalVouchersGetBody {
	return ManualJournalVouchersGetBody{}
}

type ManualJournalVouchersGetBody struct {
}

func (r *ManualJournalVouchersGet) RequestBody() *ManualJournalVouchersGetBody {
	return nil
}

func (r *ManualJournalVouchersGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ManualJournalVouchersGet) SetRequestBody(body ManualJournalVouchersGetBody) {
	r.requestBody = body
}

func (r *ManualJournalVouchersGet) NewResponseBody() *ManualJournalVouchersGetResponseBody {
	return &ManualJournalVouchersGetResponseBody{}
}

type ManualJournalVouchersGetResponseBody struct {
	Data    ManualJournalVouchers `json:"data"`
	Success bool                  `json:"success"`
	Count   int                   `json:"count"`
}

func (r *ManualJournalVouchersGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/Voucher/ManualJournalVoucher/", r.PathParams())
	return &u
}

func (r *ManualJournalVouchersGet) Do() (ManualJournalVouchersGetResponseBody, error) {
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
