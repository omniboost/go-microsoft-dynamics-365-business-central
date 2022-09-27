package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewVATCodesGet() VATCodesGet {
	r := VATCodesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VATCodesGet struct {
	client      *Client
	queryParams *VATCodesGetQueryParams
	pathParams  *VATCodesGetPathParams
	method      string
	headers     http.Header
	requestBody VATCodesGetBody
}

func (r VATCodesGet) NewQueryParams() *VATCodesGetQueryParams {
	return &VATCodesGetQueryParams{}
}

type VATCodesGetQueryParams struct {
}

func (p VATCodesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VATCodesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VATCodesGet) NewPathParams() *VATCodesGetPathParams {
	return &VATCodesGetPathParams{}
}

type VATCodesGetPathParams struct {
}

func (p *VATCodesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATCodesGet) PathParams() *VATCodesGetPathParams {
	return r.pathParams
}

func (r *VATCodesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VATCodesGet) SetMethod(method string) {
	r.method = method
}

func (r *VATCodesGet) Method() string {
	return r.method
}

func (r VATCodesGet) NewRequestBody() VATCodesGetBody {
	return VATCodesGetBody{}
}

type VATCodesGetBody struct {
}

func (r *VATCodesGet) RequestBody() *VATCodesGetBody {
	return nil
}

func (r *VATCodesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VATCodesGet) SetRequestBody(body VATCodesGetBody) {
	r.requestBody = body
}

func (r *VATCodesGet) NewResponseBody() *VATCodesGetResponseBody {
	return &VATCodesGetResponseBody{}
}

type VATCodesGetResponseBody struct {
	Data    VATCodes `json:"data"`
	Success bool     `json:"success"`
	Count   int      `json:"count"`
}

func (r *VATCodesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/VatCode", r.PathParams())
	return &u
}

func (r *VATCodesGet) Do() (VATCodesGetResponseBody, error) {
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
