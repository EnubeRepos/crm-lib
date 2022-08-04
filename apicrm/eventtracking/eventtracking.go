package eventtracking

import (
	"encoding/json"
	"fmt"
)

func (svc *APIEventTrackingService) Get() (ResponseEventTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTracking, "")
	if err != nil {
		return ResponseEventTracking{}, err
	}

	return convertMarchalResponseEventTracking(response)
}

func (svc *APIEventTrackingService) GetById(id string) (DomainEventTracking, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainEventTracking{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityEventTracking, "/"+id)
	if err != nil {
		return DomainEventTracking{}, err
	}

	return convertMarchalEventTracking(response)
}

func (svc *APIEventTrackingService) GetByFilter(filter string) (ResponseEventTracking, error) {
	if filter == "" {
		return ResponseEventTracking{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityEventTracking, "?"+filter)
	if err != nil {
		return ResponseEventTracking{}, err
	}

	return convertMarchalResponseEventTracking(response)
}

func (svc *APIEventTrackingService) Post(EventTracking DomainEventTracking) (DomainEventTracking, error) {
	payload, err := json.Marshal(EventTracking)

	if err != nil {
		return DomainEventTracking{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityEventTracking, payload)
	if err != nil {
		return DomainEventTracking{}, err
	}

	return convertMarchalEventTracking(response)
}

func (svc *APIEventTrackingService) Put(event DomainEventTrackingBase) (DomainEventTracking, error) {
	payload, err := json.Marshal(event)

	if err != nil {
		return DomainEventTracking{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityEventTracking+"/"+event.ID, payload)
	if err != nil {
		return DomainEventTracking{}, err
	}

	return convertMarchalEventTracking(response)
}

func (svc *APIEventTrackingService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityEventTracking, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseEventTracking(response []byte) (ResponseEventTracking, error) {
	var result ResponseEventTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseEventTracking{}, err
	}

	return result, nil
}

func convertMarchalEventTracking(response []byte) (DomainEventTracking, error) {
	var result DomainEventTracking

	err := json.Unmarshal(response, &result)
	if err != nil {

		return DomainEventTracking{}, err
	}

	return result, nil
}
