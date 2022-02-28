package account

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIAccountService struct {
	client crmapi.CRMAPIClient
}

func NewAPIAccountService(client crmapi.CRMAPIClient) *APIAccountService {
	return &APIAccountService{
		client: client,
	}
}
