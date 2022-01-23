package accountsrv

import (
	"encoding/json"

	"github.com/EnubeRepos/CRM_framework/internal/core/domains"
	"github.com/EnubeRepos/CRM_framework/internal/crmapi"
)

type APIAccountService struct {
	client crmapi.CRMAPIClient
}

func NewAPIAccountService(client crmapi.CRMAPIClient) *APIAccountService {
	return &APIAccountService{
		client: client,
	}
}

func (svc *APIAccountService) GetAccount() (domains.ResponseAccount, error) {
	response, err := svc.client.CRMHandlerGetService(domains.EntityAccount, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return domains.ResponseAccount{}, err
	}

	return convertMarchalResponseAccount(response)
}

func (svc *APIAccountService) GetAccountById(id string) (domains.Account, error) {
	response, err := svc.client.CRMHandlerGetService(domains.EntityAccount, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return domains.Account{}, err
	}

	return convertMarchalAccount(response)
}

func (svc *APIAccountService) GetAccountByFilter(filter string) (domains.ResponseAccount, error) {
	response, err := svc.client.CRMHandlerGetService(domains.EntityAccount, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return domains.ResponseAccount{}, err
	}

	return convertMarchalResponseAccount(response)
}

func (svc *APIAccountService) PostAccount(account domains.Account) (domains.Account, error) {
	payload, err := json.Marshal(account)

	if err != nil {
		return domains.Account{}, err
	}

	response, err := svc.client.CRMHandlerPostService(domains.EntityAccount, payload)
	if err != nil {
		return domains.Account{}, err
	}

	return convertMarchalAccount(response)
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

func convertMarchalAccount(response []byte) (domains.Account, error) {
	var result domains.Account

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return domains.Account{}, err
	}

	return result, err
}
