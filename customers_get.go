package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewCustomersGet() CustomersGet {
	r := CustomersGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomersGet struct {
	client      *Client
	queryParams *CustomersGetQueryParams
	pathParams  *CustomersGetPathParams
	method      string
	headers     http.Header
	requestBody CustomersGetBody
}

func (r CustomersGet) NewQueryParams() *CustomersGetQueryParams {
	selectFields, _ := utils.Fields(&Customer{})
	return &CustomersGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CustomersGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomersGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomersGet) QueryParams() *CustomersGetQueryParams {
	return r.queryParams
}

func (r CustomersGet) NewPathParams() *CustomersGetPathParams {
	return &CustomersGetPathParams{}
}

type CustomersGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *CustomersGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *CustomersGet) PathParams() *CustomersGetPathParams {
	return r.pathParams
}

func (r *CustomersGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomersGet) SetMethod(method string) {
	r.method = method
}

func (r *CustomersGet) Method() string {
	return r.method
}

func (r *CustomersGet) Headers() http.Header {
	return r.headers
}

func (r CustomersGet) NewRequestBody() CustomersGetBody {
	return CustomersGetBody{}
}

type CustomersGetBody struct {
}

func (r *CustomersGet) RequestBody() *CustomersGetBody {
	return nil
}

func (r *CustomersGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomersGet) SetRequestBody(body CustomersGetBody) {
	r.requestBody = body
}

func (r *CustomersGet) NewResponseBody() *CustomersGetResponseBody {
	return &CustomersGetResponseBody{}
}

type CustomersGetResponseBody struct {
	OdataContext string `json:"@odata.context"`
	Value        Customers `json:"value"`
}

func (r *CustomersGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/customers", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CustomersGet) Do() (CustomersGetResponseBody, error) {
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
