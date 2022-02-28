package bankaccount

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIBankAccountService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIBankAccountService {
	return &APIBankAccountService{
		client: client,
	}
}
