package card

import (
	"encoding/json"
	"fmt"
)

func (svc *APIVirtualCardService) Get() (ResponseVirtualCard, error) {
	response, err := svc.client.CRMHandlerGetService(EntityVirtualCard, "")
	if err != nil {
		return ResponseVirtualCard{}, err
	}

	return convertMarchalResponseVirtualCard(response)
}

func (svc *APIVirtualCardService) GetById(id string) (DomainVirtualCard, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainVirtualCard{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	responseHttp, err := svc.client.CRMHandlerGetService(EntityVirtualCard, "/"+id)
	if err != nil {
		return DomainVirtualCard{}, err
	}

	response := DomainVirtualCard{}

	err = convertMarchalVirtualCard(responseHttp, &response)
	return response, err
}

func (svc *APIVirtualCardService) GetByFilter(filter string) (ResponseVirtualCard, error) {
	if filter == "" {
		return ResponseVirtualCard{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityVirtualCard, "?"+filter)
	if err != nil {
		return ResponseVirtualCard{}, err
	}

	return convertMarchalResponseVirtualCard(response)
}

func (svc *APIVirtualCardService) Post(debitCard DomainVirtualCard) (DomainVirtualCard, error) {
	payload, err := json.Marshal(debitCard)

	if err != nil {
		return DomainVirtualCard{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPostService(EntityVirtualCard, payload)
	if err != nil {
		return DomainVirtualCard{}, err
	}

	response := DomainVirtualCard{}

	err = convertMarchalVirtualCard(responseHttp, &response)
	return response, err
}

func (svc *APIVirtualCardService) Put(balance DomainVirtualCard) (DomainVirtualCard, error) {

	payload, err := json.Marshal(balance)

	if err != nil {
		return DomainVirtualCard{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPutService(EntityVirtualCard+"/"+balance.ID, payload)
	if err != nil {
		return DomainVirtualCard{}, err
	}

	response := DomainVirtualCard{}

	err = convertMarchalVirtualCard(responseHttp, &response)
	return response, err
}

func (svc *APIVirtualCardService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(EntityVirtualCard, "/"+id)
	if err != nil {
		return false, err
	}

	return true, err
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseVirtualCard(response []byte) (ResponseVirtualCard, error) {
	var result ResponseVirtualCard

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseVirtualCard{}, err
	}

	return result, nil
}

func convertMarchalVirtualCard(responseByte []byte, responseObj interface{}) error {
	err := json.Unmarshal(responseByte, responseObj)
	if err != nil {
		return err
	}

	return nil
}
