package parcels

import (
	"encoding/json"
	"fmt"
)

func (svc *APIParcelsService) Get() (ResponseParcels, error) {
	response, err := svc.client.CRMHandlerGetService(EntityParcels, "")
	if err != nil {
		return ResponseParcels{}, err
	}

	return convertMarchalResponseParcels(response)
}

func (svc *APIParcelsService) GetById(id string) (DomainParcels, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainParcels{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityParcels, "/"+id)
	if err != nil {
		return DomainParcels{}, err
	}

	return convertMarchalParcels(response)
}

func (svc *APIParcelsService) GetByFilter(filter string) (ResponseParcels, error) {
	if filter == "" {
		return ResponseParcels{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityParcels, "?"+filter)
	if err != nil {
		return ResponseParcels{}, err
	}

	return convertMarchalResponseParcels(response)
}

func (svc *APIParcelsService) Post(Parcels DomainParcels) (DomainParcels, error) {
	payload, err := json.Marshal(Parcels)

	if err != nil {
		return DomainParcels{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityParcels, payload)
	if err != nil {
		return DomainParcels{}, err
	}

	return convertMarchalParcels(response)
}

func (svc *APIParcelsService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityParcels, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (svc *APIParcelsService) PutAuthCode(ModelPutAuthCode DomainParcelsPutAuthCode) (DomainParcels, error) {
	payload, err := json.Marshal(ModelPutAuthCode)

	if err != nil {
		return DomainParcels{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityParcels+"/"+ModelPutAuthCode.ID, payload)
	if err != nil {
		return DomainParcels{}, err
	}

	return convertMarchalParcels(response)
}

func (svc *APIParcelsService) PutStatus(ModelPutStatus DomainParcelsPutStatus) (DomainParcels, error) {
	payload, err := json.Marshal(ModelPutStatus)

	if err != nil {
		return DomainParcels{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityParcels+"/"+ModelPutStatus.ID, payload)
	if err != nil {
		return DomainParcels{}, err
	}

	return convertMarchalParcels(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseParcels(response []byte) (ResponseParcels, error) {
	var result ResponseParcels

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseParcels{}, err
	}

	return result, nil
}

func convertMarchalParcels(response []byte) (DomainParcels, error) {
	var result DomainParcels

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainParcels{}, err
	}

	return result, nil
}
