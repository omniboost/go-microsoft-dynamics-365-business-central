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

func (c *Client) NewGeneralLedgerAccountsGet() GeneralLedgerAccountsGet {
	r := GeneralLedgerAccountsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GeneralLedgerAccountsGet struct {
	client      *Client
	queryParams *GeneralLedgerAccountsGetQueryParams
	pathParams  *GeneralLedgerAccountsGetPathParams
	method      string
	headers     http.Header
	requestBody GeneralLedgerAccountsGetBody
}

func (r GeneralLedgerAccountsGet) NewQueryParams() *GeneralLedgerAccountsGetQueryParams {
	selectFields, _ := utils.Fields(&GeneralLedgerAccount{})
	expandFields := []string{}
	return &GeneralLedgerAccountsGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type GeneralLedgerAccountsGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p GeneralLedgerAccountsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GeneralLedgerAccountsGet) QueryParams() *GeneralLedgerAccountsGetQueryParams {
	return r.queryParams
}

func (r GeneralLedgerAccountsGet) NewPathParams() *GeneralLedgerAccountsGetPathParams {
	return &GeneralLedgerAccountsGetPathParams{}
}

type GeneralLedgerAccountsGetPathParams struct {
}

func (p *GeneralLedgerAccountsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GeneralLedgerAccountsGet) PathParams() *GeneralLedgerAccountsGetPathParams {
	return r.pathParams
}

func (r *GeneralLedgerAccountsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GeneralLedgerAccountsGet) SetMethod(method string) {
	r.method = method
}

func (r *GeneralLedgerAccountsGet) Method() string {
	return r.method
}

func (r GeneralLedgerAccountsGet) NewRequestBody() GeneralLedgerAccountsGetBody {
	return GeneralLedgerAccountsGetBody{}
}

type GeneralLedgerAccountsGetBody struct {
}

func (r *GeneralLedgerAccountsGet) RequestBody() *GeneralLedgerAccountsGetBody {
	return nil
}

func (r *GeneralLedgerAccountsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *GeneralLedgerAccountsGet) SetRequestBody(body GeneralLedgerAccountsGetBody) {
	r.requestBody = body
}

func (r *GeneralLedgerAccountsGet) NewResponseBody() *GeneralLedgerAccountsGetResponseBody {
	return &GeneralLedgerAccountsGetResponseBody{}
}

type GeneralLedgerAccountsGetResponseBody struct {
	Data    GeneralLedgerAccounts `json:"data"`
	Success bool                  `json:"success"`
	Count   int                   `json:"count"`
}

func (r *GeneralLedgerAccountsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedgerAccount", r.PathParams())
	return &u
}

func (r *GeneralLedgerAccountsGet) Do() (GeneralLedgerAccountsGetResponseBody, error) {
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
