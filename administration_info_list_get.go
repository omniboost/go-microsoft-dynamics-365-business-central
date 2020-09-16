package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewAdministrationInfoListGetRequest() AdministrationInfoListGetRequest {
	r := AdministrationInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewAdministrationInfoListGetQueryParams()
	r.pathParams = r.NewAdministrationInfoListGetPathParams()
	r.requestBody = r.NewAdministrationInfoListGetRequestBody()
	return r
}

type AdministrationInfoListGetRequest struct {
	client      *Client
	queryParams *AdministrationInfoListGetQueryParams
	pathParams  *AdministrationInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody AdministrationInfoListGetRequestBody
}

func (r AdministrationInfoListGetRequest) NewAdministrationInfoListGetQueryParams() *AdministrationInfoListGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &AdministrationInfoListGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type AdministrationInfoListGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p AdministrationInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AdministrationInfoListGetRequest) QueryParams() *AdministrationInfoListGetQueryParams {
	return r.queryParams
}

func (r AdministrationInfoListGetRequest) NewAdministrationInfoListGetPathParams() *AdministrationInfoListGetPathParams {
	return &AdministrationInfoListGetPathParams{}
}

type AdministrationInfoListGetPathParams struct {
}

func (p *AdministrationInfoListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AdministrationInfoListGetRequest) PathParams() *AdministrationInfoListGetPathParams {
	return r.pathParams
}

func (r *AdministrationInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AdministrationInfoListGetRequest) Method() string {
	return r.method
}

func (r AdministrationInfoListGetRequest) NewAdministrationInfoListGetRequestBody() AdministrationInfoListGetRequestBody {
	return AdministrationInfoListGetRequestBody{}
}

type AdministrationInfoListGetRequestBody struct{}

func (r *AdministrationInfoListGetRequest) RequestBody() *AdministrationInfoListGetRequestBody {
	return &r.requestBody
}

func (r *AdministrationInfoListGetRequest) SetRequestBody(body AdministrationInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *AdministrationInfoListGetRequest) NewResponseBody() *AdministrationInfoListGetResponseBody {
	return &AdministrationInfoListGetResponseBody{}
}

type AdministrationInfoListGetResponseBody []struct {
	AdministrationID    string   `json:"administrationId"`
	BackupDate          string   `json:"backupDate"`
	Code                string   `json:"code"`
	Description         string   `json:"description"`
	GroupID             int      `json:"groupId"`
	OnlineState         int      `json:"onlineState"`
	PreviousOnlineState int      `json:"previousOnlineState"`
	ReportPath          string   `json:"reportPath"`
	ShortName           string   `json:"shortName"`
	Users               []string `json:"users"`
	Version             int      `json:"version"`
}

func (r *AdministrationInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/AdministrationInfoList", r.PathParams())
}

func (r *AdministrationInfoListGetRequest) Do() (AdministrationInfoListGetResponseBody, error) {
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
