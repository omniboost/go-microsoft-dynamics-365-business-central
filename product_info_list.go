package multivers

// https://api.online.unit4.nl/V14/Help/Api/GET-api-database-ProductInfoList_productId_shortName_description_productGroupId

import (
	"context"
)

const (
	ProductInfoListPath = "/ProductInfoList"
)

func NewProductInfoListService(client *Client) *ProductInfoListService {
	return &ProductInfoListService{Client: client}
}

type ProductInfoListService struct {
	Client *Client
}

func (s *ProductInfoListService) Get(database string, requestParams *ProductInfoListGetParams, ctx context.Context) (*ProductInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewProductInfoListGetResponse()
	path := apiPrefix + ProductInfoListPath + ".json"

	// Process path parameters
	if database != "" {
		path = apiPrefix + "/" + database + ProductInfoListPath + ".json"
	}

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductInfoListGetParams() *ProductInfoListGetParams {
	return &ProductInfoListGetParams{}
}

type ProductInfoListGetParams struct {
	ProductID      string `url:"productId"`
	ShortName      string `url:"shortName"`
	Description    string `url:"description"`
	ProductGroupID string `url:"productGroupId"`
}

func NewProductInfoListGetResponse() *ProductInfoListGetResponse {
	return &ProductInfoListGetResponse{}
}

type ProductInfoListGetResponse []ProductInfo
