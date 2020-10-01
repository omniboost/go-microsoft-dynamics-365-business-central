package tripletex

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewTokenSessionCreateRequest() TokenSessionCreateRequest {
	r := TokenSessionCreateRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewTokenSessionCreateQueryParams()
	r.pathParams = r.NewTokenSessionCreatePathParams()
	r.requestBody = r.NewTokenSessionCreateRequestBody()
	return r
}

type TokenSessionCreateRequest struct {
	client      *Client
	queryParams *TokenSessionCreateQueryParams
	pathParams  *TokenSessionCreatePathParams
	method      string
	headers     http.Header
	requestBody TokenSessionCreateRequestBody
}

func (r TokenSessionCreateRequest) NewTokenSessionCreateQueryParams() *TokenSessionCreateQueryParams {
	return &TokenSessionCreateQueryParams{}
}

type TokenSessionCreateQueryParams struct {
	ConsumerToken  string `schema:"consumerToken"`
	EmployeeToken  string `schema:"employeeToken"`
	ExpirationDate Date   `schema:"expirationDate"`
}

func (p TokenSessionCreateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TokenSessionCreateRequest) QueryParams() *TokenSessionCreateQueryParams {
	return r.queryParams
}

func (r TokenSessionCreateRequest) NewTokenSessionCreatePathParams() *TokenSessionCreatePathParams {
	return &TokenSessionCreatePathParams{}
}

type TokenSessionCreatePathParams struct {
}

func (p *TokenSessionCreatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TokenSessionCreateRequest) PathParams() *TokenSessionCreatePathParams {
	return r.pathParams
}

func (r *TokenSessionCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *TokenSessionCreateRequest) Method() string {
	return r.method
}

func (r TokenSessionCreateRequest) NewTokenSessionCreateRequestBody() TokenSessionCreateRequestBody {
	return TokenSessionCreateRequestBody{}
}

type TokenSessionCreateRequestBody struct{}

func (r *TokenSessionCreateRequest) RequestBody() *TokenSessionCreateRequestBody {
	return &r.requestBody
}

func (r *TokenSessionCreateRequest) SetRequestBody(body TokenSessionCreateRequestBody) {
	r.requestBody = body
}

func (r *TokenSessionCreateRequest) NewResponseBody() *TokenSessionCreateResponseBody {
	return &TokenSessionCreateResponseBody{}
}

type TokenSessionCreateResponseBody struct {
	Value struct {
		ID            int    `json:"id"`
		Version       int    `json:"version"`
		URL           string `json:"url"`
		ConsumerToken struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"consumerToken"`
		EmployeeToken struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"employeeToken"`
		ExpirationDate string      `json:"expirationDate"`
		Token          string      `json:"token"`
		EncryptionKey  interface{} `json:"encryptionKey"`
	} `json:"value"`
}

func (r *TokenSessionCreateRequest) URL() url.URL {
	return r.client.GetEndpointURL("/token/session/:create", r.PathParams())
}

func (r *TokenSessionCreateRequest) Do() (TokenSessionCreateResponseBody, error) {
	if r.QueryParams().ConsumerToken == "" {
		r.QueryParams().ConsumerToken = r.client.ConsumerToken()
	}

	if r.QueryParams().EmployeeToken == "" {
		r.QueryParams().EmployeeToken = r.client.EmployeeToken()
	}

	if r.QueryParams().ExpirationDate.IsZero() {
		r.QueryParams().ExpirationDate = Date{time.Now().AddDate(0, 0, 1)}
	}

	// fetch a new token if it isn't set already
	// if r.client.token == "" {
	// 	var err error
	// 	r.client.token, err = r.client.NewToken()
	// 	if err != nil {
	// 		return *r.NewResponseBody(), err
	// 	}
	// }

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
