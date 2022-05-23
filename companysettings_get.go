package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCompanysettingsGet() CompanysettingsGet {
	r := CompanysettingsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompanysettingsGet struct {
	client      *Client
	queryParams *CompanysettingsGetQueryParams
	pathParams  *CompanysettingsGetPathParams
	method      string
	headers     http.Header
	requestBody CompanysettingsGetBody
}

func (r CompanysettingsGet) NewQueryParams() *CompanysettingsGetQueryParams {
	return &CompanysettingsGetQueryParams{}
}

type CompanysettingsGetQueryParams struct {
}

func (p CompanysettingsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CompanysettingsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r CompanysettingsGet) NewPathParams() *CompanysettingsGetPathParams {
	return &CompanysettingsGetPathParams{}
}

type CompanysettingsGetPathParams struct {
}

func (p *CompanysettingsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompanysettingsGet) PathParams() *CompanysettingsGetPathParams {
	return r.pathParams
}

func (r *CompanysettingsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CompanysettingsGet) SetMethod(method string) {
	r.method = method
}

func (r *CompanysettingsGet) Method() string {
	return r.method
}

func (r CompanysettingsGet) NewRequestBody() CompanysettingsGetBody {
	return CompanysettingsGetBody{}
}

type CompanysettingsGetBody struct {
}

func (r *CompanysettingsGet) RequestBody() *CompanysettingsGetBody {
	return nil
}

func (r *CompanysettingsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CompanysettingsGet) SetRequestBody(body CompanysettingsGetBody) {
	r.requestBody = body
}

func (r *CompanysettingsGet) NewResponseBody() *CompanysettingsGetResponseBody {
	return &CompanysettingsGetResponseBody{}
}

type CompanysettingsGetResponseBody Companysettings

func (r *CompanysettingsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/companysettings", r.PathParams())
	return &u
}

func (r *CompanysettingsGet) Do() (CompanysettingsGetResponseBody, error) {
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
