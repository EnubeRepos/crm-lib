package account

import (
	"encoding/json"
)

func (svc *APIAccountService) Get() (ResponseAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccount, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseAccount{}, err
	}

	return convertMarchalResponseAccount(response)
}

func (svc *APIAccountService) GetById(id string) (DomainAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccount, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainAccount{}, err
	}

	return convertMarchalAccount(response)
}

func (svc *APIAccountService) GetByFilter(filter string) (ResponseAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccount, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseAccount{}, err
	}

	return convertMarchalResponseAccount(response)
}

func (svc *APIAccountService) Post(account DomainAccount) (DomainAccount, error) {
	payload, err := json.Marshal(account)

	if err != nil {
		return DomainAccount{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityAccount, payload)
	if err != nil {
		return DomainAccount{}, err
	}

	return convertMarchalAccount(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseAccount(response []byte) (ResponseAccount, error) {
	var result ResponseAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseAccount{}, err
	}

	return result, err
}

func convertMarchalAccount(response []byte) (DomainAccount, error) {
	var result DomainAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainAccount{}, err
	}

	return result, err
}
