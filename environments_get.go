package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewEnvironmentsGet() EnvironmentsGet {
	r := EnvironmentsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EnvironmentsGet struct {
	client      *Client
	queryParams *EnvironmentsGetQueryParams
	pathParams  *EnvironmentsGetPathParams
	method      string
	headers     http.Header
	requestBody EnvironmentsGetBody
}

func (r EnvironmentsGet) NewQueryParams() *EnvironmentsGetQueryParams {
	return &EnvironmentsGetQueryParams{}
}

type EnvironmentsGetQueryParams struct {
}

func (p EnvironmentsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EnvironmentsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r EnvironmentsGet) NewPathParams() *EnvironmentsGetPathParams {
	return &EnvironmentsGetPathParams{}
}

type EnvironmentsGetPathParams struct {
}

func (p *EnvironmentsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EnvironmentsGet) PathParams() *EnvironmentsGetPathParams {
	return r.pathParams
}

func (r *EnvironmentsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EnvironmentsGet) SetMethod(method string) {
	r.method = method
}

func (r *EnvironmentsGet) Method() string {
	return r.method
}

func (r *EnvironmentsGet) Headers() http.Header {
	return r.headers
}

func (r EnvironmentsGet) NewRequestBody() EnvironmentsGetBody {
	return EnvironmentsGetBody{}
}

type EnvironmentsGetBody struct {
}

func (r *EnvironmentsGet) RequestBody() *EnvironmentsGetBody {
	return nil
}

func (r *EnvironmentsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *EnvironmentsGet) SetRequestBody(body EnvironmentsGetBody) {
	r.requestBody = body
}

func (r *EnvironmentsGet) NewResponseBody() *EnvironmentsGetResponseBody {
	return &EnvironmentsGetResponseBody{}
}

type EnvironmentsGetResponseBody struct {
	Value []struct {
		AadTenantID       string `json:"aadTenantId"`
		ApplicationFamily string `json:"applicationFamily"`
		Type              string `json:"type"`
		Name              string `json:"name"`
		CountryCode       string `json:"countryCode"`
		WebServiceURL     string `json:"webServiceUrl"`
		WebClientLoginURL string `json:"webClientLoginUrl"`
	} `json:"value"`
}

func (r *EnvironmentsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/environments/v1.1", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *EnvironmentsGet) Do() (EnvironmentsGetResponseBody, error) {
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
