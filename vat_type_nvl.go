package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewVATTypeNVLRequest() VATTypeNVLRequest {
	r := VATTypeNVLRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewVATTypeNVLQueryParams()
	r.pathParams = r.NewVATTypeNVLPathParams()
	r.requestBody = r.NewVATTypeNVLRequestBody()
	return r
}

type VATTypeNVLRequest struct {
	client      *Client
	queryParams *VATTypeNVLQueryParams
	pathParams  *VATTypeNVLPathParams
	method      string
	headers     http.Header
	requestBody VATTypeNVLRequestBody
}

func (r VATTypeNVLRequest) NewVATTypeNVLQueryParams() *VATTypeNVLQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &VATTypeNVLQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type VATTypeNVLQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p VATTypeNVLQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *VATTypeNVLRequest) QueryParams() *VATTypeNVLQueryParams {
	return r.queryParams
}

func (r VATTypeNVLRequest) NewVATTypeNVLPathParams() *VATTypeNVLPathParams {
	return &VATTypeNVLPathParams{}
}

type VATTypeNVLPathParams struct {
}

func (p *VATTypeNVLPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATTypeNVLRequest) PathParams() *VATTypeNVLPathParams {
	return r.pathParams
}

func (r *VATTypeNVLRequest) SetMethod(method string) {
	r.method = method
}

func (r *VATTypeNVLRequest) Method() string {
	return r.method
}

func (r VATTypeNVLRequest) NewVATTypeNVLRequestBody() VATTypeNVLRequestBody {
	return VATTypeNVLRequestBody{}
}

type VATTypeNVLRequestBody struct{}

func (r *VATTypeNVLRequest) RequestBody() *VATTypeNVLRequestBody {
	return &r.requestBody
}

func (r *VATTypeNVLRequest) SetRequestBody(body VATTypeNVLRequestBody) {
	r.requestBody = body
}

func (r *VATTypeNVLRequest) NewResponseBody() *VATTypeNVLResponseBody {
	return &VATTypeNVLResponseBody{}
}

type VATTypeNVLResponseBody []struct {
	Name  string
	Value string
}

func (r *VATTypeNVLRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/VatTypeNVL", r.PathParams())
}

func (r *VATTypeNVLRequest) Do() (VATTypeNVLResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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
