package blocklistexternal

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIBlockListExternalService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIBlockListExternalService {
	return &APIBlockListExternalService{
		client: client,
	}
}
