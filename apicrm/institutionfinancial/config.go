package institutionfinancial

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIInstitutionFinancialService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIInstitutionFinancialService {
	return &APIInstitutionFinancialService{
		client: client,
	}
}
