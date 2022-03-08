package bankaccount

import (
	"encoding/json"
)

func (svc *APIBankAccountService) Get() (ResponseBankAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBankAccount, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseBankAccount{}, err
	}

	return convertMarchalResponseBankAccount(response)
}

func (svc *APIBankAccountService) GetById(id string) (DomainBankAccount, error) {
	responseHttp, err := svc.client.CRMHandlerGetService(EntityBankAccount, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainBankAccount{}, err
	}

	response := DomainBankAccount{}

	err = convertMarchalBankAccount(responseHttp, &response)
	return response, err
}

func (svc *APIBankAccountService) GetByFilter(filter string) (ResponseBankAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBankAccount, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseBankAccount{}, err
	}

	return convertMarchalResponseBankAccount(response)
}

func (svc *APIBankAccountService) Post(BankAccount DomainBankAccountCreateRequest) (DomainBankAccountCreateResponse, error) {
	payload, err := json.Marshal(BankAccount)

	if err != nil {
		return DomainBankAccountCreateResponse{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPostService(EntityBankAccount, payload)
	if err != nil {
		return DomainBankAccountCreateResponse{}, err
	}

	response := DomainBankAccountCreateResponse{}

	err = convertMarchalBankAccount(responseHttp, &response)
	return response, err
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseBankAccount(response []byte) (ResponseBankAccount, error) {
	var result ResponseBankAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseBankAccount{}, err
	}

	return result, nil
}

func convertMarchalBankAccount(responseByte []byte, responseObj interface{}) error {
	err := json.Unmarshal(responseByte, responseObj)
	if err != nil {
		return err
	}

	return nil
}
