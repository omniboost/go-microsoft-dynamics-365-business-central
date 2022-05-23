package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerGetByCD() CustomerGetByCD {
	r := CustomerGetByCD{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerGetByCD struct {
	client      *Client
	queryParams *CustomerGetByCDQueryParams
	pathParams  *CustomerGetByCDPathParams
	method      string
	headers     http.Header
	requestBody CustomerGetByCDBody
}

func (r CustomerGetByCD) NewQueryParams() *CustomerGetByCDQueryParams {
	return &CustomerGetByCDQueryParams{}
}

type CustomerGetByCDQueryParams struct{}

func (p CustomerGetByCDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerGetByCD) QueryParams() *CustomerGetByCDQueryParams {
	return r.queryParams
}

func (r CustomerGetByCD) NewPathParams() *CustomerGetByCDPathParams {
	return &CustomerGetByCDPathParams{}
}

type CustomerGetByCDPathParams struct {
	CustomerCD string `schema:"customer_cd"`
}

func (p *CustomerGetByCDPathParams) Params() map[string]string {
	return map[string]string{
		"customer_cd": p.CustomerCD,
	}
}

func (r *CustomerGetByCD) PathParams() *CustomerGetByCDPathParams {
	return r.pathParams
}

func (r *CustomerGetByCD) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerGetByCD) SetMethod(method string) {
	r.method = method
}

func (r *CustomerGetByCD) Method() string {
	return r.method
}

func (r CustomerGetByCD) NewRequestBody() CustomerGetByCDBody {
	return CustomerGetByCDBody{}
}

type CustomerGetByCDBody struct {
}

func (r *CustomerGetByCD) RequestBody() *CustomerGetByCDBody {
	return nil
}

func (r *CustomerGetByCD) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerGetByCD) SetRequestBody(body CustomerGetByCDBody) {
	r.requestBody = body
}

func (r *CustomerGetByCD) NewResponseBody() *CustomerGetByCDResponseBody {
	return &CustomerGetByCDResponseBody{}
}

type CustomerGetByCDResponseBody Customer

func (r *CustomerGetByCD) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customer/{{.customer_cd}}", r.PathParams())
	return &u
}

func (r *CustomerGetByCD) Do() (CustomerGetByCDResponseBody, error) {
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
