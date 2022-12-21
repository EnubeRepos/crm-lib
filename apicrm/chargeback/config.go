package chargeback

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIChargebackService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIChargebackService {
	return &APIChargebackService{
		client: client,
	}
}
