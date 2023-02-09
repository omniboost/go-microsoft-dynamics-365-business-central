package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewDimensionValuesGet() DimensionValuesGet {
	r := DimensionValuesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DimensionValuesGet struct {
	client      *Client
	queryParams *DimensionValuesGetQueryParams
	pathParams  *DimensionValuesGetPathParams
	method      string
	headers     http.Header
	requestBody DimensionValuesGetBody
}

func (r DimensionValuesGet) NewQueryParams() *DimensionValuesGetQueryParams {
	selectFields, _ := utils.Fields(&Dimension{})
	return &DimensionValuesGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type DimensionValuesGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p DimensionValuesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DimensionValuesGet) QueryParams() *DimensionValuesGetQueryParams {
	return r.queryParams
}

func (r DimensionValuesGet) NewPathParams() *DimensionValuesGetPathParams {
	return &DimensionValuesGetPathParams{}
}

type DimensionValuesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	DimensionID     string `schema:"dimensionID"`
}

func (p *DimensionValuesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"dimensionID":     p.DimensionID,
	}
}

func (r *DimensionValuesGet) PathParams() *DimensionValuesGetPathParams {
	return r.pathParams
}

func (r *DimensionValuesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DimensionValuesGet) SetMethod(method string) {
	r.method = method
}

func (r *DimensionValuesGet) Method() string {
	return r.method
}

func (r *DimensionValuesGet) Headers() http.Header {
	return r.headers
}

func (r DimensionValuesGet) NewRequestBody() DimensionValuesGetBody {
	return DimensionValuesGetBody{}
}

type DimensionValuesGetBody struct {
}

func (r *DimensionValuesGet) RequestBody() *DimensionValuesGetBody {
	return nil
}

func (r *DimensionValuesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *DimensionValuesGet) SetRequestBody(body DimensionValuesGetBody) {
	r.requestBody = body
}

func (r *DimensionValuesGet) NewResponseBody() *DimensionValuesGetResponseBody {
	return &DimensionValuesGetResponseBody{}
}

type DimensionValuesGetResponseBody struct {
	OdataContext string          `json:"@odata.context"`
	Value        DimensionValues `json:"value"`
}

func (r *DimensionValuesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/dimensions({{.dimensionID}})/dimensionValues", r.PathParams())
	return &u
}

func (r *DimensionValuesGet) Do() (DimensionValuesGetResponseBody, error) {
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
