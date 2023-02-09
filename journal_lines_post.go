package central

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewJournalLinesPost() JournalLinesPost {
	r := JournalLinesPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalLinesPost struct {
	client      *Client
	queryParams *JournalLinesPostQueryParams
	pathParams  *JournalLinesPostPathParams
	method      string
	headers     http.Header
	requestBody JournalLinesPostBody
}

func (r JournalLinesPost) NewQueryParams() *JournalLinesPostQueryParams {
	return &JournalLinesPostQueryParams{}
}

type JournalLinesPostQueryParams struct {
}

func (p JournalLinesPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalLinesPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r JournalLinesPost) NewPathParams() *JournalLinesPostPathParams {
	return &JournalLinesPostPathParams{}
}

type JournalLinesPostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	JournalID       string `schema:"journalID"`
}

func (p *JournalLinesPostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"journalID":       p.JournalID,
	}
}

func (r *JournalLinesPost) PathParams() *JournalLinesPostPathParams {
	return r.pathParams
}

func (r *JournalLinesPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalLinesPost) SetMethod(method string) {
	r.method = method
}

func (r *JournalLinesPost) Method() string {
	return r.method
}

func (r *JournalLinesPost) Headers() http.Header {
	return r.headers
}

func (r JournalLinesPost) NewRequestBody() JournalLinesPostBody {
	return JournalLinesPostBody{}
}

type JournalLinesPostBody struct {
	JournalLine
}

func (r *JournalLinesPost) RequestBody() *JournalLinesPostBody {
	return &r.requestBody
}

func (r *JournalLinesPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *JournalLinesPost) SetRequestBody(body JournalLinesPostBody) {
	r.requestBody = body
}

func (r *JournalLinesPost) NewResponseBody() *JournalLinesPostResponseBody {
	return &JournalLinesPostResponseBody{}
}

type JournalLinesPostResponseBody struct {
	OdataContext string       `json:"@odata.context"`
	Value        JournalLines `json:"value"`
}

func (r *JournalLinesPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/journals({{.journalID}})/journalLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *JournalLinesPost) Do() (JournalLinesPostResponseBody, error) {
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
