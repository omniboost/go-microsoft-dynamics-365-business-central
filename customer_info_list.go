package multivers

// https://api.online.unit4.nl/V19/Help/Api/GET-api-database-CustomerInfoList

import (
	"context"
	"fmt"
)

const (
	CustomerInfoListPath = "api/%s/CustomerInfoList.json"
)

func NewCustomerInfoListService(client *Client) *CustomerInfoListService {
	return &CustomerInfoListService{Client: client}
}

type CustomerInfoListService struct {
	Client *Client
}

func (s *CustomerInfoListService) Get(database string, ctx context.Context) (*CustomerInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerInfoListGetResponse()
	path := fmt.Sprintf(CustomerInfoListPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerInfoListGetResponse() *CustomerInfoListGetResponse {
	return &CustomerInfoListGetResponse{}
}

type CustomerInfoListGetResponse []CustomerInfoList

type CustomerInfoList struct {
	AccountManagerID    string `json:"accountManagerId"`
	ApplyOrderSurcharge bool   `json:"applyOrderSurcharge"`
	BusinessNumber      string `json:"businessNumber"`
	ChargeVatTypeID     int    `json:"chargeVatTypeId"`
	City                string `json:"city"`
	CocCity             string `json:"cocCity"`
	CocDate             *Time  `json:"cocDate"`
	CocRegistration     string `json:"cocRegistration"`
	ContactPerson       string `json:"contactPerson"`
	CountryID           string `json:"countryId"`
	CreditSqueezeID     string `json:"creditSqueezeId"`
	CurrencyID          string `json:"currencyId"`
	CustomerID          string `json:"customerId"`
	CustomerStateID     string `json:"customerStateId"`
	DateChanged         *Time  `json:"dateChanged"`
	DateCreated         *Time  `json:"dateCreated"`
	Email               string `json:"email"`
	Fax                 string `json:"fax"`
	Homepage            string `json:"homepage"`
	IsDunForPayment     bool   `json:"isDunForPayment"`
	LanguageID          string `json:"languageId"`
	MobilePhone         string `json:"mobilePhone"`
	Name                string `json:"name"`
	OrganizationID      int    `json:"organizationId"`
	PaymentConditionID  string `json:"paymentConditionId"`
	RevenueAccountID    string `json:"revenueAccountId"`
	ShortName           string `json:"shortName"`
	Street1             string `json:"street1"`
	Street2             string `json:"street2"`
	Telephone           string `json:"telephone"`
	VatNumber           string `json:"vatNumber"`
	VatScenarioID       int    `json:"vatScenarioId"`
	VatVerificationDate *Time  `json:"vatVerificationDate"`
	ZipCode             string `json:"zipCode"`
}
