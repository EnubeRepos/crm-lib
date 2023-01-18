package occupation

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIOccupationService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIOccupationService {
	return &APIOccupationService{
		client: client,
	}
}
