package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-vismaonline/utils"
)

func (c *Client) NewTermsofpaymentsGet() TermsofpaymentsGet {
	r := TermsofpaymentsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TermsofpaymentsGet struct {
	client      *Client
	queryParams *TermsofpaymentsGetQueryParams
	pathParams  *TermsofpaymentsGetPathParams
	method      string
	headers     http.Header
	requestBody TermsofpaymentsGetBody
}

func (r TermsofpaymentsGet) NewQueryParams() *TermsofpaymentsGetQueryParams {
	return &TermsofpaymentsGetQueryParams{}
}

type TermsofpaymentsGetQueryParams struct {
}

func (p TermsofpaymentsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TermsofpaymentsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r TermsofpaymentsGet) NewPathParams() *TermsofpaymentsGetPathParams {
	return &TermsofpaymentsGetPathParams{}
}

type TermsofpaymentsGetPathParams struct {
}

func (p *TermsofpaymentsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TermsofpaymentsGet) PathParams() *TermsofpaymentsGetPathParams {
	return r.pathParams
}

func (r *TermsofpaymentsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TermsofpaymentsGet) SetMethod(method string) {
	r.method = method
}

func (r *TermsofpaymentsGet) Method() string {
	return r.method
}

func (r TermsofpaymentsGet) NewRequestBody() TermsofpaymentsGetBody {
	return TermsofpaymentsGetBody{}
}

type TermsofpaymentsGetBody struct {
}

func (r *TermsofpaymentsGet) RequestBody() *TermsofpaymentsGetBody {
	return nil
}

func (r *TermsofpaymentsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *TermsofpaymentsGet) SetRequestBody(body TermsofpaymentsGetBody) {
	r.requestBody = body
}

func (r *TermsofpaymentsGet) NewResponseBody() *TermsofpaymentsGetResponseBody {
	return &TermsofpaymentsGetResponseBody{}
}

type TermsofpaymentsGetResponseBody struct {
	Meta Meta            `json:"Meta"`
	Data TermsOfPayments `json:"Data"`
}

func (r *TermsofpaymentsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/termsofpayments", r.PathParams())
	return &u
}

func (r *TermsofpaymentsGet) Do() (TermsofpaymentsGetResponseBody, error) {
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
