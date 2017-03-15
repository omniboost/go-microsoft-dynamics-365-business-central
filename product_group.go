package multivers

// https://api.online.unit4.nl/V16/Help/Api/GET-api-database-ProductGroup-productGroupId

import (
	"context"
	"fmt"
)

const (
	ProductGroupPath = "/api/%s/ProductGroup/%d.json"
)

func NewProductGroupService(client *Client) *ProductGroupService {
	return &ProductGroupService{Client: client}
}

type ProductGroupService struct {
	Client *Client
}

func (s *ProductGroupService) Get(database string, productGroupID int, ctx context.Context) (*ProductGroupGetResponse, error) {
	method := "GET"
	responseBody := NewProductGroupGetResponse()
	path := fmt.Sprintf(ProductGroupPath, database, productGroupID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductGroupGetResponse() *ProductGroupGetResponse {
	return &ProductGroupGetResponse{}
}

type ProductGroupGetResponse ProductGroup

type ProductGroup struct {
	Messages                   []Message `json:"messages"`
	CanChange                  bool      `json:"canChange"`
	Description                string    `json:"description"`
	ProductGroupID             string    `json:"productGroupId"`
	ProjectSurchargePercentage float64   `json:"projectSurchargePercentage"`
}
