package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewOutgoingInvoicePost() OutgoingInvoicePost {
	r := OutgoingInvoicePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OutgoingInvoicePost struct {
	client      *Client
	queryParams *OutgoingInvoicePostQueryParams
	pathParams  *OutgoingInvoicePostPathParams
	method      string
	headers     http.Header
	requestBody OutgoingInvoicePostBody
}

func (r OutgoingInvoicePost) NewQueryParams() *OutgoingInvoicePostQueryParams {
	return &OutgoingInvoicePostQueryParams{}
}

type OutgoingInvoicePostQueryParams struct{}

func (p OutgoingInvoicePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OutgoingInvoicePost) QueryParams() *OutgoingInvoicePostQueryParams {
	return r.queryParams
}

func (r OutgoingInvoicePost) NewPathParams() *OutgoingInvoicePostPathParams {
	return &OutgoingInvoicePostPathParams{}
}

type OutgoingInvoicePostPathParams struct {
}

func (p *OutgoingInvoicePostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OutgoingInvoicePost) PathParams() *OutgoingInvoicePostPathParams {
	return r.pathParams
}

func (r *OutgoingInvoicePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OutgoingInvoicePost) SetMethod(method string) {
	r.method = method
}

func (r *OutgoingInvoicePost) Method() string {
	return r.method
}

func (r OutgoingInvoicePost) NewRequestBody() OutgoingInvoicePostBody {
	return OutgoingInvoicePostBody{}
}

type OutgoingInvoicePostBody OutgoingInvoice

func (v OutgoingInvoicePostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *OutgoingInvoicePost) RequestBody() *OutgoingInvoicePostBody {
	return &r.requestBody
}

func (r *OutgoingInvoicePost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *OutgoingInvoicePost) SetRequestBody(body OutgoingInvoicePostBody) {
	r.requestBody = body
}

func (r *OutgoingInvoicePost) NewResponseBody() *OutgoingInvoicePostResponseBody {
	return &OutgoingInvoicePostResponseBody{}
}

type OutgoingInvoicePostResponseBody struct{}

func (r *OutgoingInvoicePost) URL() *url.URL {
	u := r.client.GetEndpointURL("/OutgoingInvoice/", r.PathParams())
	return &u
}

func (r *OutgoingInvoicePost) Do() (OutgoingInvoicePostResponseBody, error) {
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
