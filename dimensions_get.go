package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewDimensionsGet() DimensionsGet {
	r := DimensionsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DimensionsGet struct {
	client      *Client
	queryParams *DimensionsGetQueryParams
	pathParams  *DimensionsGetPathParams
	method      string
	headers     http.Header
	requestBody DimensionsGetBody
}

func (r DimensionsGet) NewQueryParams() *DimensionsGetQueryParams {
	selectFields, _ := utils.Fields(&Dimension{})
	return &DimensionsGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type DimensionsGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p DimensionsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DimensionsGet) QueryParams() *DimensionsGetQueryParams {
	return r.queryParams
}

func (r DimensionsGet) NewPathParams() *DimensionsGetPathParams {
	return &DimensionsGetPathParams{}
}

type DimensionsGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *DimensionsGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *DimensionsGet) PathParams() *DimensionsGetPathParams {
	return r.pathParams
}

func (r *DimensionsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DimensionsGet) SetMethod(method string) {
	r.method = method
}

func (r *DimensionsGet) Method() string {
	return r.method
}

func (r *DimensionsGet) Headers() http.Header {
	return r.headers
}

func (r DimensionsGet) NewRequestBody() DimensionsGetBody {
	return DimensionsGetBody{}
}

type DimensionsGetBody struct {
}

func (r *DimensionsGet) RequestBody() *DimensionsGetBody {
	return nil
}

func (r *DimensionsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *DimensionsGet) SetRequestBody(body DimensionsGetBody) {
	r.requestBody = body
}

func (r *DimensionsGet) NewResponseBody() *DimensionsGetResponseBody {
	return &DimensionsGetResponseBody{}
}

type DimensionsGetResponseBody struct {
	OdataContext string     `json:"@odata.context"`
	Value        Dimensions `json:"value"`
}

func (r *DimensionsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/dimensions", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *DimensionsGet) Do() (DimensionsGetResponseBody, error) {
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
