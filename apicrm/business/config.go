package registration

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIRegistrationService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIRegistrationService {
	return &APIRegistrationService{
		client: client,
	}
}
