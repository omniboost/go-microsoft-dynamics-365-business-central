package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/odata"
	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCustomerInfoListGetRequest() CustomerInfoListGetRequest {
	r := CustomerInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerInfoListGetQueryParams()
	r.pathParams = r.NewCustomerInfoListGetPathParams()
	r.requestBody = r.NewCustomerInfoListGetRequestBody()
	return r
}

type CustomerInfoListGetRequest struct {
	client      *Client
	queryParams *CustomerInfoListGetQueryParams
	pathParams  *CustomerInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerInfoListGetRequestBody
}

func (r CustomerInfoListGetRequest) NewCustomerInfoListGetQueryParams() *CustomerInfoListGetQueryParams {
	body := CustomerInfoListGetResponseBody{{}}
	fields, _ := utils.Fields(body[0])
	return &CustomerInfoListGetQueryParams{
		Select: odata.NewSelect(fields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CustomerInfoListGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInfoListGetRequest) QueryParams() *CustomerInfoListGetQueryParams {
	return r.queryParams
}

func (r CustomerInfoListGetRequest) NewCustomerInfoListGetPathParams() *CustomerInfoListGetPathParams {
	return &CustomerInfoListGetPathParams{}
}

type CustomerInfoListGetPathParams struct {
}

func (p *CustomerInfoListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerInfoListGetRequest) PathParams() *CustomerInfoListGetPathParams {
	return r.pathParams
}

func (r *CustomerInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInfoListGetRequest) Method() string {
	return r.method
}

func (r CustomerInfoListGetRequest) NewCustomerInfoListGetRequestBody() CustomerInfoListGetRequestBody {
	return CustomerInfoListGetRequestBody{}
}

type CustomerInfoListGetRequestBody struct{}

func (r *CustomerInfoListGetRequest) RequestBody() *CustomerInfoListGetRequestBody {
	return &r.requestBody
}

func (r *CustomerInfoListGetRequest) SetRequestBody(body CustomerInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerInfoListGetRequest) NewResponseBody() *CustomerInfoListGetResponseBody {
	return &CustomerInfoListGetResponseBody{}
}

type CustomerInfoListGetResponseBody []Customer

func (r *CustomerInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CustomerInfoList", r.PathParams())
}

func (r *CustomerInfoListGetRequest) Do() (CustomerInfoListGetResponseBody, error) {
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
