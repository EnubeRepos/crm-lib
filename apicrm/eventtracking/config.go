package eventtracking

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIEventTrackingService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIEventTrackingService {
	return &APIEventTrackingService{
		client: client,
	}
}
