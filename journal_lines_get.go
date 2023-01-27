package central

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewJournalLinesGet() JournalLinesGet {
	r := JournalLinesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalLinesGet struct {
	client      *Client
	queryParams *JournalLinesGetQueryParams
	pathParams  *JournalLinesGetPathParams
	method      string
	headers     http.Header
	requestBody JournalLinesGetBody
}

func (r JournalLinesGet) NewQueryParams() *JournalLinesGetQueryParams {
	return &JournalLinesGetQueryParams{}
}

type JournalLinesGetQueryParams struct {
}

func (p JournalLinesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalLinesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r JournalLinesGet) NewPathParams() *JournalLinesGetPathParams {
	return &JournalLinesGetPathParams{}
}

type JournalLinesGetPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	JournalID       string `schema:"journalID"`
}

func (p *JournalLinesGetPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"journalID":       p.JournalID,
	}
}

func (r *JournalLinesGet) PathParams() *JournalLinesGetPathParams {
	return r.pathParams
}

func (r *JournalLinesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalLinesGet) SetMethod(method string) {
	r.method = method
}

func (r *JournalLinesGet) Method() string {
	return r.method
}

func (r JournalLinesGet) NewRequestBody() JournalLinesGetBody {
	return JournalLinesGetBody{}
}

type JournalLinesGetBody struct {
}

func (r *JournalLinesGet) RequestBody() *JournalLinesGetBody {
	return nil
}

func (r *JournalLinesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalLinesGet) SetRequestBody(body JournalLinesGetBody) {
	r.requestBody = body
}

func (r *JournalLinesGet) NewResponseBody() *JournalLinesGetResponseBody {
	return &JournalLinesGetResponseBody{}
}

type JournalLinesGetResponseBody struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		OdataEtag              string    `json:"@odata.etag"`
		ID                     string    `json:"id"`
		JournalID              string    `json:"journalId"`
		JournalDisplayName     string    `json:"journalDisplayName"`
		LineNumber             int       `json:"lineNumber"`
		AccountType            string    `json:"accountType"`
		AccountID              string    `json:"accountId"`
		AccountNumber          string    `json:"accountNumber"`
		PostingDate            string    `json:"postingDate"`
		DocumentNumber         string    `json:"documentNumber"`
		ExternalDocumentNumber string    `json:"externalDocumentNumber"`
		Amount                 int       `json:"amount"`
		Description            string    `json:"description"`
		Comment                string    `json:"comment"`
		TaxCode                string    `json:"taxCode"`
		BalanceAccountType     string    `json:"balanceAccountType"`
		BalancingAccountID     string    `json:"balancingAccountId"`
		BalancingAccountNumber string    `json:"balancingAccountNumber"`
		LastModifiedDateTime   time.Time `json:"lastModifiedDateTime"`
	} `json:"value"`
}

func (r *JournalLinesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/companies({{.companyID}})/journals({{.journalID}})/journalLines", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *JournalLinesGet) Do() (JournalLinesGetResponseBody, error) {
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
