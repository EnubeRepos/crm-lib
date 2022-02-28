package eventtraching

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIEventTrachingService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIEventTrachingService {
	return &APIEventTrachingService{
		client: client,
	}
}
