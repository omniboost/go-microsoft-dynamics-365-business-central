package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/odata"
	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCustomerInfoListGetRequest() CustomerInfoListGetRequest {
	r := CustomerInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCustomerInfoListGetQueryParams()
	r.pathParams = r.NewCustomerInfoListGetPathParams()
	r.requestBody = r.NewCustomerInfoListGetRequestBody()
	return r
}

type CustomerInfoListGetRequest struct {
	client      *Client
	queryParams *CustomerInfoListGetQueryParams
	pathParams  *CustomerInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody CustomerInfoListGetRequestBody
}

func (r CustomerInfoListGetRequest) NewCustomerInfoListGetQueryParams() *CustomerInfoListGetQueryParams {
	body := CustomerInfoListGetResponseBody{{}}
	fields, _ := utils.Fields(body[0])
	return &CustomerInfoListGetQueryParams{
		Select: odata.NewSelect(fields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CustomerInfoListGetQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CustomerInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerInfoListGetRequest) QueryParams() *CustomerInfoListGetQueryParams {
	return r.queryParams
}

func (r CustomerInfoListGetRequest) NewCustomerInfoListGetPathParams() *CustomerInfoListGetPathParams {
	return &CustomerInfoListGetPathParams{}
}

type CustomerInfoListGetPathParams struct {
}

func (p *CustomerInfoListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerInfoListGetRequest) PathParams() *CustomerInfoListGetPathParams {
	return r.pathParams
}

func (r *CustomerInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInfoListGetRequest) Method() string {
	return r.method
}

func (r CustomerInfoListGetRequest) NewCustomerInfoListGetRequestBody() CustomerInfoListGetRequestBody {
	return CustomerInfoListGetRequestBody{}
}

type CustomerInfoListGetRequestBody struct{}

func (r *CustomerInfoListGetRequest) RequestBody() *CustomerInfoListGetRequestBody {
	return &r.requestBody
}

func (r *CustomerInfoListGetRequest) SetRequestBody(body CustomerInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerInfoListGetRequest) NewResponseBody() *CustomerInfoListGetResponseBody {
	return &CustomerInfoListGetResponseBody{}
}

type CustomerInfoListGetResponseBody []struct {
	AccountManagerID                    string      `json:"accountManagerId"`
	ApplyOrderSurcharge                 bool        `json:"applyOrderSurcharge"`
	BusinessNumber                      string      `json:"businessNumber"`
	ChargeVatTypeID                     int         `json:"chargeVatTypeId"`
	City                                string      `json:"city"`
	CocCity                             string      `json:"cocCity"`
	CocDate                             string      `json:"cocDate"`
	CocRegistration                     string      `json:"cocRegistration"`
	CollectiveInvoiceSystemID           string      `json:"collectiveInvoiceSystemId"`
	CombineInvoicesForElectronicBanking bool        `json:"combineInvoicesForElectronicBanking"`
	ContactPerson                       string      `json:"contactPerson"`
	CountryID                           string      `json:"countryId"`
	CreditLimit                         float64     `json:"creditLimit"`
	CreditSqueezeID                     string      `json:"creditSqueezeId"`
	CurrencyID                          string      `json:"currencyId"`
	CustomerGroupID                     int         `json:"customerGroupId"`
	CustomerID                          string      `json:"customerId"`
	CustomerOrganizationID              int         `json:"customerOrganizationId"`
	CustomerStateID                     string      `json:"customerStateId"`
	DateChanged                         string      `json:"dateChanged"`
	DateCreated                         string      `json:"dateCreated"`
	DeliveryConditionID                 string      `json:"deliveryConditionId"`
	DiscountPercentage                  float64     `json:"discountPercentage"`
	Email                               string      `json:"email"`
	Fax                                 string      `json:"fax"`
	GovernmentDigitalID                 string      `json:"governmentDigitalId"`
	Homepage                            string      `json:"homepage"`
	IncludeVatOnOrderByDefault          bool        `json:"includeVatOnOrderByDefault"`
	IntrastatGoodsCodeID                int         `json:"intrastatGoodsCodeId"`
	IntrastatGoodsDistributionID        interface{} `json:"intrastatGoodsDistributionId"`
	IntrastatStatSystemID               interface{} `json:"intrastatStatSystemId"`
	IntrastatTrafficRegionID            interface{} `json:"intrastatTrafficRegionId"`
	IntrastatTransactionTypeID          string      `json:"intrastatTransactionTypeId"`
	IntrastatTransportTypeID            interface{} `json:"intrastatTransportTypeId"`
	InvoiceOnBehalfOfMembers            bool        `json:"invoiceOnBehalfOfMembers"`
	IsDunForPayment                     bool        `json:"isDunForPayment"`
	IsInFactoring                       bool        `json:"isInFactoring"`
	IsPaymentRefRequired                bool        `json:"isPaymentRefRequired"`
	IsPurchaseOrganization              bool        `json:"isPurchaseOrganization"`
	LanguageID                          string      `json:"languageId"`
	MobilePhone                         string      `json:"mobilePhone"`
	Name                                string      `json:"name"`
	OrganizationID                      int         `json:"organizationId"`
	PaymentConditionID                  string      `json:"paymentConditionId"`
	PricelistID                         string      `json:"pricelistId"`
	PrintPurchaseDetails                bool        `json:"printPurchaseDetails"`
	PurchaseOrganizationID              string      `json:"purchaseOrganizationId"`
	PurchaseOrganizationMemberID        string      `json:"purchaseOrganizationMemberId"`
	RevenueAccountID                    string      `json:"revenueAccountId"`
	ShortName                           string      `json:"shortName"`
	Street1                             string      `json:"street1"`
	Street2                             string      `json:"street2"`
	Telephone                           string      `json:"telephone"`
	UsesUBLInvoice                      bool        `json:"usesUBLInvoice"`
	VatNumber                           string      `json:"vatNumber"`
	VatScenarioID                       interface{} `json:"vatScenarioId"`
	VatVerificationDate                 string      `json:"vatVerificationDate"`
	ZipCode                             string      `json:"zipCode"`
}

func (r *CustomerInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CustomerInfoList", r.PathParams())
}

func (r *CustomerInfoListGetRequest) Do() (CustomerInfoListGetResponseBody, error) {
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
