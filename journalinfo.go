package multivers

import (
	"context"
	"fmt"
)

const (
	JournalInfoPath = "/api/%s/JournalInfo/%d"
	// JournalInfoPath = "api/{database}/JournalInfo/{journalId}?fiscalYear={fiscalYear}"
)

func NewJournalInfoService(client *Client) *JournalInfoService {
	return &JournalInfoService{Client: client}
}

type JournalInfoService struct {
	Client *Client
}

func (s *JournalInfoService) Get(database string, journalId int, requestParams *JournalInfoGetParams, ctx context.Context) (*JournalInfoGetResponse, error) {
	method := "GET"
	responseBody := NewJournalInfoGetResponse()
	path := fmt.Sprintf(JournalInfoPath, database, journalId)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewJournalInfoGetResponse() *JournalInfoGetResponse {
	return &JournalInfoGetResponse{}
}

type JournalInfoGetResponse JournalInfo

func NewJournalInfoGetParams() *JournalInfoGetParams {
	return &JournalInfoGetParams{}
}

type JournalInfoGetParams struct {
	FiscalYear int `schema:"fiscalYear"`
}

type JournalInfo struct {
	AccountID                     string  `json:"accountId"`
	JournalID                     string  `json:"journalId"`
	AmendmentBalance              float64 `json:"amendmentBalance"`
	BankAccountID                 string  `json:"bankAccountId"`
	CurrencyID                    string  `json:"currencyId"`
	Description                   string  `json:"description"`
	EndBalance                    float64 `json:"endBalance"`
	EntriesToMatchCount           int     `json:"entriesToMatchCount"`
	FiscalYear                    int     `json:"fiscalYear"`
	JournalKind                   int     `json:"journalKind"`
	JournalTransaction            int     `json:"journalTransaction"`
	JournalType                   int     `json:"journalType"`
	LastElectronicTransactionDate string  `json:"lastEletronicTransactionDate"`
	LastTransactionDate           string  `json:"lastTransactionDate"`
	OpeningBalanceJournal         bool    `json:"openingBalanceJournal"`
	RangeID                       string  `json:"rangeId"`
	sequenceNumber                int     `json:"sequenceNumber"`
	SheetNumberDay                int     `json:"sheetNumberDay"`
	SheetNumberPeriod             int     `json:"sheetNumberPeriod"`
	StartBalance                  float64 `json:"startBalance"`
}
