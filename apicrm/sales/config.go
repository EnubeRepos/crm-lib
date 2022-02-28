package sales

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APISalesService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APISalesService {
	return &APISalesService{
		client: client,
	}
}
