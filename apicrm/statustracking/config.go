package statustracking

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIStatusTrackingService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIStatusTrackingService {
	return &APIStatusTrackingService{
		client: client,
	}
}
