package multivers

// https://api.online.unit4.nl/V14/Help/Api/GET-api-database-ProjectEntryTypeNVL_productId_shortName_description_productGroupId

import "context"

const (
	ProjectEntryTypeNVLPath = "/ProjectEntryTypeNVL"
)

func NewProjectEntryTypeNVLService(client *Client) *ProjectEntryTypeNVLService {
	return &ProjectEntryTypeNVLService{Client: client}
}

type ProjectEntryTypeNVLService struct {
	Client *Client
}

func (s *ProjectEntryTypeNVLService) Get(database string, ctx context.Context) (*ProjectEntryTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewProjectEntryTypeNVLGetResponse()
	path := apiPrefix + ProjectEntryTypeNVLPath + ".json"

	// Process path parameters
	if database != "" {
		path = apiPrefix + "/" + database + ProjectEntryTypeNVLPath + ".json"
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

func NewProjectEntryTypeNVLGetResponse() *ProjectEntryTypeNVLGetResponse {
	return &ProjectEntryTypeNVLGetResponse{}
}

type ProjectEntryTypeNVLGetResponse []NVL
