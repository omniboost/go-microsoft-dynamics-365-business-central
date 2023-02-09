package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewDimensionSetLinesGet() DimensionSetLinesGet {
	r := DimensionSetLinesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DimensionSetLinesGet struct {
	client      *Client
	queryParams *DimensionSetLinesGetQueryParams
	pathParams  *DimensionSetLinesGetPathParams
	method      string
	headers     http.Header
	requestBody DimensionSetLinesGetBody
}

func (r DimensionSetLinesGet) NewQueryParams() *DimensionSetLinesGetQueryParams {
	return &DimensionSetLinesGetQueryParams{}
}

type DimensionSetLinesGetQueryParams struct {
}

func (p DimensionSetLinesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DimensionSetLinesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r DimensionSetLinesGet) NewPathParams() *DimensionSetLinesGetPathParams {
	return &DimensionSetLinesGetPathParams{}
}

type DimensionSetLinesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	Object          string `schema:"object"`
	ObjectID        string `schema:"objectID"`
}

func (p *DimensionSetLinesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"object":          p.Object,
		"objectID":        p.ObjectID,
	}
}

func (r *DimensionSetLinesGet) PathParams() *DimensionSetLinesGetPathParams {
	return r.pathParams
}

func (r *DimensionSetLinesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DimensionSetLinesGet) SetMethod(method string) {
	r.method = method
}

func (r *DimensionSetLinesGet) Method() string {
	return r.method
}

func (r *DimensionSetLinesGet) Headers() http.Header {
	return r.headers
}

func (r DimensionSetLinesGet) NewRequestBody() DimensionSetLinesGetBody {
	return DimensionSetLinesGetBody{}
}

type DimensionSetLinesGetBody struct {
}

func (r *DimensionSetLinesGet) RequestBody() *DimensionSetLinesGetBody {
	return nil
}

func (r *DimensionSetLinesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *DimensionSetLinesGet) SetRequestBody(body DimensionSetLinesGetBody) {
	r.requestBody = body
}

func (r *DimensionSetLinesGet) NewResponseBody() *DimensionSetLinesGetResponseBody {
	return &DimensionSetLinesGetResponseBody{}
}

type DimensionSetLinesGetResponseBody struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		OdataEtag        string `json:"@odata.etag"`
		ID               string `json:"id"`
		Code             string `json:"code"`
		ParentID         string `json:"parentId"`
		ParentType       string `json:"parentType"`
		DisplayName      string `json:"displayName"`
		ValueID          string `json:"valueId"`
		ValueCode        string `json:"valueCode"`
		ValueDisplayName string `json:"valueDisplayName"`
	} `json:"value"`
}

func (r *DimensionSetLinesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/{{.object}}({{.objectID}})/dimensionSetLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *DimensionSetLinesGet) Do() (DimensionSetLinesGetResponseBody, error) {
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
