package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewCustomerledgeritemswithvoucherPost() CustomerledgeritemswithvoucherPost {
	r := CustomerledgeritemswithvoucherPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerledgeritemswithvoucherPost struct {
	client      *Client
	queryParams *CustomerledgeritemswithvoucherPostQueryParams
	pathParams  *CustomerledgeritemswithvoucherPostPathParams
	method      string
	headers     http.Header
	requestBody CustomerledgeritemswithvoucherPostBody
}

func (r CustomerledgeritemswithvoucherPost) NewQueryParams() *CustomerledgeritemswithvoucherPostQueryParams {
	return &CustomerledgeritemswithvoucherPostQueryParams{}
}

type CustomerledgeritemswithvoucherPostQueryParams struct {
}

func (p CustomerledgeritemswithvoucherPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerledgeritemswithvoucherPost) QueryParams() *CustomerledgeritemswithvoucherPostQueryParams {
	return r.queryParams
}

func (r CustomerledgeritemswithvoucherPost) NewPathParams() *CustomerledgeritemswithvoucherPostPathParams {
	return &CustomerledgeritemswithvoucherPostPathParams{}
}

type CustomerledgeritemswithvoucherPostPathParams struct {
}

func (p *CustomerledgeritemswithvoucherPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerledgeritemswithvoucherPost) PathParams() *CustomerledgeritemswithvoucherPostPathParams {
	return r.pathParams
}

func (r *CustomerledgeritemswithvoucherPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerledgeritemswithvoucherPost) SetMethod(method string) {
	r.method = method
}

func (r *CustomerledgeritemswithvoucherPost) Method() string {
	return r.method
}

func (r CustomerledgeritemswithvoucherPost) NewRequestBody() CustomerledgeritemswithvoucherPostBody {
	return CustomerledgeritemswithvoucherPostBody{}
}

type CustomerledgeritemswithvoucherPostBody CustomerLedgerItem

func (v CustomerledgeritemswithvoucherPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *CustomerledgeritemswithvoucherPost) RequestBody() *CustomerledgeritemswithvoucherPostBody {
	return &r.requestBody
}

func (r *CustomerledgeritemswithvoucherPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerledgeritemswithvoucherPost) SetRequestBody(body CustomerledgeritemswithvoucherPostBody) {
	r.requestBody = body
}

func (r *CustomerledgeritemswithvoucherPost) NewResponseBody() *CustomerledgeritemswithvoucherPostResponseBody {
	return &CustomerledgeritemswithvoucherPostResponseBody{}
}

type CustomerledgeritemswithvoucherPostResponseBody CustomerLedgerItem

func (r *CustomerledgeritemswithvoucherPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/customerledgeritems/customerledgeritemswithvoucher", r.PathParams())
	return &u
}

func (r *CustomerledgeritemswithvoucherPost) Do() (CustomerledgeritemswithvoucherPostResponseBody, error) {
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
