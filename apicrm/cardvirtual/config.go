package cardvirtual

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIVirtualCardService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIVirtualCardService {
	return &APIVirtualCardService{
		client: client,
	}
}
