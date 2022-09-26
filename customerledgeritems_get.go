package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/odata"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewCustomerledgeritemsGet() CustomerledgeritemsGet {
	r := CustomerledgeritemsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerledgeritemsGet struct {
	client      *Client
	queryParams *CustomerledgeritemsGetQueryParams
	pathParams  *CustomerledgeritemsGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerledgeritemsGetBody
}

func (r CustomerledgeritemsGet) NewQueryParams() *CustomerledgeritemsGetQueryParams {
	selectFields, _ := utils.Fields(&Customer{})
	expandFields := []string{}
	return &CustomerledgeritemsGetQueryParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CustomerledgeritemsGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerledgeritemsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerledgeritemsGet) QueryParams() *CustomerledgeritemsGetQueryParams {
	return r.queryParams
}

func (r CustomerledgeritemsGet) NewPathParams() *CustomerledgeritemsGetPathParams {
	return &CustomerledgeritemsGetPathParams{}
}

type CustomerledgeritemsGetPathParams struct {
}

func (p *CustomerledgeritemsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerledgeritemsGet) PathParams() *CustomerledgeritemsGetPathParams {
	return r.pathParams
}

func (r *CustomerledgeritemsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerledgeritemsGet) SetMethod(method string) {
	r.method = method
}

func (r *CustomerledgeritemsGet) Method() string {
	return r.method
}

func (r CustomerledgeritemsGet) NewRequestBody() CustomerledgeritemsGetBody {
	return CustomerledgeritemsGetBody{}
}

type CustomerledgeritemsGetBody struct {
}

func (r *CustomerledgeritemsGet) RequestBody() *CustomerledgeritemsGetBody {
	return nil
}

func (r *CustomerledgeritemsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerledgeritemsGet) SetRequestBody(body CustomerledgeritemsGetBody) {
	r.requestBody = body
}

func (r *CustomerledgeritemsGet) NewResponseBody() *CustomerledgeritemsGetResponseBody {
	return &CustomerledgeritemsGetResponseBody{}
}

type CustomerledgeritemsGetResponseBody struct {
	Meta Meta                `json:"Meta"`
	Data CustomerLedgerItems `json:"Data"`
}

func (r *CustomerledgeritemsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/customerledgeritems", r.PathParams())
	return &u
}

func (r *CustomerledgeritemsGet) Do() (CustomerledgeritemsGetResponseBody, error) {
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
