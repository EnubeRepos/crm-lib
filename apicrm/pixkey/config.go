package pixkey

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIPixKeyService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIPixKeyService {
	return &APIPixKeyService{
		client: client,
	}
}
