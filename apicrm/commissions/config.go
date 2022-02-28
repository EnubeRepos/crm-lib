package commissions

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APICommissionsService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APICommissionsService {
	return &APICommissionsService{
		client: client,
	}
}
