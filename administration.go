package multivers

import "context"

const (
	AdministrationPath = "/api/administration"
)

func NewAdministrationService(client *Client) *AdministrationService {
	return &AdministrationService{Client: client}
}

type AdministrationService struct {
	Client *Client
}

func (s *AdministrationService) Get(administrationID string, ctx context.Context) (*AdministrationGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationGetResponse()
	path := AdministrationPath + "/" + administrationID + ".json"

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAdministrationGetResponse() *AdministrationGetResponse {
	return &AdministrationGetResponse{}
}

type AdministrationGetResponse Administration

type Administration struct {
	BackupDate          string   `json:"backupDate"`
	Version             int      `json:"version"`
	GroupID             int      `json:"groupId"`
	Code                string   `json:"code"`
	ShortName           string   `json:"shortName"`
	PreviousOnlineState int      `json:"previousOnlineState"`
	OnlineState         int      `json:"onlineState"`
	Description         string   `json:"description"`
	AdministrationID    string   `json:"administrationId"`
	Users               []string `json:"users"`
	ReportPath          string   `json:"reportPath"`
}
