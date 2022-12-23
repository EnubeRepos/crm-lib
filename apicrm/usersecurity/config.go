package usersecurity

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIUserService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIUserService {
	return &APIUserService{
		client: client,
	}
}
