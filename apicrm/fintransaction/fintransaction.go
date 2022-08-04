package fintransaction

import (
	"encoding/json"
	"fmt"
)

func (svc *APIFinTransactionService) Get() (ResponseFinTransaction, error) {
	response, err := svc.client.CRMHandlerGetService(EntityFinTransaction, "")
	if err != nil {
		return ResponseFinTransaction{}, err
	}

	return convertMarchalResponseFinTransaction(response)
}

func (svc *APIFinTransactionService) GetById(id string) (DomainFinTransaction, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainFinTransaction{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityFinTransaction, "/"+id)
	if err != nil {
		return DomainFinTransaction{}, err
	}

	return convertMarchalFinTransaction(response)
}

func (svc *APIFinTransactionService) GetByFilter(filter string) (ResponseFinTransaction, error) {
	if filter == "" {
		return ResponseFinTransaction{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityFinTransaction, "?"+filter)
	if err != nil {
		return ResponseFinTransaction{}, err
	}

	return convertMarchalResponseFinTransaction(response)
}

func (svc *APIFinTransactionService) Post(FinTransaction interface{}) (DomainFinTransaction, error) {
	payload, err := json.Marshal(FinTransaction)

	if err != nil {
		return DomainFinTransaction{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityFinTransaction, payload)
	if err != nil {
		return DomainFinTransaction{}, err
	}

	return convertMarchalFinTransaction(response)
}

func (svc *APIFinTransactionService) Put(Fintransaction DomainFinTransaction) (DomainFinTransaction, error) {
	payload, err := json.Marshal(Fintransaction)

	if err != nil {
		return DomainFinTransaction{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityFinTransaction+"/"+Fintransaction.ID, payload)

  if err != nil {
		return DomainFinTransaction{}, err
	}

	return convertMarchalFinTransaction(response)
}

func (svc *APIFinTransactionService) Delete(id string) (bool, error) {
	_, err := svc.client.CRMHandlerDeleteService(EntityFinTransaction, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseFinTransaction(response []byte) (ResponseFinTransaction, error) {
	var result ResponseFinTransaction

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseFinTransaction{}, err
	}

	return result, nil
}

func convertMarchalFinTransaction(response []byte) (DomainFinTransaction, error) {
	var result DomainFinTransaction

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainFinTransaction{}, err
	}

	return result, nil
}
