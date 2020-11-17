package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewOrderListPostRequest() OrderListPostRequest {
	r := OrderListPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewOrderListPostQueryParams()
	r.pathParams = r.NewOrderListPostPathParams()
	r.requestBody = r.NewOrderListPostRequestBody()
	return r
}

type OrderListPostRequest struct {
	client      *Client
	queryParams *OrderListPostQueryParams
	pathParams  *OrderListPostPathParams
	method      string
	headers     http.Header
	requestBody OrderListPostRequestBody
}

func (r OrderListPostRequest) NewOrderListPostQueryParams() *OrderListPostQueryParams {
	return &OrderListPostQueryParams{}
}

type OrderListPostQueryParams struct{}

func (p OrderListPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OrderListPostRequest) QueryParams() *OrderListPostQueryParams {
	return r.queryParams
}

func (r OrderListPostRequest) NewOrderListPostPathParams() *OrderListPostPathParams {
	return &OrderListPostPathParams{}
}

type OrderListPostPathParams struct {
}

func (p *OrderListPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OrderListPostRequest) PathParams() *OrderListPostPathParams {
	return r.pathParams
}

func (r *OrderListPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *OrderListPostRequest) Method() string {
	return r.method
}

func (r OrderListPostRequest) NewOrderListPostRequestBody() OrderListPostRequestBody {
	return OrderListPostRequestBody{}
}

type OrderListPostRequestBody Orders

func (r *OrderListPostRequest) RequestBody() *OrderListPostRequestBody {
	return &r.requestBody
}

func (r *OrderListPostRequest) SetRequestBody(body OrderListPostRequestBody) {
	r.requestBody = body
}

func (r *OrderListPostRequest) NewResponseBody() *OrderListPostResponseBody {
	return &OrderListPostResponseBody{}
}

type OrderListPostResponseBody struct{}

func (r *OrderListPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("/order/list", r.PathParams())
}

func (r *OrderListPostRequest) Do() (OrderListPostResponseBody, error) {
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
