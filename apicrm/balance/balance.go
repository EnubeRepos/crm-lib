package balance

import (
	"encoding/json"
	"fmt"
)

func (svc *APIBalanceService) Get() (ResponseBalance, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBalance, "")
	if err != nil {
		return ResponseBalance{}, err
	}

	return convertMarchalResponseBalance(response)
}

func (svc *APIBalanceService) GetById(id string) (DomainBalance, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainBalance{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	responseHttp, err := svc.client.CRMHandlerGetService(EntityBalance, "/"+id)
	if err != nil {
		return DomainBalance{}, err
	}

	response := DomainBalance{}

	err = convertMarchalBalance(responseHttp, &response)
	return response, err
}

func (svc *APIBalanceService) GetByFilter(filter string) (ResponseBalance, error) {
	if filter == "" {
		return ResponseBalance{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityBalance, "?"+filter)
	if err != nil {
		return ResponseBalance{}, err
	}

	return convertMarchalResponseBalance(response)
}

func (svc *APIBalanceService) Post(Balance DomainBalanceCreateRequest) (DomainBalanceCreateResponse, error) {
	payload, err := json.Marshal(Balance)

	if err != nil {
		return DomainBalanceCreateResponse{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPostService(EntityBalance, payload)
	if err != nil {
		return DomainBalanceCreateResponse{}, err
	}

	response := DomainBalanceCreateResponse{}

	err = convertMarchalBalance(responseHttp, &response)
	return response, err
}

func (svc *APIBalanceService) Put(balance DomainBalanceCreateRequest) (DomainBalanceCreateResponse, error) {

	payload, err := json.Marshal(balance)

	if err != nil {
		return DomainBalanceCreateResponse{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPutService(EntityBalance+"/"+balance.ID, payload)
	if err != nil {
		return DomainBalanceCreateResponse{}, err
	}

	response := DomainBalanceCreateResponse{}

	err = convertMarchalBalance(responseHttp, &response)
	return response, err
}

func (svc *APIBalanceService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(EntityBalance, "/"+id)
	if err != nil {
		return false, err
	}

	return true, err
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseBalance(response []byte) (ResponseBalance, error) {
	var result ResponseBalance

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseBalance{}, err
	}

	return result, nil
}

func convertMarchalBalance(responseByte []byte, responseObj interface{}) error {
	err := json.Unmarshal(responseByte, responseObj)
	if err != nil {
		return err
	}

	return nil
}
