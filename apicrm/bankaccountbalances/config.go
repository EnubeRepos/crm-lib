package bankaccount

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIBankAccountBalanceService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIBankAccountBalanceService {
	return &APIBankAccountBalanceService{
		client: client,
	}
}
