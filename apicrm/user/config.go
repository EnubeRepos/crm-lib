package user

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIUserService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIUserService {
	return &APIUserService{
		client: client,
	}
}
