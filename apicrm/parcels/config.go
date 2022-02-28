package parcels

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIParcelsService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIParcelsService {
	return &APIParcelsService{
		client: client,
	}
}
