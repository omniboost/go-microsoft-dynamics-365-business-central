package multivers

import "context"

const (
	VatCodeInfoListPath = "/VatCodeInfoList"
)

func NewVatCodeInfoListService(client *Client) *VatCodeInfoListService {
	return &VatCodeInfoListService{Client: client}
}

type VatCodeInfoListService struct {
	Client *Client
}

func (s *VatCodeInfoListService) Get(database string, ctx context.Context) (*VatCodeInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewVatCodeInfoListGetResponse()
	path := apiPrefix + VatCodeInfoListPath + ".json"

	// Process path parameters
	if database != "" {
		path = apiPrefix + "/" + database + VatCodeInfoListPath + ".json"
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

func NewVatCodeInfoListGetResponse() *VatCodeInfoListGetResponse {
	return &VatCodeInfoListGetResponse{}
}

type VatCodeInfoListGetResponse []VatCodeInfo
