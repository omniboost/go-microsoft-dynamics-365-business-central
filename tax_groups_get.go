package central

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewTaxGroupsGet() TaxGroupsGet {
	r := TaxGroupsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxGroupsGet struct {
	client      *Client
	queryParams *TaxGroupsGetQueryParams
	pathParams  *TaxGroupsGetPathParams
	method      string
	headers     http.Header
	requestBody TaxGroupsGetBody
}

func (r TaxGroupsGet) NewQueryParams() *TaxGroupsGetQueryParams {
	return &TaxGroupsGetQueryParams{}
}

type TaxGroupsGetQueryParams struct {
}

func (p TaxGroupsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxGroupsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r TaxGroupsGet) NewPathParams() *TaxGroupsGetPathParams {
	return &TaxGroupsGetPathParams{}
}

type TaxGroupsGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *TaxGroupsGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *TaxGroupsGet) PathParams() *TaxGroupsGetPathParams {
	return r.pathParams
}

func (r *TaxGroupsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxGroupsGet) SetMethod(method string) {
	r.method = method
}

func (r *TaxGroupsGet) Method() string {
	return r.method
}

func (r *TaxGroupsGet) Headers() http.Header {
	return r.headers
}

func (r TaxGroupsGet) NewRequestBody() TaxGroupsGetBody {
	return TaxGroupsGetBody{}
}

type TaxGroupsGetBody struct {
}

func (r *TaxGroupsGet) RequestBody() *TaxGroupsGetBody {
	return nil
}

func (r *TaxGroupsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxGroupsGet) SetRequestBody(body TaxGroupsGetBody) {
	r.requestBody = body
}

func (r *TaxGroupsGet) NewResponseBody() *TaxGroupsGetResponseBody {
	return &TaxGroupsGetResponseBody{}
}

type TaxGroupsGetResponseBody struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		OdataEtag            string    `json:"@odata.etag"`
		ID                   string    `json:"id"`
		Code                 string    `json:"code"`
		DisplayName          string    `json:"displayName"`
		TaxType              string    `json:"taxType"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"value"`
}

func (r *TaxGroupsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/taxGroups", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *TaxGroupsGet) Do() (TaxGroupsGetResponseBody, error) {
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
