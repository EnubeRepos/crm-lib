package account

import (
	"encoding/json"
	"fmt"
)

func (svc *APIAccountService) Get() (ResponseAccount, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccount, "")
	if err != nil {
		return ResponseAccount{}, err
	}

	return convertMarchalResponseAccount(response)
}

func (svc *APIAccountService) GetById(id string) (DomainAccount, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainAccount{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	response, err := svc.client.CRMHandlerGetService(EntityAccount, "/"+id)
	if err != nil {
		return DomainAccount{}, err
	}

	return convertMarchalAccount(response)
}

func (svc *APIAccountService) GetByFilter(filter string) (ResponseAccount, error) {
	if filter == "" {
		return ResponseAccount{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}

	response, err := svc.client.CRMHandlerGetService(EntityAccount, "?"+filter)
	if err != nil {
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

func (svc *APIAccountService) Put(account DomainAccountBase) (DomainAccount, error) {

	payload, err := json.Marshal(account)

	if err != nil {
		return DomainAccount{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityAccount+"/"+account.ID, payload)
	if err != nil {
		return DomainAccount{}, err
	}

	return convertMarchalAccount(response)

}

func (svc *APIAccountService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(EntityAccount, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseAccount(response []byte) (ResponseAccount, error) {
	var result ResponseAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseAccount{}, nil
	}

	return result, err
}

func convertMarchalAccount(response []byte) (DomainAccount, error) {
	var result DomainAccount

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainAccount{}, err
	}

	return result, nil
}
