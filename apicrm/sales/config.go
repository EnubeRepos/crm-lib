package sales

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APISalesService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APISalesService {
	return &APISalesService{
		client: client,
	}
}
