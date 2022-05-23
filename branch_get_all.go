package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewBranchGetAll() BranchGetAll {
	r := BranchGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BranchGetAll struct {
	client      *Client
	queryParams *BranchGetAllQueryParams
	pathParams  *BranchGetAllPathParams
	method      string
	headers     http.Header
	requestBody BranchGetAllBody
}

func (r BranchGetAll) NewQueryParams() *BranchGetAllQueryParams {
	return &BranchGetAllQueryParams{}
}

type BranchGetAllQueryParams struct {
}

func (p BranchGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BranchGetAll) QueryParams() QueryParams {
	return r.queryParams
}

func (r BranchGetAll) NewPathParams() *BranchGetAllPathParams {
	return &BranchGetAllPathParams{}
}

type BranchGetAllPathParams struct {
}

func (p *BranchGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BranchGetAll) PathParams() *BranchGetAllPathParams {
	return r.pathParams
}

func (r *BranchGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BranchGetAll) SetMethod(method string) {
	r.method = method
}

func (r *BranchGetAll) Method() string {
	return r.method
}

func (r BranchGetAll) NewRequestBody() BranchGetAllBody {
	return BranchGetAllBody{}
}

type BranchGetAllBody struct {
}

func (r *BranchGetAll) RequestBody() *BranchGetAllBody {
	return nil
}

func (r *BranchGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *BranchGetAll) SetRequestBody(body BranchGetAllBody) {
	r.requestBody = body
}

func (r *BranchGetAll) NewResponseBody() *BranchGetAllResponseBody {
	return &BranchGetAllResponseBody{}
}

type BranchGetAllResponseBody Branches

func (r *BranchGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("controller/api/v1/branch", r.PathParams())
	return &u
}

func (r *BranchGetAll) Do() (BranchGetAllResponseBody, error) {
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
