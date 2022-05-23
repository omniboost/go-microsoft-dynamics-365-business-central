package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewJournalTransactionV2GetAll() JournalTransactionV2GetAll {
	r := JournalTransactionV2GetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalTransactionV2GetAll struct {
	client      *Client
	queryParams *JournalTransactionV2GetAllQueryParams
	pathParams  *JournalTransactionV2GetAllPathParams
	method      string
	headers     http.Header
	requestBody JournalTransactionV2GetAllBody
}

func (r JournalTransactionV2GetAll) NewQueryParams() *JournalTransactionV2GetAllQueryParams {
	return &JournalTransactionV2GetAllQueryParams{}
}

type JournalTransactionV2GetAllQueryParams struct {
	PeriodID string `schema:"periodId"`
}

func (p JournalTransactionV2GetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalTransactionV2GetAll) QueryParams() *JournalTransactionV2GetAllQueryParams {
	return r.queryParams
}

func (r JournalTransactionV2GetAll) NewPathParams() *JournalTransactionV2GetAllPathParams {
	return &JournalTransactionV2GetAllPathParams{}
}

type JournalTransactionV2GetAllPathParams struct {
}

func (p *JournalTransactionV2GetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalTransactionV2GetAll) PathParams() *JournalTransactionV2GetAllPathParams {
	return r.pathParams
}

func (r *JournalTransactionV2GetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalTransactionV2GetAll) SetMethod(method string) {
	r.method = method
}

func (r *JournalTransactionV2GetAll) Method() string {
	return r.method
}

func (r JournalTransactionV2GetAll) NewRequestBody() JournalTransactionV2GetAllBody {
	return JournalTransactionV2GetAllBody{}
}

type JournalTransactionV2GetAllBody struct {
}

func (r *JournalTransactionV2GetAll) RequestBody() *JournalTransactionV2GetAllBody {
	return nil
}

func (r *JournalTransactionV2GetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalTransactionV2GetAll) SetRequestBody(body JournalTransactionV2GetAllBody) {
	r.requestBody = body
}

func (r *JournalTransactionV2GetAll) NewResponseBody() *JournalTransactionV2GetAllResponseBody {
	return &JournalTransactionV2GetAllResponseBody{}
}

type JournalTransactionV2GetAllResponseBody JournalTransactions

func (r *JournalTransactionV2GetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v2/journaltransaction", r.PathParams())
	return &u
}

func (r *JournalTransactionV2GetAll) Do() (JournalTransactionV2GetAllResponseBody, error) {
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
