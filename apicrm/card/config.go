package card

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIDebitCardService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIDebitCardService {
	return &APIDebitCardService{
		client: client,
	}
}

type APIVirtualCardService struct {
	client crmapi.CRMAPIClient
}

func NewVirtual(client crmapi.CRMAPIClient) *APIVirtualCardService {
	return &APIVirtualCardService{
		client: client,
	}
}
