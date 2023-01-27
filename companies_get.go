package central

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewCompaniesGet() CompaniesGet {
	r := CompaniesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompaniesGet struct {
	client      *Client
	queryParams *CompaniesGetQueryParams
	pathParams  *CompaniesGetPathParams
	method      string
	headers     http.Header
	requestBody CompaniesGetBody
}

func (r CompaniesGet) NewQueryParams() *CompaniesGetQueryParams {
	return &CompaniesGetQueryParams{}
}

type CompaniesGetQueryParams struct {
}

func (p CompaniesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CompaniesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r CompaniesGet) NewPathParams() *CompaniesGetPathParams {
	return &CompaniesGetPathParams{}
}

type CompaniesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
}

func (p *CompaniesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
	}
}

func (r *CompaniesGet) PathParams() *CompaniesGetPathParams {
	return r.pathParams
}

func (r *CompaniesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CompaniesGet) SetMethod(method string) {
	r.method = method
}

func (r *CompaniesGet) Method() string {
	return r.method
}

func (r CompaniesGet) NewRequestBody() CompaniesGetBody {
	return CompaniesGetBody{}
}

type CompaniesGetBody struct {
}

func (r *CompaniesGet) RequestBody() *CompaniesGetBody {
	return nil
}

func (r *CompaniesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CompaniesGet) SetRequestBody(body CompaniesGetBody) {
	r.requestBody = body
}

func (r *CompaniesGet) NewResponseBody() *CompaniesGetResponseBody {
	return &CompaniesGetResponseBody{}
}

type CompaniesGetResponseBody struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		ID                string    `json:"id"`
		SystemVersion     string    `json:"systemVersion"`
		Timestamp         int       `json:"timestamp"`
		Name              string    `json:"name"`
		DisplayName       string    `json:"displayName"`
		BusinessProfileID string    `json:"businessProfileId"`
		SystemCreatedAt   time.Time `json:"systemCreatedAt"`
		SystemCreatedBy   string    `json:"systemCreatedBy"`
		SystemModifiedAt  time.Time `json:"systemModifiedAt"`
		SystemModifiedBy  string    `json:"systemModifiedBy"`
	} `json:"value"`
}

func (r *CompaniesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CompaniesGet) Do() (CompaniesGetResponseBody, error) {
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
