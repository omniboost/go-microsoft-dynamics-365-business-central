package dkplus

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewGeneralLedgerAccountGetRequest() GeneralLedgerAccountGetRequest {
	r := GeneralLedgerAccountGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GeneralLedgerAccountGetRequest struct {
	client      *Client
	queryParams *GeneralLedgerAccountGetQueryParams
	pathParams  *GeneralLedgerAccountGetPathParams
	method      string
	headers     http.Header
	requestBody GeneralLedgerAccountGetRequestBody
}

func (r GeneralLedgerAccountGetRequest) NewQueryParams() *GeneralLedgerAccountGetQueryParams {
	return &GeneralLedgerAccountGetQueryParams{}
}

type GeneralLedgerAccountGetQueryParams struct {
}

func (p GeneralLedgerAccountGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GeneralLedgerAccountGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GeneralLedgerAccountGetRequest) NewPathParams() *GeneralLedgerAccountGetPathParams {
	return &GeneralLedgerAccountGetPathParams{}
}

type GeneralLedgerAccountGetPathParams struct {
}

func (p *GeneralLedgerAccountGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GeneralLedgerAccountGetRequest) PathParams() *GeneralLedgerAccountGetPathParams {
	return r.pathParams
}

func (r *GeneralLedgerAccountGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GeneralLedgerAccountGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GeneralLedgerAccountGetRequest) Method() string {
	return r.method
}

func (r GeneralLedgerAccountGetRequest) NewRequestBody() GeneralLedgerAccountGetRequestBody {
	return GeneralLedgerAccountGetRequestBody{}
}

type GeneralLedgerAccountGetRequestBody struct {
}

func (r *GeneralLedgerAccountGetRequest) RequestBody() *GeneralLedgerAccountGetRequestBody {
	return &r.requestBody
}

func (r *GeneralLedgerAccountGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GeneralLedgerAccountGetRequest) SetRequestBody(body GeneralLedgerAccountGetRequestBody) {
	r.requestBody = body
}

func (r *GeneralLedgerAccountGetRequest) NewResponseBody() *GeneralLedgerAccountGetResponseBody {
	return &GeneralLedgerAccountGetResponseBody{}
}

type GeneralLedgerAccountGetResponseBody []struct {
	Number         string    `json:"Number"`
	Name           string    `json:"Name"`
	Group          int       `json:"Group"`
	AccountType    int       `json:"AccountType"`
	TaxCode        string    `json:"TaxCode"`
	Modified       time.Time `json:"Modified"`
	Closed         bool      `json:"Closed"`
	UseCurrency    bool      `json:"UseCurrency"`
	Currency       string    `json:"Currency"`
	SumReference   string    `json:"SumReference"`
	ReverseAccount string    `json:"ReverseAccount"`
}

func (r *GeneralLedgerAccountGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("generalLedger/account", r.PathParams())
	return &u
}

func (r *GeneralLedgerAccountGetRequest) Do() (GeneralLedgerAccountGetResponseBody, error) {
	// if r.QueryParams().ConsumerToken == "" {
	// 	r.QueryParams().ConsumerToken = r.client.ConsumerToken()
	// }

	// if r.QueryParams().EmployeeToken == "" {
	// 	r.QueryParams().EmployeeToken = r.client.EmployeeToken()
	// }

	// if r.QueryParams().ExpirationDate.IsZero() {
	// 	r.QueryParams().ExpirationDate = Date{time.Now().AddDate(0, 0, 1)}
	// }

	// fetch a new token if it isn't set already
	// if r.client.token == "" {
	// 	var err error
	// 	r.client.token, err = r.client.NewToken()
	// 	if err != nil {
	// 		return *r.NewResponseBody(), err
	// 	}
	// }

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
