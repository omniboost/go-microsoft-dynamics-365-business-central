package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewDimensionSetLinesPost() DimensionSetLinesPost {
	r := DimensionSetLinesPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DimensionSetLinesPost struct {
	client      *Client
	queryParams *DimensionSetLinesPostQueryParams
	pathParams  *DimensionSetLinesPostPathParams
	method      string
	headers     http.Header
	requestBody DimensionSetLinesPostBody
}

func (r DimensionSetLinesPost) NewQueryParams() *DimensionSetLinesPostQueryParams {
	return &DimensionSetLinesPostQueryParams{}
}

type DimensionSetLinesPostQueryParams struct {
}

func (p DimensionSetLinesPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DimensionSetLinesPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r DimensionSetLinesPost) NewPathParams() *DimensionSetLinesPostPathParams {
	return &DimensionSetLinesPostPathParams{}
}

type DimensionSetLinesPostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	Object          string `schema:"object"`
	ObjectID        string `schema:"objectID"`
}

func (p *DimensionSetLinesPostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"object":          p.Object,
		"objectID":        p.ObjectID,
	}
}

func (r *DimensionSetLinesPost) PathParams() *DimensionSetLinesPostPathParams {
	return r.pathParams
}

func (r *DimensionSetLinesPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DimensionSetLinesPost) SetMethod(method string) {
	r.method = method
}

func (r *DimensionSetLinesPost) Method() string {
	return r.method
}

func (r *DimensionSetLinesPost) Headers() http.Header {
	return r.headers
}

func (r DimensionSetLinesPost) NewRequestBody() DimensionSetLinesPostBody {
	return DimensionSetLinesPostBody{}
}

type DimensionSetLinesPostBody struct {
	DimensionSetLine
}

func (r *DimensionSetLinesPost) RequestBody() *DimensionSetLinesPostBody {
	return &r.requestBody
}

func (r *DimensionSetLinesPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *DimensionSetLinesPost) SetRequestBody(body DimensionSetLinesPostBody) {
	r.requestBody = body
}

func (r *DimensionSetLinesPost) NewResponseBody() *DimensionSetLinesPostResponseBody {
	return &DimensionSetLinesPostResponseBody{}
}

type DimensionSetLinesPostResponseBody struct {
	OdataContext string            `json:"@odata.context"`
	Value        DimensionSetLines `json:"value"`
}

func (r *DimensionSetLinesPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/{{.object}}({{.objectID}})/dimensionSetLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *DimensionSetLinesPost) Do() (DimensionSetLinesPostResponseBody, error) {
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
