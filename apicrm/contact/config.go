package contact

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIContactService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIContactService {
	return &APIContactService{
		client: client,
	}
}
