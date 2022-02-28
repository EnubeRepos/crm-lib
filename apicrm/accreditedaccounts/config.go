package accreditedaccounts

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIAccreditedAccountsService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIAccreditedAccountsService {
	return &APIAccreditedAccountsService{
		client: client,
	}
}
