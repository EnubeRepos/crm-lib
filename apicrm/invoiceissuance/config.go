package invoiceissuance

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIInvoiceIssuanceService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIInvoiceIssuanceService {
	return &APIInvoiceIssuanceService{
		client: client,
	}
}
