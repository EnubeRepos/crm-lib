package preregistration

import (
	"encoding/json"
	"fmt"
)

func (svc *APIPRERegistrationService) Get() (ResponseRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityRegistration, "")
	if err != nil {
		return ResponseRegistration{}, err
	}

	return convertMarchalResponseRegistration(response)
}

func (svc *APIPRERegistrationService) GetById(id string) (DomainRegistration, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainRegistration{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityRegistration, "/"+id)
	if err != nil {
		return DomainRegistration{}, err
	}

	return convertMarchalRegistration(response)
}

func (svc *APIPRERegistrationService) GetByFilter(filter string) (ResponseRegistration, error) {
	if filter == "" {
		return ResponseRegistration{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityRegistration, "?"+filter)
	if err != nil {
		return ResponseRegistration{}, err
	}

	return convertMarchalResponseRegistration(response)
}

func (svc *APIPRERegistrationService) Post(Registration DomainRegistration) (DomainRegistration, error) {
	payload, err := json.Marshal(Registration)

	if err != nil {
		return DomainRegistration{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityRegistration, payload)
	if err != nil {
		return DomainRegistration{}, err
	}

	return convertMarchalRegistration(response)
}

func (svc *APIPRERegistrationService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityRegistration, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (svc *APIPRERegistrationService) Put(Registration DomainRegistrationBase) (DomainRegistration, error) {
	payload, err := json.Marshal(Registration)

	if err != nil {
		return DomainRegistration{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityRegistration+"/"+Registration.ID, payload)
	if err != nil {
		return DomainRegistration{}, err
	}

	return convertMarchalRegistration(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseRegistration(response []byte) (ResponseRegistration, error) {
	var result ResponseRegistration

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseRegistration{}, err
	}

	return result, nil
}

func convertMarchalRegistration(response []byte) (DomainRegistration, error) {
	var result DomainRegistration

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainRegistration{}, err
	}

	return result, nil
}
