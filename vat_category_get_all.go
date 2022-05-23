package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewVATCategoryGetAll() VATCategoryGetAll {
	r := VATCategoryGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VATCategoryGetAll struct {
	client      *Client
	queryParams *VATCategoryGetAllQueryParams
	pathParams  *VATCategoryGetAllPathParams
	method      string
	headers     http.Header
	requestBody VATCategoryGetAllBody
}

func (r VATCategoryGetAll) NewQueryParams() *VATCategoryGetAllQueryParams {
	return &VATCategoryGetAllQueryParams{}
}

type VATCategoryGetAllQueryParams struct {
}

func (p VATCategoryGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VATCategoryGetAll) QueryParams() QueryParams {
	return r.queryParams
}

func (r VATCategoryGetAll) NewPathParams() *VATCategoryGetAllPathParams {
	return &VATCategoryGetAllPathParams{}
}

type VATCategoryGetAllPathParams struct {
}

func (p *VATCategoryGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATCategoryGetAll) PathParams() *VATCategoryGetAllPathParams {
	return r.pathParams
}

func (r *VATCategoryGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VATCategoryGetAll) SetMethod(method string) {
	r.method = method
}

func (r *VATCategoryGetAll) Method() string {
	return r.method
}

func (r VATCategoryGetAll) NewRequestBody() VATCategoryGetAllBody {
	return VATCategoryGetAllBody{}
}

type VATCategoryGetAllBody struct {
}

func (r *VATCategoryGetAll) RequestBody() *VATCategoryGetAllBody {
	return nil
}

func (r *VATCategoryGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *VATCategoryGetAll) SetRequestBody(body VATCategoryGetAllBody) {
	r.requestBody = body
}

func (r *VATCategoryGetAll) NewResponseBody() *VATCategoryGetAllResponseBody {
	return &VATCategoryGetAllResponseBody{}
}

type VATCategoryGetAllResponseBody VATCategories

func (r *VATCategoryGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("controller/api/v1/vatCategory", r.PathParams())
	return &u
}

func (r *VATCategoryGetAll) Do() (VATCategoryGetAllResponseBody, error) {
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
