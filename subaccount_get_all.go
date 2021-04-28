package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewSubaccountGetAll() SubaccountGetAll {
	r := SubaccountGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SubaccountGetAll struct {
	client      *Client
	queryParams *SubaccountGetAllQueryParams
	pathParams  *SubaccountGetAllPathParams
	method      string
	headers     http.Header
	requestBody SubaccountGetAllBody
}

func (r SubaccountGetAll) NewQueryParams() *SubaccountGetAllQueryParams {
	return &SubaccountGetAllQueryParams{}
}

type SubaccountGetAllQueryParams struct {
}

func (p SubaccountGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SubaccountGetAll) QueryParams() QueryParams {
	return r.queryParams
}

func (r SubaccountGetAll) NewPathParams() *SubaccountGetAllPathParams {
	return &SubaccountGetAllPathParams{}
}

type SubaccountGetAllPathParams struct {
}

func (p *SubaccountGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SubaccountGetAll) PathParams() *SubaccountGetAllPathParams {
	return r.pathParams
}

func (r *SubaccountGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SubaccountGetAll) SetMethod(method string) {
	r.method = method
}

func (r *SubaccountGetAll) Method() string {
	return r.method
}

func (r SubaccountGetAll) NewRequestBody() SubaccountGetAllBody {
	return SubaccountGetAllBody{}
}

type SubaccountGetAllBody struct {
}

func (r *SubaccountGetAll) RequestBody() *SubaccountGetAllBody {
	return nil
}

func (r *SubaccountGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *SubaccountGetAll) SetRequestBody(body SubaccountGetAllBody) {
	r.requestBody = body
}

func (r *SubaccountGetAll) NewResponseBody() *SubaccountGetAllResponseBody {
	return &SubaccountGetAllResponseBody{}
}

type SubaccountGetAllResponseBody []struct {
	SubaccountNumber     string `json:"subaccountNumber"`
	SubaccountID         int    `json:"subaccountId"`
	Description          string `json:"description"`
	LastModifiedDateTime Time   `json:"lastModifiedDateTime"`
	Active               bool   `json:"active"`
	Segments             []struct {
		SegmentID               int    `json:"segmentId"`
		SegmentDescription      string `json:"segmentDescription"`
		SegmentValue            string `json:"segmentValue"`
		SegmentValueDescription string `json:"segmentValueDescription"`
	} `json:"segments"`
	ErrorInfo string   `json:"errorInfo"`
	Metadata  Metadata `json:"metadata"`
}

func (r *SubaccountGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("controller/api/v1/subaccount", r.PathParams())
	return &u
}

func (r *SubaccountGetAll) Do() (SubaccountGetAllResponseBody, error) {
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
