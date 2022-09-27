package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewManualJournalVoucherPost() ManualJournalVoucherPost {
	r := ManualJournalVoucherPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ManualJournalVoucherPost struct {
	client      *Client
	queryParams *ManualJournalVoucherPostQueryParams
	pathParams  *ManualJournalVoucherPostPathParams
	method      string
	headers     http.Header
	requestBody ManualJournalVoucherPostBody
}

func (r ManualJournalVoucherPost) NewQueryParams() *ManualJournalVoucherPostQueryParams {
	return &ManualJournalVoucherPostQueryParams{}
}

type ManualJournalVoucherPostQueryParams struct{}

func (p ManualJournalVoucherPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ManualJournalVoucherPost) QueryParams() *ManualJournalVoucherPostQueryParams {
	return r.queryParams
}

func (r ManualJournalVoucherPost) NewPathParams() *ManualJournalVoucherPostPathParams {
	return &ManualJournalVoucherPostPathParams{}
}

type ManualJournalVoucherPostPathParams struct {
}

func (p *ManualJournalVoucherPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ManualJournalVoucherPost) PathParams() *ManualJournalVoucherPostPathParams {
	return r.pathParams
}

func (r *ManualJournalVoucherPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ManualJournalVoucherPost) SetMethod(method string) {
	r.method = method
}

func (r *ManualJournalVoucherPost) Method() string {
	return r.method
}

func (r ManualJournalVoucherPost) NewRequestBody() ManualJournalVoucherPostBody {
	return ManualJournalVoucherPostBody{}
}

type ManualJournalVoucherPostBody ManualJournalVoucher

func (v ManualJournalVoucherPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *ManualJournalVoucherPost) RequestBody() *ManualJournalVoucherPostBody {
	return &r.requestBody
}

func (r *ManualJournalVoucherPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ManualJournalVoucherPost) SetRequestBody(body ManualJournalVoucherPostBody) {
	r.requestBody = body
}

func (r *ManualJournalVoucherPost) NewResponseBody() *ManualJournalVoucherPostResponseBody {
	return &ManualJournalVoucherPostResponseBody{}
}

type ManualJournalVoucherPostResponseBody struct{}

func (r *ManualJournalVoucherPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/Voucher/ManualJournalVoucher/", r.PathParams())
	return &u
}

func (r *ManualJournalVoucherPost) Do() (ManualJournalVoucherPostResponseBody, error) {
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
