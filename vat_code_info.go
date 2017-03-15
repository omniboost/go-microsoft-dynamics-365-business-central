package multivers

import (
	"context"
	"fmt"
)

const (
	VatCodeInfoPath = "/api/%s/VatCodeInfo/%d.json"
	// VatCodeInfoPath = "/api/{database}/VatCodeInfo/{vatCode}"
)

func NewVatCodeInfoService(client *Client) *VatCodeInfoService {
	return &VatCodeInfoService{Client: client}
}

type VatCodeInfoService struct {
	Client *Client
}

func (s *VatCodeInfoService) Get(database string, vatCodeID int, ctx context.Context) (*VatCodeInfoGetResponse, error) {
	method := "GET"
	responseBody := NewVatCodeInfoGetResponse()
	path := fmt.Sprintf(VatCodeInfoPath, database, vatCodeID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewVatCodeInfoGetResponse() *VatCodeInfoGetResponse {
	return &VatCodeInfoGetResponse{}
}

type VatCodeInfoGetResponse VatCodeInfo

type VatCodeInfo struct {
	Description   string  `json:"description"`
	VatCodeID     int     `json:"vatCodeID"`
	VatPercentage float64 `json:"vatPercentage"`
}
