package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/odata"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewJournalsGet() JournalsGet {
	r := JournalsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalsGet struct {
	client      *Client
	queryParams *JournalsGetQueryParams
	pathParams  *JournalsGetPathParams
	method      string
	headers     http.Header
	requestBody JournalsGetBody
}

func (r JournalsGet) NewQueryParams() *JournalsGetQueryParams {
	selectFields, _ := utils.Fields(&Journal{})
	return &JournalsGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type JournalsGetQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p JournalsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalsGet) QueryParams() *JournalsGetQueryParams {
	return r.queryParams
}

func (r JournalsGet) NewPathParams() *JournalsGetPathParams {
	return &JournalsGetPathParams{}
}

type JournalsGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
}

func (p *JournalsGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
	}
}

func (r *JournalsGet) PathParams() *JournalsGetPathParams {
	return r.pathParams
}

func (r *JournalsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalsGet) SetMethod(method string) {
	r.method = method
}

func (r *JournalsGet) Method() string {
	return r.method
}

func (r JournalsGet) NewRequestBody() JournalsGetBody {
	return JournalsGetBody{}
}

type JournalsGetBody struct {
}

func (r *JournalsGet) RequestBody() *JournalsGetBody {
	return nil
}

func (r *JournalsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalsGet) SetRequestBody(body JournalsGetBody) {
	r.requestBody = body
}

func (r *JournalsGet) NewResponseBody() *JournalsGetResponseBody {
	return &JournalsGetResponseBody{}
}

type JournalsGetResponseBody struct {
	OdataContext string   `json:"@odata.context"`
	Value        Journals `json:"value"`
}

func (r *JournalsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/journals", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *JournalsGet) Do() (JournalsGetResponseBody, error) {
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
