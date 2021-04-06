package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCurrencyGetAll() CurrencyGetAll {
	r := CurrencyGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CurrencyGetAll struct {
	client      *Client
	queryParams *CurrencyGetAllQueryParams
	pathParams  *CurrencyGetAllPathParams
	method      string
	headers     http.Header
	requestBody CurrencyGetAllBody
}

func (r CurrencyGetAll) NewQueryParams() *CurrencyGetAllQueryParams {
	return &CurrencyGetAllQueryParams{}
}

type CurrencyGetAllQueryParams struct{}

func (p CurrencyGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CurrencyGetAll) QueryParams() *CurrencyGetAllQueryParams {
	return r.queryParams
}

func (r CurrencyGetAll) NewPathParams() *CurrencyGetAllPathParams {
	return &CurrencyGetAllPathParams{}
}

type CurrencyGetAllPathParams struct {
}

func (p *CurrencyGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CurrencyGetAll) PathParams() *CurrencyGetAllPathParams {
	return r.pathParams
}

func (r *CurrencyGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CurrencyGetAll) SetMethod(method string) {
	r.method = method
}

func (r *CurrencyGetAll) Method() string {
	return r.method
}

func (r CurrencyGetAll) NewRequestBody() CurrencyGetAllBody {
	return CurrencyGetAllBody{}
}

type CurrencyGetAllBody struct {
}

func (r *CurrencyGetAll) RequestBody() *CurrencyGetAllBody {
	return nil
}

func (r *CurrencyGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *CurrencyGetAll) SetRequestBody(body CurrencyGetAllBody) {
	r.requestBody = body
}

func (r *CurrencyGetAll) NewResponseBody() *CurrencyGetAllResponseBody {
	return &CurrencyGetAllResponseBody{}
}

type CurrencyGetAllResponseBody Currencies

func (r *CurrencyGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/currency", r.PathParams())
	return &u
}

func (r *CurrencyGetAll) Do() (CurrencyGetAllResponseBody, error) {
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
