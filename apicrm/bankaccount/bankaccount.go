package bankaccount

import (
	"encoding/json"
	"fmt"
)

func (svc *APIBankAccountService) Get() (ResponseBankAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBankAccount, "")
	if err != nil {
		return ResponseBankAccount{}, err
	}

	return convertMarchalResponseBankAccount(response)
}

func (svc *APIBankAccountService) GetById(id string) (DomainBankAccount, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainBankAccount{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	responseHttp, err := svc.client.CRMHandlerGetService(EntityBankAccount, "/"+id)
	if err != nil {
		return DomainBankAccount{}, err
	}

	response := DomainBankAccount{}

	err = convertMarchalBankAccount(responseHttp, &response)
	return response, err
}

func (svc *APIBankAccountService) GetByFilter(filter string) (ResponseBankAccount, error) {
	if filter == "" {
		return ResponseBankAccount{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityBankAccount, "?"+filter)
	if err != nil {
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

func (svc *APIBankAccountService) Put(account DomainBankAccountCreateRequest) (DomainBankAccountCreateResponse, error) {

	payload, err := json.Marshal(account)

	if err != nil {
		return DomainBankAccountCreateResponse{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPutService(EntityBankAccount+"/"+account.ID, payload)
	if err != nil {
		return DomainBankAccountCreateResponse{}, err
	}

	response := DomainBankAccountCreateResponse{}

	err = convertMarchalBankAccount(responseHttp, &response)
	return response, err

}

func (svc *APIBankAccountService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(EntityBankAccount, "/"+id)
	if err != nil {
		return false, err
	}

	return true, err
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseBankAccount(response []byte) (ResponseBankAccount, error) {
	var result ResponseBankAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
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
