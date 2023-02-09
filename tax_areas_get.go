package central

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewTaxAreasGet() TaxAreasGet {
	r := TaxAreasGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxAreasGet struct {
	client      *Client
	queryParams *TaxAreasGetQueryParams
	pathParams  *TaxAreasGetPathParams
	method      string
	headers     http.Header
	requestBody TaxAreasGetBody
}

func (r TaxAreasGet) NewQueryParams() *TaxAreasGetQueryParams {
	return &TaxAreasGetQueryParams{}
}

type TaxAreasGetQueryParams struct {
}

func (p TaxAreasGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxAreasGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r TaxAreasGet) NewPathParams() *TaxAreasGetPathParams {
	return &TaxAreasGetPathParams{}
}

type TaxAreasGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *TaxAreasGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *TaxAreasGet) PathParams() *TaxAreasGetPathParams {
	return r.pathParams
}

func (r *TaxAreasGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxAreasGet) SetMethod(method string) {
	r.method = method
}

func (r *TaxAreasGet) Method() string {
	return r.method
}

func (r *TaxAreasGet) Headers() http.Header {
	return r.headers
}

func (r TaxAreasGet) NewRequestBody() TaxAreasGetBody {
	return TaxAreasGetBody{}
}

type TaxAreasGetBody struct {
}

func (r *TaxAreasGet) RequestBody() *TaxAreasGetBody {
	return nil
}

func (r *TaxAreasGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxAreasGet) SetRequestBody(body TaxAreasGetBody) {
	r.requestBody = body
}

func (r *TaxAreasGet) NewResponseBody() *TaxAreasGetResponseBody {
	return &TaxAreasGetResponseBody{}
}

type TaxAreasGetResponseBody struct {
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

func (r *TaxAreasGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/taxAreas", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *TaxAreasGet) Do() (TaxAreasGetResponseBody, error) {
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
