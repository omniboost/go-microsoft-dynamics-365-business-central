package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewProductsGetRequest() ProductsGetRequest {
	r := ProductsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ProductsGetRequest struct {
	client      *Client
	queryParams *ProductsGetQueryParams
	pathParams  *ProductsGetPathParams
	method      string
	headers     http.Header
	requestBody ProductsGetRequestBody
}

func (r ProductsGetRequest) NewQueryParams() *ProductsGetQueryParams {
	return &ProductsGetQueryParams{}
}

type ProductsGetQueryParams struct {
}

func (p ProductsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ProductsGetRequest) QueryParams() *ProductsGetQueryParams {
	return r.queryParams
}

func (r ProductsGetRequest) NewPathParams() *ProductsGetPathParams {
	return &ProductsGetPathParams{}
}

type ProductsGetPathParams struct {
	Objects bool `schema:"objects"`
}

func (p *ProductsGetPathParams) Params() map[string]string {
	return map[string]string{
		"objects": func() string {
			if p.Objects {
				return "true"
			}
			return "false"
		}(),
	}
}

func (r *ProductsGetRequest) PathParams() *ProductsGetPathParams {
	return r.pathParams
}

func (r *ProductsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ProductsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProductsGetRequest) Method() string {
	return r.method
}

func (r ProductsGetRequest) NewRequestBody() ProductsGetRequestBody {
	return ProductsGetRequestBody{}
}

type ProductsGetRequestBody struct {
}

func (r *ProductsGetRequest) RequestBody() *ProductsGetRequestBody {
	return &r.requestBody
}

func (r *ProductsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ProductsGetRequest) SetRequestBody(body ProductsGetRequestBody) {
	r.requestBody = body
}

func (r *ProductsGetRequest) NewResponseBody() *ProductsGetResponseBody {
	return &ProductsGetResponseBody{}
}

type ProductsGetResponseBody []ProductsGetResponseBodyProduct

type ProductsGetResponseBodyProduct Product

func (r *ProductsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("Product", r.PathParams())
	return &u
}

func (r *ProductsGetRequest) Do() (ProductsGetResponseBody, error) {
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
