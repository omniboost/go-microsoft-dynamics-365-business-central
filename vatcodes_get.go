package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewVatcodesGet() VatcodesGet {
	r := VatcodesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VatcodesGet struct {
	client      *Client
	queryParams *VatcodesGetQueryParams
	pathParams  *VatcodesGetPathParams
	method      string
	headers     http.Header
	requestBody VatcodesGetBody
}

func (r VatcodesGet) NewQueryParams() *VatcodesGetQueryParams {
	return &VatcodesGetQueryParams{}
}

type VatcodesGetQueryParams struct {
}

func (p VatcodesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VatcodesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VatcodesGet) NewPathParams() *VatcodesGetPathParams {
	return &VatcodesGetPathParams{}
}

type VatcodesGetPathParams struct {
}

func (p *VatcodesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VatcodesGet) PathParams() *VatcodesGetPathParams {
	return r.pathParams
}

func (r *VatcodesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VatcodesGet) SetMethod(method string) {
	r.method = method
}

func (r *VatcodesGet) Method() string {
	return r.method
}

func (r VatcodesGet) NewRequestBody() VatcodesGetBody {
	return VatcodesGetBody{}
}

type VatcodesGetBody struct {
}

func (r *VatcodesGet) RequestBody() *VatcodesGetBody {
	return nil
}

func (r *VatcodesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VatcodesGet) SetRequestBody(body VatcodesGetBody) {
	r.requestBody = body
}

func (r *VatcodesGet) NewResponseBody() *VatcodesGetResponseBody {
	return &VatcodesGetResponseBody{}
}

type VatcodesGetResponseBody struct {
	Meta Meta     `json:"Meta"`
	Data Vatcodes `json:"Data"`
}

func (r *VatcodesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/vatcodes", r.PathParams())
	return &u
}

func (r *VatcodesGet) Do() (VatcodesGetResponseBody, error) {
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
