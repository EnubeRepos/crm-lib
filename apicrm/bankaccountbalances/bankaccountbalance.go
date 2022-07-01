package bankaccountbalances

import (
	"encoding/json"
	"fmt"
)

func (svc *APIBankAccountBalanceService) Get() (ResponseBankAccountBalance, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBankAccountBalance, "")
	if err != nil {
		return ResponseBankAccountBalance{}, err
	}

	return convertMarchalResponseBankAccountBalance(response)
}

func (svc *APIBankAccountBalanceService) GetById(id string) (DomainBankAccountBalance, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainBankAccountBalance{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	responseHttp, err := svc.client.CRMHandlerGetService(EntityBankAccountBalance, "/"+id)
	if err != nil {
		return DomainBankAccountBalance{}, err
	}

	response := DomainBankAccountBalance{}

	err = convertMarchalBankAccountBalance(responseHttp, &response)
	return response, err
}

func (svc *APIBankAccountBalanceService) GetByFilter(filter string) (ResponseBankAccountBalance, error) {
	if filter == "" {
		return ResponseBankAccountBalance{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityBankAccountBalance, "?"+filter)
	if err != nil {
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

func (svc *APIBankAccountBalanceService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(EntityBankAccountBalance, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (svc *APIBankAccountBalanceService) Put(Registration DomainBankAccountBalanceCreateRequest) (DomainBankAccountBalanceCreateResponse, error) {
	payload, err := json.Marshal(Registration)

	if err != nil {
		return DomainBankAccountBalanceCreateResponse{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPutService(EntityBankAccountBalance+"/"+Registration.ID, payload)
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
