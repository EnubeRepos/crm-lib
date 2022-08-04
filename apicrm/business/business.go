package business

import (
	"encoding/json"
	"fmt"
)

func (svc *APIRegistrationService) Get() (ResponseBusiness, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBusiness, "")
	if err != nil {
		return ResponseBusiness{}, err
	}

	return convertMarchalResponseBusiness(response)

}

func (svc *APIRegistrationService) GetById(id string) (DomainBusiness, error) {

	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainBusiness{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	response, err := svc.client.CRMHandlerGetService(EntityBusiness, "/"+id)
	if err != nil {
		return DomainBusiness{}, err
	}

	return convertMarchalBusiness(response)
}

func (svc *APIRegistrationService) GetByFilter(filter string) (ResponseBusiness, error) {
	if filter == "" {
		return ResponseBusiness{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityBusiness, "/"+filter)
	if err != nil {
		return ResponseBusiness{}, err
	}

	return convertMarchalResponseBusiness(response)
}

func (svc *APIRegistrationService) Post(request DomainBusiness) (DomainBusiness, error) {
	payload, err := json.Marshal(request)

	if err != nil {
		return DomainBusiness{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityBusiness, payload)
	if err != nil {
		return DomainBusiness{}, err
	}

	return convertMarchalBusiness(response)
}

func (svc *APIRegistrationService) Put(business DomainBusiness) (DomainBusiness, error) {
	payload, err := json.Marshal(business)

	if err != nil {
		return DomainBusiness{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityBusiness+"/"+business.ID, payload)
	if err != nil {
		return DomainBusiness{}, err
	}

	return convertMarchalBusiness(response)
}

func (svc *APIRegistrationService) Delete(id string) (bool, error) {

	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityBusiness, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil

}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalBusiness(response []byte) (DomainBusiness, error) {
	var result DomainBusiness

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainBusiness{}, err
	}

	return result, nil
}

func convertMarchalResponseBusiness(response []byte) (ResponseBusiness, error) {
	var result ResponseBusiness

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseBusiness{}, err
	}

	return result, nil
}
