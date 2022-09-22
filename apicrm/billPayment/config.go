package billPayment

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIBillPaymentService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIBillPaymentService {
	return &APIBillPaymentService{
		client: client,
	}
}
