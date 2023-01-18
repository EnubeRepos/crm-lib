package occupation

import (
	"encoding/json"
	"fmt"
)

func (svc *APIOccupationService) Get() (ResponseOccupation, error) {
	response, err := svc.client.CRMHandlerGetService(EntityOccupation, "")
	if err != nil {
		return ResponseOccupation{}, err
	}

	return convertMarchalResponseOccupation(response)
}

func (svc *APIOccupationService) GetById(id string) (DomainOccupation, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainOccupation{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityOccupation, "/"+id)
	if err != nil {
		return DomainOccupation{}, err
	}

	return convertMarchalOccupation(response)
}

func (svc *APIOccupationService) GetByFilter(filter string) (ResponseOccupation, error) {
	if filter == "" {
		return ResponseOccupation{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityOccupation, "?"+filter)
	if err != nil {
		return ResponseOccupation{}, err
	}

	return convertMarchalResponseOccupation(response)
}

func (svc *APIOccupationService) Post(Occupation DomainOccupation) (DomainOccupation, error) {
	payload, err := json.Marshal(Occupation)

	if err != nil {
		return DomainOccupation{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityOccupation, payload)
	if err != nil {
		return DomainOccupation{}, err
	}

	return convertMarchalOccupation(response)
}

func (svc *APIOccupationService) Put(Status DomainOccupationBase) (DomainOccupation, error) {

	payload, err := json.Marshal(Status)

	if err != nil {
		return DomainOccupation{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityOccupation+"/"+Status.ID, payload)
	if err != nil {
		return DomainOccupation{}, err
	}

	return convertMarchalOccupation(response)

}

func (svc *APIOccupationService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityOccupation, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseOccupation(response []byte) (ResponseOccupation, error) {
	var result ResponseOccupation

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseOccupation{}, err
	}

	return result, nil
}

func convertMarchalOccupation(response []byte) (DomainOccupation, error) {
	var result DomainOccupation

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainOccupation{}, err
	}

	return result, nil
}
