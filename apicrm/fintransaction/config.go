package fintransaction

import "github.com/EnubeRepos/crm-lib/internal/crmapi"

type APIFinTransactionService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIFinTransactionService {
	return &APIFinTransactionService{
		client: client,
	}
}
