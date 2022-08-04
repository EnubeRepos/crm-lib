package commissions

import (
	"encoding/json"
	"fmt"
)

func (svc *APICommissionsService) Get() (ResponseCommissions, error) {
	response, err := svc.client.CRMHandlerGetService(EntityCommissions, "")
	if err != nil {
		return ResponseCommissions{}, err
	}

	return convertMarchalResponseCommissions(response)
}

func (svc *APICommissionsService) GetById(id string) (DomainCommissions, error) {

	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainCommissions{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	response, err := svc.client.CRMHandlerGetService(EntityCommissions, "/"+id)
	if err != nil {
		return DomainCommissions{}, err
	}

	return convertMarchalCommissions(response)
}

func (svc *APICommissionsService) GetByFilter(filter string) (ResponseCommissions, error) {
	if filter == "" {
		return ResponseCommissions{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityCommissions, "?"+filter)
	if err != nil {
		return ResponseCommissions{}, err
	}

	return convertMarchalResponseCommissions(response)
}

func (svc *APICommissionsService) Post(Commissions DomainCommissions) (DomainCommissions, error) {
	payload, err := json.Marshal(Commissions)

	if err != nil {
		return DomainCommissions{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityCommissions, payload)
	if err != nil {
		return DomainCommissions{}, err
	}

	return convertMarchalCommissions(response)
}

func (svc *APICommissionsService) PutStatus(ModelPutStatus DomainComissionsPutStatus) (DomainCommissions, error) {
	payload, err := json.Marshal(ModelPutStatus)

	if err != nil {
		return DomainCommissions{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityCommissions+"/"+ModelPutStatus.ID, payload)

  if err != nil {
		return DomainCommissions{}, err
	}

	return convertMarchalCommissions(response)
}

func (svc *APICommissionsService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityCommissions, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseCommissions(response []byte) (ResponseCommissions, error) {
	var result ResponseCommissions

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseCommissions{}, err
	}

	return result, nil
}

func convertMarchalCommissions(response []byte) (DomainCommissions, error) {
	var result DomainCommissions

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainCommissions{}, err
	}

	return result, nil
}
