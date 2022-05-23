package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewVATGetAll() VATGetAll {
	r := VATGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VATGetAll struct {
	client      *Client
	queryParams *VATGetAllQueryParams
	pathParams  *VATGetAllPathParams
	method      string
	headers     http.Header
	requestBody VATGetAllBody
}

func (r VATGetAll) NewQueryParams() *VATGetAllQueryParams {
	return &VATGetAllQueryParams{}
}

type VATGetAllQueryParams struct {
}

func (p VATGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VATGetAll) QueryParams() QueryParams {
	return r.queryParams
}

func (r VATGetAll) NewPathParams() *VATGetAllPathParams {
	return &VATGetAllPathParams{}
}

type VATGetAllPathParams struct {
}

func (p *VATGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATGetAll) PathParams() *VATGetAllPathParams {
	return r.pathParams
}

func (r *VATGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VATGetAll) SetMethod(method string) {
	r.method = method
}

func (r *VATGetAll) Method() string {
	return r.method
}

func (r VATGetAll) NewRequestBody() VATGetAllBody {
	return VATGetAllBody{}
}

type VATGetAllBody struct {
}

func (r *VATGetAll) RequestBody() *VATGetAllBody {
	return nil
}

func (r *VATGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *VATGetAll) SetRequestBody(body VATGetAllBody) {
	r.requestBody = body
}

func (r *VATGetAll) NewResponseBody() *VATGetAllResponseBody {
	return &VATGetAllResponseBody{}
}

type VATGetAllResponseBody VATs

func (r *VATGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("controller/api/v1/vat", r.PathParams())
	return &u
}

func (r *VATGetAll) Do() (VATGetAllResponseBody, error) {
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
