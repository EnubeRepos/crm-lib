package user

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIAuthService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIAuthService {
	return &APIAuthService{
		client: client,
	}
}
