package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-vismaonline/odata"
	"github.com/omniboost/go-vismaonline/utils"
)

func (c *Client) NewArticlesGet() ArticlesGet {
	r := ArticlesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ArticlesGet struct {
	client      *Client
	queryParams *ArticlesGetQueryParams
	pathParams  *ArticlesGetPathParams
	method      string
	headers     http.Header
	requestBody ArticlesGetBody
}

func (r ArticlesGet) NewQueryParams() *ArticlesGetQueryParams {
	selectFields, _ := utils.Fields(&Article{})
	expandFields := []string{}
	return &ArticlesGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type ArticlesGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p ArticlesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ArticlesGet) QueryParams() *ArticlesGetQueryParams {
	return r.queryParams
}

func (r ArticlesGet) NewPathParams() *ArticlesGetPathParams {
	return &ArticlesGetPathParams{}
}

type ArticlesGetPathParams struct {
}

func (p *ArticlesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ArticlesGet) PathParams() *ArticlesGetPathParams {
	return r.pathParams
}

func (r *ArticlesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ArticlesGet) SetMethod(method string) {
	r.method = method
}

func (r *ArticlesGet) Method() string {
	return r.method
}

func (r ArticlesGet) NewRequestBody() ArticlesGetBody {
	return ArticlesGetBody{}
}

type ArticlesGetBody struct {
}

func (r *ArticlesGet) RequestBody() *ArticlesGetBody {
	return nil
}

func (r *ArticlesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ArticlesGet) SetRequestBody(body ArticlesGetBody) {
	r.requestBody = body
}

func (r *ArticlesGet) NewResponseBody() *ArticlesGetResponseBody {
	return &ArticlesGetResponseBody{}
}

type ArticlesGetResponseBody struct {
	Meta Meta     `json:"Meta"`
	Data Articles `json:"Data"`
}

func (r *ArticlesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/articles", r.PathParams())
	return &u
}

func (r *ArticlesGet) Do() (ArticlesGetResponseBody, error) {
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
