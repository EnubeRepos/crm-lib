package image

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIImageService struct {
	client crmapi.CRMAPIClient
}

func NewAPIImageService(client crmapi.CRMAPIClient) *APIImageService {
	return &APIImageService{
		client: client,
	}
}
