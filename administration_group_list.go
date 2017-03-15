package multivers

import "context"

const (
	AdministrationGroupListPath = "/api/AdministrationGroupList"
)

func NewAdministrationGroupListService(client *Client) *AdministrationGroupListService {
	return &AdministrationGroupListService{Client: client}
}

type AdministrationGroupListService struct {
	Client *Client
}

func (s *AdministrationGroupListService) Get(ctx context.Context) (*AdministrationGroupListGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationGroupListGetResponse()
	path := AdministrationGroupListPath + ".json"

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAdministrationGroupListGetResponse() *AdministrationGroupListGetResponse {
	return &AdministrationGroupListGetResponse{}
}

type AdministrationGroupListGetResponse []AdministrationGroup
