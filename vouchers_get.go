package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-vismaonline/odata"
	"github.com/omniboost/go-vismaonline/utils"
)

func (c *Client) NewVouchersGet() VouchersGet {
	r := VouchersGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VouchersGet struct {
	client      *Client
	queryParams *VouchersGetQueryParams
	pathParams  *VouchersGetPathParams
	method      string
	headers     http.Header
	requestBody VouchersGetBody
}

func (r VouchersGet) NewQueryParams() *VouchersGetQueryParams {
	selectFields, _ := utils.Fields(&Customer{})
	expandFields := []string{}
	return &VouchersGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type VouchersGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p VouchersGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VouchersGet) QueryParams() *VouchersGetQueryParams {
	return r.queryParams
}

func (r VouchersGet) NewPathParams() *VouchersGetPathParams {
	return &VouchersGetPathParams{}
}

type VouchersGetPathParams struct {
}

func (p *VouchersGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VouchersGet) PathParams() *VouchersGetPathParams {
	return r.pathParams
}

func (r *VouchersGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VouchersGet) SetMethod(method string) {
	r.method = method
}

func (r *VouchersGet) Method() string {
	return r.method
}

func (r VouchersGet) NewRequestBody() VouchersGetBody {
	return VouchersGetBody{}
}

type VouchersGetBody struct {
}

func (r *VouchersGet) RequestBody() *VouchersGetBody {
	return nil
}

func (r *VouchersGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VouchersGet) SetRequestBody(body VouchersGetBody) {
	r.requestBody = body
}

func (r *VouchersGet) NewResponseBody() *VouchersGetResponseBody {
	return &VouchersGetResponseBody{}
}

type VouchersGetResponseBody struct {
	Meta Meta     `json:"Meta"`
	Data Vouchers `json:"Data"`
}

func (r *VouchersGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/vouchers", r.PathParams())
	return &u
}

func (r *VouchersGet) Do() (VouchersGetResponseBody, error) {
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
