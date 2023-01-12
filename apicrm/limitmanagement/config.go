package limitmanagement

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APILimitManagementService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APILimitManagementService {
	return &APILimitManagementService{
		client: client,
	}
}
