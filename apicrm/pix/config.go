package pix

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIPixService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIPixService {
	return &APIPixService{
		client: client,
	}
}
