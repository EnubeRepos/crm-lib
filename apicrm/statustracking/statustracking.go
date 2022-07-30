package statustracking

import (
	"encoding/json"
	"fmt"
)

func (svc *APIStatusTrackingService) Get() (ResponseStatusTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityStatusTracking, "")
	if err != nil {
		return ResponseStatusTracking{}, err
	}

	return convertMarchalResponseStatusTracking(response)
}

func (svc *APIStatusTrackingService) GetById(id string) (DomainStatusTracking, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainStatusTracking{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityStatusTracking, "/"+id)
	if err != nil {
		return DomainStatusTracking{}, err
	}

	return convertMarchalStatusTracking(response)
}

func (svc *APIStatusTrackingService) GetByFilter(filter string) (ResponseStatusTracking, error) {
	if filter == "" {
		return ResponseStatusTracking{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityStatusTracking, "?"+filter)
	if err != nil {
		return ResponseStatusTracking{}, err
	}

	return convertMarchalResponseStatusTracking(response)
}

func (svc *APIStatusTrackingService) Post(StatusTracking DomainStatusTracking) (DomainStatusTracking, error) {
	payload, err := json.Marshal(StatusTracking)

	if err != nil {
		return DomainStatusTracking{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityStatusTracking, payload)
	if err != nil {
		return DomainStatusTracking{}, err
	}

	return convertMarchalStatusTracking(response)
}

func (svc *APIStatusTrackingService) Put(Status DomainStatusTrackingBase) (DomainStatusTracking, error) {

	payload, err := json.Marshal(Status)

	if err != nil {
		return DomainStatusTracking{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityStatusTracking+"/"+Status.ID, payload)
	if err != nil {
		return DomainStatusTracking{}, err
	}

	return convertMarchalStatusTracking(response)

}

func (svc *APIStatusTrackingService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityStatusTracking, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseStatusTracking(response []byte) (ResponseStatusTracking, error) {
	var result ResponseStatusTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseStatusTracking{}, err
	}

	return result, nil
}

func convertMarchalStatusTracking(response []byte) (DomainStatusTracking, error) {
	var result DomainStatusTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainStatusTracking{}, err
	}

	return result, nil
}
