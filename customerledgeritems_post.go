package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewCustomerledgeritemsPost() CustomerledgeritemsPost {
	r := CustomerledgeritemsPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerledgeritemsPost struct {
	client      *Client
	queryParams *CustomerledgeritemsPostQueryParams
	pathParams  *CustomerledgeritemsPostPathParams
	method      string
	headers     http.Header
	requestBody CustomerledgeritemsPostBody
}

func (r CustomerledgeritemsPost) NewQueryParams() *CustomerledgeritemsPostQueryParams {
	return &CustomerledgeritemsPostQueryParams{}
}

type CustomerledgeritemsPostQueryParams struct {
}

func (p CustomerledgeritemsPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerledgeritemsPost) QueryParams() *CustomerledgeritemsPostQueryParams {
	return r.queryParams
}

func (r CustomerledgeritemsPost) NewPathParams() *CustomerledgeritemsPostPathParams {
	return &CustomerledgeritemsPostPathParams{}
}

type CustomerledgeritemsPostPathParams struct {
}

func (p *CustomerledgeritemsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerledgeritemsPost) PathParams() *CustomerledgeritemsPostPathParams {
	return r.pathParams
}

func (r *CustomerledgeritemsPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerledgeritemsPost) SetMethod(method string) {
	r.method = method
}

func (r *CustomerledgeritemsPost) Method() string {
	return r.method
}

func (r CustomerledgeritemsPost) NewRequestBody() CustomerledgeritemsPostBody {
	return CustomerledgeritemsPostBody{}
}

type CustomerledgeritemsPostBody CustomerLedgerItem

func (v CustomerledgeritemsPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *CustomerledgeritemsPost) RequestBody() *CustomerledgeritemsPostBody {
	return &r.requestBody
}

func (r *CustomerledgeritemsPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerledgeritemsPost) SetRequestBody(body CustomerledgeritemsPostBody) {
	r.requestBody = body
}

func (r *CustomerledgeritemsPost) NewResponseBody() *CustomerledgeritemsPostResponseBody {
	return &CustomerledgeritemsPostResponseBody{}
}

type CustomerledgeritemsPostResponseBody CustomerLedgerItem

func (r *CustomerledgeritemsPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/customerledgeritems", r.PathParams())
	return &u
}

func (r *CustomerledgeritemsPost) Do() (CustomerledgeritemsPostResponseBody, error) {
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
