package contact

import (
	"encoding/json"
	"fmt"
)

func (svc *APIContactService) Get() (ResponseContact, error) {
	response, err := svc.client.CRMHandlerGetService(EntityContact, "")
	if err != nil {
		return ResponseContact{}, err
	}

	return convertMarchalResponseContact(response)
}

func (svc *APIContactService) GetById(id string) (DomainContact, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainContact{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityContact, "/"+id)
	if err != nil {
		return DomainContact{}, err
	}

	return convertMarchalContact(response)
}

func (svc *APIContactService) GetByFilter(filter string) (ResponseContact, error) {
	if filter == "" {
		return ResponseContact{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityContact, "?"+filter)
	if err != nil {
		return ResponseContact{}, err
	}

	return convertMarchalResponseContact(response)
}

func (svc *APIContactService) Post(Contact DomainContact) (DomainContact, error) {
	payload, err := json.Marshal(Contact)

	if err != nil {
		return DomainContact{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityContact, payload)
	if err != nil {
		return DomainContact{}, err
	}

	return convertMarchalContact(response)
}

func (svc *APIContactService) Put(contact DomainContactBase) (DomainContact, error) {
	payload, err := json.Marshal(contact)

	if err != nil {
		return DomainContact{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityContact+"/"+contact.ID, payload)
	if err != nil {
		return DomainContact{}, err
	}

	return convertMarchalContact(response)
}

func (svc *APIContactService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityContact, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseContact(response []byte) (ResponseContact, error) {
	var result ResponseContact

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseContact{}, err
	}

	return result, err
}

func convertMarchalContact(response []byte) (DomainContact, error) {
	var result DomainContact

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainContact{}, err
	}

	return result, err
}
