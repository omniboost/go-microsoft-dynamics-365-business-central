package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewGeneralLedgerJournalPostRequest() GeneralLedgerJournalPostRequest {
	r := GeneralLedgerJournalPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GeneralLedgerJournalPostRequest struct {
	client      *Client
	queryParams *GeneralLedgerJournalPostQueryParams
	pathParams  *GeneralLedgerJournalPostPathParams
	method      string
	headers     http.Header
	requestBody GeneralLedgerJournalPostRequestBody
}

func (r GeneralLedgerJournalPostRequest) NewQueryParams() *GeneralLedgerJournalPostQueryParams {
	return &GeneralLedgerJournalPostQueryParams{}
}

type GeneralLedgerJournalPostQueryParams struct {
}

func (p GeneralLedgerJournalPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GeneralLedgerJournalPostRequest) QueryParams() *GeneralLedgerJournalPostQueryParams {
	return r.queryParams
}

func (r GeneralLedgerJournalPostRequest) NewPathParams() *GeneralLedgerJournalPostPathParams {
	return &GeneralLedgerJournalPostPathParams{}
}

type GeneralLedgerJournalPostPathParams struct {
}

func (p *GeneralLedgerJournalPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GeneralLedgerJournalPostRequest) PathParams() *GeneralLedgerJournalPostPathParams {
	return r.pathParams
}

func (r *GeneralLedgerJournalPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GeneralLedgerJournalPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *GeneralLedgerJournalPostRequest) Method() string {
	return r.method
}

func (r GeneralLedgerJournalPostRequest) NewRequestBody() GeneralLedgerJournalPostRequestBody {
	return GeneralLedgerJournalPostRequestBody{}
}

type GeneralLedgerJournalPostRequestBody struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
	Period      int    `json:"Period,omitempty"`
	Options     struct {
		Post            bool `json:"Post"`
		GenerateVoucher bool `json:"GenerateVoucher"`
	} `json:"Options,omitempty"`
	Lines []struct {
		Account     string   `json:"Account"`
		Amount      float64  `json:"Amount"`
		Currency    string   `json:"Currency,omitempty"`
		Date        DateTime `json:"Date,omitempty"`
		DueDate     DateTime `json:"DueDate,omitempty"`
		Dim1        string   `json:"Dim1,omitempty"`
		Reference   string   `json:"Reference,omitempty"`
		Text        string   `json:"Text,omitempty"`
		Voucher     string   `json:"Voucher,omitempty"`
		JournalType string   `json:"JournalType"`
	}
}

func (r *GeneralLedgerJournalPostRequest) RequestBody() *GeneralLedgerJournalPostRequestBody {
	return &r.requestBody
}

func (r *GeneralLedgerJournalPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GeneralLedgerJournalPostRequest) SetRequestBody(body GeneralLedgerJournalPostRequestBody) {
	r.requestBody = body
}

func (r *GeneralLedgerJournalPostRequest) NewResponseBody() *GeneralLedgerJournalPostResponseBody {
	return &GeneralLedgerJournalPostResponseBody{}
}

type GeneralLedgerJournalPostResponseBody struct {
}

func (r *GeneralLedgerJournalPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("generalledger/journal", r.PathParams())
	return &u
}

func (r *GeneralLedgerJournalPostRequest) Do() (GeneralLedgerJournalPostResponseBody, error) {
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
