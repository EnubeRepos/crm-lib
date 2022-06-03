package userregistration

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIUserRegistrationService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIUserRegistrationService {
	return &APIUserRegistrationService{
		client: client,
	}
}
