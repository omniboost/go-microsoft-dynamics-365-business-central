package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewVatCodeInfoListGetRequest() VatCodeInfoListGetRequest {
	r := VatCodeInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewVatCodeInfoListGetQueryParams()
	r.pathParams = r.NewVatCodeInfoListGetPathParams()
	r.requestBody = r.NewVatCodeInfoListGetRequestBody()
	return r
}

type VatCodeInfoListGetRequest struct {
	client      *Client
	queryParams *VatCodeInfoListGetQueryParams
	pathParams  *VatCodeInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody VatCodeInfoListGetRequestBody
}

func (r VatCodeInfoListGetRequest) NewVatCodeInfoListGetQueryParams() *VatCodeInfoListGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &VatCodeInfoListGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type VatCodeInfoListGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p VatCodeInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *VatCodeInfoListGetRequest) QueryParams() *VatCodeInfoListGetQueryParams {
	return r.queryParams
}

func (r VatCodeInfoListGetRequest) NewVatCodeInfoListGetPathParams() *VatCodeInfoListGetPathParams {
	return &VatCodeInfoListGetPathParams{}
}

type VatCodeInfoListGetPathParams struct {
}

func (p *VatCodeInfoListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VatCodeInfoListGetRequest) PathParams() *VatCodeInfoListGetPathParams {
	return r.pathParams
}

func (r *VatCodeInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *VatCodeInfoListGetRequest) Method() string {
	return r.method
}

func (r VatCodeInfoListGetRequest) NewVatCodeInfoListGetRequestBody() VatCodeInfoListGetRequestBody {
	return VatCodeInfoListGetRequestBody{}
}

type VatCodeInfoListGetRequestBody struct{}

func (r *VatCodeInfoListGetRequest) RequestBody() *VatCodeInfoListGetRequestBody {
	return &r.requestBody
}

func (r *VatCodeInfoListGetRequest) SetRequestBody(body VatCodeInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *VatCodeInfoListGetRequest) NewResponseBody() *VatCodeInfoListGetResponseBody {
	return &VatCodeInfoListGetResponseBody{}
}

type VatCodeInfoListGetResponseBody []struct {
	Description          string  `json:"description"`
	VatCodeID            int     `json:"vatCodeId"`
	VatPercentage        float64 `json:"vatPercentage"`
	VatPercentagesByDate []struct {
		StartDate     string  `json:"startDate"`
		VatCodeID     int     `json:"vatCodeId"`
		VatPercentage float64 `json:"vatPercentage"`
	} `json:"vatPercentagesByDate"`
}

func (r *VatCodeInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/VatCodeInfoList", r.PathParams())
}

func (r *VatCodeInfoListGetRequest) Do() (VatCodeInfoListGetResponseBody, error) {
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
