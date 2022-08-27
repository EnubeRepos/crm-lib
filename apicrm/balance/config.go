package balance

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIBalanceService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIBalanceService {
	return &APIBalanceService{
		client: client,
	}
}
