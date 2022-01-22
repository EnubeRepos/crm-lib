package accountsrv

import (
	"encoding/json"

	"github.com/EnubeRepos/CRM_framework/internal/core/domains"
	"github.com/EnubeRepos/CRM_framework/internal/crmapi"
)

type APIAccountService struct {
	// crmAPIRepo ports.CRMAPIRepository
	client crmapi.CRMAPIClient
}

// NewAPIAccountService Initialization ...
// func NewAPIAccountService(repository ports.CRMAPIRepository, client crmapi.CRMAPIClient) *APIAccountService {
// 	return &APIAccountService{
// 		// crmAPIRepo: repository,
// 		client: client,
// 	}
// }

func NewAPIAccountService(client crmapi.CRMAPIClient) *APIAccountService {
	return &APIAccountService{
		// crmAPIRepo: repository,
		client: client,
	}
}

func (svc *APIAccountService) GetAccount() (domains.ResponseAccount, error) {
	response, err := svc.client.CRMHandlerGetService("Account", "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return domains.ResponseAccount{}, err
	}

	return convertMarchalResponseAccount(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseAccount(response []byte) (domains.ResponseAccount, error) {
	var result domains.ResponseAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return domains.ResponseAccount{}, err
	}

	return result, err
}
