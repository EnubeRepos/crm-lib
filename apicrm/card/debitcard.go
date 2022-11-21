package card

import (
	"encoding/json"
	"fmt"
)

func (svc *APIDebitCardService) Get() (ResponseDebitCard, error) {
	response, err := svc.client.CRMHandlerGetService(EntityDebitCard, "")
	if err != nil {
		return ResponseDebitCard{}, err
	}

	return convertMarchalResponseDebitCard(response)
}

func (svc *APIDebitCardService) GetById(id string) (DomainDebitCard, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainDebitCard{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	responseHttp, err := svc.client.CRMHandlerGetService(EntityDebitCard, "/"+id)
	if err != nil {
		return DomainDebitCard{}, err
	}

	response := DomainDebitCard{}

	err = convertMarchalDebitCard(responseHttp, &response)
	return response, err
}

func (svc *APIDebitCardService) GetByFilter(filter string) (ResponseDebitCard, error) {
	if filter == "" {
		return ResponseDebitCard{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityDebitCard, "?"+filter)
	if err != nil {
		return ResponseDebitCard{}, err
	}

	return convertMarchalResponseDebitCard(response)
}

func (svc *APIDebitCardService) Post(debitCard DomainDebitCard) (DomainDebitCard, error) {
	payload, err := json.Marshal(debitCard)

	if err != nil {
		return DomainDebitCard{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPostService(EntityDebitCard, payload)
	if err != nil {
		return DomainDebitCard{}, err
	}

	response := DomainDebitCard{}

	err = convertMarchalDebitCard(responseHttp, &response)
	return response, err
}

func (svc *APIDebitCardService) Put(balance DomainDebitCard) (DomainDebitCard, error) {

	payload, err := json.Marshal(balance)

	if err != nil {
		return DomainDebitCard{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPutService(EntityDebitCard+"/"+balance.ID, payload)
	if err != nil {
		return DomainDebitCard{}, err
	}

	response := DomainDebitCard{}

	err = convertMarchalDebitCard(responseHttp, &response)
	return response, err
}

func (svc *APIDebitCardService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(EntityDebitCard, "/"+id)
	if err != nil {
		return false, err
	}

	return true, err
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseDebitCard(response []byte) (ResponseDebitCard, error) {
	var result ResponseDebitCard

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseDebitCard{}, err
	}

	return result, nil
}

func convertMarchalDebitCard(responseByte []byte, responseObj interface{}) error {
	err := json.Unmarshal(responseByte, responseObj)
	if err != nil {
		return err
	}

	return nil
}
