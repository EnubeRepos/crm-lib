package massaction

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIMassActionService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIMassActionService {
	return &APIMassActionService{
		client: client,
	}
}
