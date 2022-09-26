package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewVouchersPost() VouchersPost {
	r := VouchersPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VouchersPost struct {
	client      *Client
	queryParams *VouchersPostQueryParams
	pathParams  *VouchersPostPathParams
	method      string
	headers     http.Header
	requestBody VouchersPostBody
}

func (r VouchersPost) NewQueryParams() *VouchersPostQueryParams {
	return &VouchersPostQueryParams{
		UseAutomaticVATCalculation: false,
		UseDefaultVATCodes:         true,
		UseDefaultVoucherSeries:    true,
	}
}

type VouchersPostQueryParams struct {
	UseAutomaticVATCalculation bool `schema:"useAutomaticVatCalculation"`
	UseDefaultVATCodes         bool `schema:"useDefaultVatCodes"`
	UseDefaultVoucherSeries    bool `schema:"useDefaultVoucherSeries"`
}

func (p VouchersPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VouchersPost) QueryParams() *VouchersPostQueryParams {
	return r.queryParams
}

func (r VouchersPost) NewPathParams() *VouchersPostPathParams {
	return &VouchersPostPathParams{}
}

type VouchersPostPathParams struct {
}

func (p *VouchersPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VouchersPost) PathParams() *VouchersPostPathParams {
	return r.pathParams
}

func (r *VouchersPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VouchersPost) SetMethod(method string) {
	r.method = method
}

func (r *VouchersPost) Method() string {
	return r.method
}

func (r VouchersPost) NewRequestBody() VouchersPostBody {
	return VouchersPostBody{}
}

type VouchersPostBody Voucher

func (v VouchersPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *VouchersPost) RequestBody() *VouchersPostBody {
	return &r.requestBody
}

func (r *VouchersPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *VouchersPost) SetRequestBody(body VouchersPostBody) {
	r.requestBody = body
}

func (r *VouchersPost) NewResponseBody() *VouchersPostResponseBody {
	return &VouchersPostResponseBody{}
}

type VouchersPostResponseBody Voucher

func (r *VouchersPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/vouchers", r.PathParams())
	return &u
}

func (r *VouchersPost) Do() (VouchersPostResponseBody, error) {
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
