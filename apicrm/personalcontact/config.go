package personalcontact

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIPersonalContactService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIPersonalContactService {
	return &APIPersonalContactService{
		client: client,
	}
}
