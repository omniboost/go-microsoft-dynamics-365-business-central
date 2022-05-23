package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerCreditNoteGetAll() CustomerCreditNoteGetAll {
	r := CustomerCreditNoteGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerCreditNoteGetAll struct {
	client      *Client
	queryParams *CustomerCreditNoteGetAllQueryParams
	pathParams  *CustomerCreditNoteGetAllPathParams
	method      string
	headers     http.Header
	requestBody CustomerCreditNoteGetAllBody
}

func (r CustomerCreditNoteGetAll) NewQueryParams() *CustomerCreditNoteGetAllQueryParams {
	return &CustomerCreditNoteGetAllQueryParams{}
}

type CustomerCreditNoteGetAllQueryParams struct{}

func (p CustomerCreditNoteGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerCreditNoteGetAll) QueryParams() *CustomerCreditNoteGetAllQueryParams {
	return r.queryParams
}

func (r CustomerCreditNoteGetAll) NewPathParams() *CustomerCreditNoteGetAllPathParams {
	return &CustomerCreditNoteGetAllPathParams{}
}

type CustomerCreditNoteGetAllPathParams struct {
}

func (p *CustomerCreditNoteGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerCreditNoteGetAll) PathParams() *CustomerCreditNoteGetAllPathParams {
	return r.pathParams
}

func (r *CustomerCreditNoteGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerCreditNoteGetAll) SetMethod(method string) {
	r.method = method
}

func (r *CustomerCreditNoteGetAll) Method() string {
	return r.method
}

func (r CustomerCreditNoteGetAll) NewRequestBody() CustomerCreditNoteGetAllBody {
	return CustomerCreditNoteGetAllBody{}
}

type CustomerCreditNoteGetAllBody struct {
}

func (r *CustomerCreditNoteGetAll) RequestBody() *CustomerCreditNoteGetAllBody {
	return nil
}

func (r *CustomerCreditNoteGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerCreditNoteGetAll) SetRequestBody(body CustomerCreditNoteGetAllBody) {
	r.requestBody = body
}

func (r *CustomerCreditNoteGetAll) NewResponseBody() *CustomerCreditNoteGetAllResponseBody {
	return &CustomerCreditNoteGetAllResponseBody{}
}

type CustomerCreditNoteGetAllResponseBody Currencies

func (r *CustomerCreditNoteGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customerCreditNote", r.PathParams())
	return &u
}

func (r *CustomerCreditNoteGetAll) Do() (CustomerCreditNoteGetAllResponseBody, error) {
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
