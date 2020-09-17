package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewFinTransPostRequest() FinTransPostRequest {
	r := FinTransPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewFinTransPostQueryParams()
	r.pathParams = r.NewFinTransPostPathParams()
	r.requestBody = r.NewFinTransPostRequestBody()
	return r
}

type FinTransPostRequest struct {
	client      *Client
	queryParams *FinTransPostQueryParams
	pathParams  *FinTransPostPathParams
	method      string
	headers     http.Header
	requestBody FinTransPostRequestBody
}

func (r FinTransPostRequest) NewFinTransPostQueryParams() *FinTransPostQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &FinTransPostQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type FinTransPostQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p FinTransPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinTransPostRequest) QueryParams() *FinTransPostQueryParams {
	return r.queryParams
}

func (r FinTransPostRequest) NewFinTransPostPathParams() *FinTransPostPathParams {
	return &FinTransPostPathParams{}
}

type FinTransPostPathParams struct {
}

func (p *FinTransPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinTransPostRequest) PathParams() *FinTransPostPathParams {
	return r.pathParams
}

func (r *FinTransPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinTransPostRequest) Method() string {
	return r.method
}

func (r FinTransPostRequest) NewFinTransPostRequestBody() FinTransPostRequestBody {
	return FinTransPostRequestBody{}
}

type FinTransPostRequestBody FinTrans

func (r *FinTransPostRequest) RequestBody() *FinTransPostRequestBody {
	return &r.requestBody
}

func (r *FinTransPostRequest) SetRequestBody(body FinTransPostRequestBody) {
	r.requestBody = body
}

func (r *FinTransPostRequest) NewResponseBody() *FinTransPostResponseBody {
	return &FinTransPostResponseBody{}
}

type FinTransPostResponseBody struct{}

func (r *FinTransPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/FinTrans", r.PathParams())
}

func (r *FinTransPostRequest) Do() (FinTransPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
