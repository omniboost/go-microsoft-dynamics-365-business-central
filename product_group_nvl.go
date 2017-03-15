package multivers

// https://api.online.unit4.nl/V17/Help/Api/GET-api-database-ProductGroupNVL

import (
	"context"
	"fmt"
)

const (
	ProductGroupNVLPath = "/api/%s/ProductGroupNVL"
)

func NewProductGroupNVLService(client *Client) *ProductGroupNVLService {
	return &ProductGroupNVLService{Client: client}
}

type ProductGroupNVLService struct {
	Client *Client
}

func (s *ProductGroupNVLService) Get(database string, ctx context.Context) (*ProductGroupNVLGetResponse, error) {
	method := "GET"
	responseBody := NewProductGroupNVLGetResponse()
	path := fmt.Sprintf(ProductGroupNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductGroupNVLGetResponse() *ProductGroupNVLGetResponse {
	return &ProductGroupNVLGetResponse{}
}

type ProductGroupNVLGetResponse []NVL
