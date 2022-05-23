package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewSalesOrderBasicV2Post() SalesOrderBasicV2Post {
	r := SalesOrderBasicV2Post{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesOrderBasicV2Post struct {
	client      *Client
	queryParams *SalesOrderBasicV2PostQueryParams
	pathParams  *SalesOrderBasicV2PostPathParams
	method      string
	headers     http.Header
	requestBody SalesOrderBasicV2PostBody
}

func (r SalesOrderBasicV2Post) NewQueryParams() *SalesOrderBasicV2PostQueryParams {
	return &SalesOrderBasicV2PostQueryParams{}
}

type SalesOrderBasicV2PostQueryParams struct{}

func (p SalesOrderBasicV2PostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesOrderBasicV2Post) QueryParams() *SalesOrderBasicV2PostQueryParams {
	return r.queryParams
}

func (r SalesOrderBasicV2Post) NewPathParams() *SalesOrderBasicV2PostPathParams {
	return &SalesOrderBasicV2PostPathParams{}
}

type SalesOrderBasicV2PostPathParams struct {
}

func (p *SalesOrderBasicV2PostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesOrderBasicV2Post) PathParams() *SalesOrderBasicV2PostPathParams {
	return r.pathParams
}

func (r *SalesOrderBasicV2Post) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesOrderBasicV2Post) SetMethod(method string) {
	r.method = method
}

func (r *SalesOrderBasicV2Post) Method() string {
	return r.method
}

func (r SalesOrderBasicV2Post) NewRequestBody() SalesOrderBasicV2PostBody {
	return SalesOrderBasicV2PostBody{}
}

type SalesOrderBasicV2PostBody struct {
	Lines []struct {
		BranchNumber    ValueString `json:"branchNumber"`
		Operation       string      `json:"operation"`
		LineNumber      ValueInt    `json:"lineNbr"`
		InventoryID     ValueString `json:"inventoryId"`
		InventoryNumber ValueString `json:"inventoryId"`
		Warehouse       ValueString `json:"warehouse"`
		UOM             ValueString `json:"uom"`
		Quantity        ValueInt    `json:"quantity"`
		UnitCost        ValueInt    `json:"unitCost"`
		UnitPrice       ValueInt    `json:"unitPrice"`
		DiscountPercent ValueInt    `json:"discountPercent"`
		DiscountAmount  ValueInt    `json:"discountAmount"`
		DiscountCode    ValueString `json:"discountCode"`
		ManualDiscount  ValueBool   `json:"manualDiscount"`
		DiscUnitPrice   ValueInt    `json:"discUnitPrice"`
		LineDescription ValueString `json:"lineDescription"`
		Note            ValueString `json:"note"`
	} `json:"lines"`
	OrderType             ValueString `json:"orderType"`
	OrderNumber           ValueString `json:"orderNumber"`
	Hold                  ValueBool   `json:"hold"`
	Date                  ValueTime   `json:"date"`
	RequestOn             ValueTime   `json:"requestOn"`
	CustomerOrder         ValueString `json:"customerOrder"`
	CustomerRefNo         ValueString `json:"customerRefNo"`
	Customer              ValueString `json:"customer"`
	Location              ValueString `json:"location"`
	ContactID             ValueInt    `json:"contactId"`
	GLN                   ValueString `json:"gln"`
	VATRegistrationID     ValueString `json:"vatRegistrationId"`
	Currency              ValueString `json:"currency"`
	Description           ValueString `json:"description"`
	RecalculationShipment ValueBool   `json:"recalculationShipment"`
	BranchNumber          ValueString `json:"branchNumber"`
	Note                  ValueString `json:"note"`
	OverrideNumberSeries  ValueBool   `json:"overrideNumberSeries"`
}

// func (r *SalesOrderBasicV2Post) MarshalJSON() ([]byte, error) {
// 	return omitempty.MarshalJSON(r)
// }

func (r *SalesOrderBasicV2Post) RequestBody() *SalesOrderBasicV2PostBody {
	return &r.requestBody
}

func (r *SalesOrderBasicV2Post) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SalesOrderBasicV2Post) SetRequestBody(body SalesOrderBasicV2PostBody) {
	r.requestBody = body
}

func (r *SalesOrderBasicV2Post) NewResponseBody() *SalesOrderBasicV2PostResponseBody {
	return &SalesOrderBasicV2PostResponseBody{}
}

type SalesOrderBasicV2PostResponseBody JournalTransactions

func (r *SalesOrderBasicV2Post) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v2/salesorderbasic", r.PathParams())
	return &u
}

func (r *SalesOrderBasicV2Post) Do() (SalesOrderBasicV2PostResponseBody, error) {
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
