package parcels

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIParcelsService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIParcelsService {
	return &APIParcelsService{
		client: client,
	}
}
