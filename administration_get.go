package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewAdministrationGetRequest() AdministrationGetRequest {
	r := AdministrationGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewAdministrationGetQueryParams()
	r.pathParams = r.NewAdministrationGetPathParams()
	r.requestBody = r.NewAdministrationGetRequestBody()
	return r
}

type AdministrationGetRequest struct {
	client      *Client
	queryParams *AdministrationGetQueryParams
	pathParams  *AdministrationGetPathParams
	method      string
	headers     http.Header
	requestBody AdministrationGetRequestBody
}

func (r AdministrationGetRequest) NewAdministrationGetQueryParams() *AdministrationGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &AdministrationGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type AdministrationGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p AdministrationGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AdministrationGetRequest) QueryParams() *AdministrationGetQueryParams {
	return r.queryParams
}

func (r AdministrationGetRequest) NewAdministrationGetPathParams() *AdministrationGetPathParams {
	return &AdministrationGetPathParams{}
}

type AdministrationGetPathParams struct {
}

func (p *AdministrationGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AdministrationGetRequest) PathParams() *AdministrationGetPathParams {
	return r.pathParams
}

func (r *AdministrationGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AdministrationGetRequest) Method() string {
	return r.method
}

func (r AdministrationGetRequest) NewAdministrationGetRequestBody() AdministrationGetRequestBody {
	return AdministrationGetRequestBody{}
}

type AdministrationGetRequestBody struct{}

func (r *AdministrationGetRequest) RequestBody() *AdministrationGetRequestBody {
	return &r.requestBody
}

func (r *AdministrationGetRequest) SetRequestBody(body AdministrationGetRequestBody) {
	r.requestBody = body
}

func (r *AdministrationGetRequest) NewResponseBody() *AdministrationGetResponseBody {
	return &AdministrationGetResponseBody{}
}

type AdministrationGetResponseBody struct {
	BackupDate          string        `json:"backupDate"`
	Version             int           `json:"version"`
	GroupID             int           `json:"groupId"`
	Code                string        `json:"code"`
	ShortName           string        `json:"shortName"`
	PreviousOnlineState int           `json:"previousOnlineState"`
	OnlineState         int           `json:"onlineState"`
	Description         string        `json:"description"`
	AdministrationID    string        `json:"administrationId"`
	Users               []string      `json:"users"`
	ReportPath          string        `json:"reportPath"`
	Messages            []interface{} `json:"messages"`
	CanChange           bool          `json:"canChange"`
	CannotChangeReason  string        `json:"cannotChangeReason"`
}

func (r *AdministrationGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/Administration/{{.administration_id}}", r.PathParams())
}

func (r *AdministrationGetRequest) Do() (AdministrationGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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
