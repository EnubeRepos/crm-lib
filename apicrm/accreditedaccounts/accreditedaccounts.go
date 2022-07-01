package accreditedaccounts

import (
	"encoding/json"
	"fmt"
)

func (svc *APIAccreditedAccountsService) Get() (ResponseAccreditedAccounts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccreditedAccounts, "")
	if err != nil {
		return ResponseAccreditedAccounts{}, err
	}

	return convertMarchalResponseAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) GetById(id string) (DomainAccreditedAccounts, error) {

	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainAccreditedAccounts{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	response, err := svc.client.CRMHandlerGetService(EntityAccreditedAccounts, "/"+id)
	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	return convertMarchalAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) GetByFilter(filter string) (ResponseAccreditedAccounts, error) {
	if filter == "" {
		return ResponseAccreditedAccounts{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityAccreditedAccounts, "?"+filter)
	if err != nil {
		return ResponseAccreditedAccounts{}, err
	}

	return convertMarchalResponseAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) Post(AccreditedAccounts DomainAccreditedAccounts) (DomainAccreditedAccounts, error) {
	payload, err := json.Marshal(AccreditedAccounts)

	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityAccreditedAccounts, payload)
	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	return convertMarchalAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) Put(account DomainAccreditedAccounts) (DomainAccreditedAccounts, error) {

	payload, err := json.Marshal(account)

	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityAccreditedAccounts+"/"+account.ID, payload)
	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	return convertMarchalAccreditedAccounts(response)

}

func (svc *APIAccreditedAccountsService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityAccreditedAccounts, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil

}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseAccreditedAccounts(response []byte) (ResponseAccreditedAccounts, error) {
	var result ResponseAccreditedAccounts

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseAccreditedAccounts{}, err
	}

	return result, nil
}

func convertMarchalAccreditedAccounts(response []byte) (DomainAccreditedAccounts, error) {
	var result DomainAccreditedAccounts

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	return result, nil
}
