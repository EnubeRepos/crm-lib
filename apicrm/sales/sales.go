package sales

import (
	"encoding/json"
	"fmt"
)

func (svc *APISalesService) Get() (ResponseSales, error) {
	response, err := svc.client.CRMHandlerGetService(EntitySales, "")
	if err != nil {
		return ResponseSales{}, err
	}

	return convertMarchalResponseSales(response)
}

func (svc *APISalesService) GetById(id string) (DomainSales, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainSales{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntitySales, "/"+id)
	if err != nil {
		return DomainSales{}, err
	}

	return convertMarchalSales(response)
}

func (svc *APISalesService) GetByFilter(filter string) (ResponseSales, error) {
	if filter == "" {
		return ResponseSales{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntitySales, "?"+filter)
	if err != nil {
		return ResponseSales{}, err
	}

	return convertMarchalResponseSales(response)
}

func (svc *APISalesService) Post(Sales DomainSales) (DomainSales, error) {
	payload, err := json.Marshal(Sales)

	if err != nil {
		return DomainSales{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntitySales, payload)
	if err != nil {
		return DomainSales{}, err
	}

	return convertMarchalSales(response)
}

func (svc *APISalesService) Put(sale DomainSalesBase) (DomainSales, error) {

	payload, err := json.Marshal(sale)

	if err != nil {
		return DomainSales{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntitySales+"/"+sale.ID, payload)
	if err != nil {
		return DomainSales{}, err
	}

	return convertMarchalSales(response)

}

func (svc *APISalesService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntitySales, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseSales(response []byte) (ResponseSales, error) {
	var result ResponseSales

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseSales{}, err
	}

	return result, nil
}

func convertMarchalSales(response []byte) (DomainSales, error) {
	var result DomainSales

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainSales{}, err
	}

	return result, nil
}
