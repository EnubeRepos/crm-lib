package bankaccount

import (
	"encoding/json"
)

func (svc *APIBankAccountBalanceService) Get() (ResponseBankAccountBalance, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBankAccountBalance, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseBankAccountBalance{}, err
	}

	return convertMarchalResponseBankAccountBalance(response)
}

func (svc *APIBankAccountBalanceService) GetById(id string) (DomainBankAccountBalance, error) {
	responseHttp, err := svc.client.CRMHandlerGetService(EntityBankAccountBalance, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainBankAccountBalance{}, err
	}

	response := DomainBankAccountBalance{}

	err = convertMarchalBankAccountBalance(responseHttp, &response)
	return response, err
}

func (svc *APIBankAccountBalanceService) GetByFilter(filter string) (ResponseBankAccountBalance, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBankAccountBalance, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseBankAccountBalance{}, err
	}

	return convertMarchalResponseBankAccountBalance(response)
}

func (svc *APIBankAccountBalanceService) Post(BankAccountBalance DomainBankAccountBalanceCreateRequest) (DomainBankAccountBalanceCreateResponse, error) {
	payload, err := json.Marshal(BankAccountBalance)

	if err != nil {
		return DomainBankAccountBalanceCreateResponse{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPostService(EntityBankAccountBalance, payload)
	if err != nil {
		return DomainBankAccountBalanceCreateResponse{}, err
	}

	response := DomainBankAccountBalanceCreateResponse{}

	err = convertMarchalBankAccountBalance(responseHttp, &response)
	return response, err
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseBankAccountBalance(response []byte) (ResponseBankAccountBalance, error) {
	var result ResponseBankAccountBalance

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseBankAccountBalance{}, err
	}

	return result, nil
}

func convertMarchalBankAccountBalance(responseByte []byte, responseObj interface{}) error {
	err := json.Unmarshal(responseByte, responseObj)
	if err != nil {
		return err
	}

	return nil
}
