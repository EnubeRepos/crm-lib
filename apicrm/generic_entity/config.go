package generic

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIGenericService struct {
	client     crmapi.CRMAPIClient
	entityName string
}

func NewAPIGenericService(client crmapi.CRMAPIClient, entityName string) *APIGenericService {
	return &APIGenericService{
		client:     client,
		entityName: entityName,
	}
}
