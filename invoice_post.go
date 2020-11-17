package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewInvoicePostRequest() InvoicePostRequest {
	r := InvoicePostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewInvoicePostQueryParams()
	r.pathParams = r.NewInvoicePostPathParams()
	r.requestBody = r.NewInvoicePostRequestBody()
	return r
}

type InvoicePostRequest struct {
	client      *Client
	queryParams *InvoicePostQueryParams
	pathParams  *InvoicePostPathParams
	method      string
	headers     http.Header
	requestBody InvoicePostRequestBody
}

func (r InvoicePostRequest) NewInvoicePostQueryParams() *InvoicePostQueryParams {
	return &InvoicePostQueryParams{}
}

type InvoicePostQueryParams struct {
	SendToCustomer bool `schema:"sendToCustomer"`
}

func (p InvoicePostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InvoicePostRequest) QueryParams() *InvoicePostQueryParams {
	return r.queryParams
}

func (r InvoicePostRequest) NewInvoicePostPathParams() *InvoicePostPathParams {
	return &InvoicePostPathParams{}
}

type InvoicePostPathParams struct {
}

func (p *InvoicePostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoicePostRequest) PathParams() *InvoicePostPathParams {
	return r.pathParams
}

func (r *InvoicePostRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicePostRequest) Method() string {
	return r.method
}

func (r InvoicePostRequest) NewInvoicePostRequestBody() InvoicePostRequestBody {
	return InvoicePostRequestBody{}
}

type InvoicePostRequestBody Invoice

func (r *InvoicePostRequest) RequestBody() *InvoicePostRequestBody {
	return &r.requestBody
}

func (r *InvoicePostRequest) SetRequestBody(body InvoicePostRequestBody) {
	r.requestBody = body
}

func (r *InvoicePostRequest) NewResponseBody() *InvoicePostResponseBody {
	return &InvoicePostResponseBody{}
}

type InvoicePostResponseBody struct{}

func (r *InvoicePostRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoice", r.PathParams())
}

func (r *InvoicePostRequest) Do() (InvoicePostResponseBody, error) {
	// fetch a new token if it isn't set already
	if r.client.token == "" {
		var err error
		r.client.token, err = r.client.NewToken()
		if err != nil {
			return *r.NewResponseBody(), err
		}
	}

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
