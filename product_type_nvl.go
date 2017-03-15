package multivers

// https://api.online.unit4.nl/V14/Help/Api/GET-api-database-ProductTypeNVL_productId_shortName_description_productGroupId

import (
	"context"
)

const (
	ProductTypeNVLPath = "/ProductTypeNVL"
)

func NewProductTypeNVLService(client *Client) *ProductTypeNVLService {
	return &ProductTypeNVLService{Client: client}
}

type ProductTypeNVLService struct {
	Client *Client
}

func (s *ProductTypeNVLService) Get(database string, ctx context.Context) (*ProductTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewProductTypeNVLGetResponse()
	path := apiPrefix + ProductTypeNVLPath + ".json"

	// Process path parameters
	if database != "" {
		path = apiPrefix + "/" + database + ProductTypeNVLPath + ".json"
	}

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductTypeNVLGetResponse() *ProductTypeNVLGetResponse {
	return &ProductTypeNVLGetResponse{}
}

type ProductTypeNVLGetResponse []NVL
