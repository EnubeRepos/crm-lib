package preregistration

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIPRERegistrationService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIPRERegistrationService {
	return &APIPRERegistrationService{
		client: client,
	}
}
