package operationfeeapplication

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APITedService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APITedService {
	return &APITedService{
		client: client,
	}
}
