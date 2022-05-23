package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewSalesOrderTypeGetAll() SalesOrderTypeGetAll {
	r := SalesOrderTypeGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesOrderTypeGetAll struct {
	client      *Client
	queryParams *SalesOrderTypeGetAllQueryParams
	pathParams  *SalesOrderTypeGetAllPathParams
	method      string
	headers     http.Header
	requestBody SalesOrderTypeGetAllBody
}

func (r SalesOrderTypeGetAll) NewQueryParams() *SalesOrderTypeGetAllQueryParams {
	return &SalesOrderTypeGetAllQueryParams{}
}

type SalesOrderTypeGetAllQueryParams struct{}

func (p SalesOrderTypeGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesOrderTypeGetAll) QueryParams() *SalesOrderTypeGetAllQueryParams {
	return r.queryParams
}

func (r SalesOrderTypeGetAll) NewPathParams() *SalesOrderTypeGetAllPathParams {
	return &SalesOrderTypeGetAllPathParams{}
}

type SalesOrderTypeGetAllPathParams struct {
}

func (p *SalesOrderTypeGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesOrderTypeGetAll) PathParams() *SalesOrderTypeGetAllPathParams {
	return r.pathParams
}

func (r *SalesOrderTypeGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesOrderTypeGetAll) SetMethod(method string) {
	r.method = method
}

func (r *SalesOrderTypeGetAll) Method() string {
	return r.method
}

func (r SalesOrderTypeGetAll) NewRequestBody() SalesOrderTypeGetAllBody {
	return SalesOrderTypeGetAllBody{}
}

type SalesOrderTypeGetAllBody struct{}

func (r *SalesOrderTypeGetAll) RequestBody() *SalesOrderTypeGetAllBody {
	return nil
}

func (r *SalesOrderTypeGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesOrderTypeGetAll) SetRequestBody(body SalesOrderTypeGetAllBody) {
	r.requestBody = body
}

func (r *SalesOrderTypeGetAll) NewResponseBody() *SalesOrderTypeGetAllResponseBody {
	return &SalesOrderTypeGetAllResponseBody{}
}

type SalesOrderTypeGetAllResponseBody SalesOrderTypes

func (r *SalesOrderTypeGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/salesordertype", r.PathParams())
	return &u
}

func (r *SalesOrderTypeGetAll) Do() (SalesOrderTypeGetAllResponseBody, error) {
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
