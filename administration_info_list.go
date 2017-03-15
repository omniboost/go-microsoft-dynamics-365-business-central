package multivers

import "context"

const (
	AdministrationInfoListPath = "/api/AdministrationInfoList"
)

func NewAdministrationInfoListService(client *Client) *AdministrationInfoListService {
	return &AdministrationInfoListService{Client: client}
}

type AdministrationInfoListService struct {
	Client *Client
}

func (s *AdministrationInfoListService) Get(groupID string, ctx context.Context) (*AdministrationInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationInfoListGetResponse()
	path := AdministrationInfoListPath + ".json"

	// Process parameters
	if groupID != "" {
		path = AdministrationInfoListPath + "/" + groupID + ".json"
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

func NewAdministrationInfoListGetResponse() *AdministrationInfoListGetResponse {
	return &AdministrationInfoListGetResponse{}
}

type AdministrationInfoListGetResponse []Administration
