package generic

import (
	"encoding/json"
	"fmt"
)

func (svc *APIGenericService) Get() ([]byte, error) {
	return svc.client.CRMHandlerGetService(svc.entityName, "")
}

func (svc *APIGenericService) GetById(id string) ([]byte, error) {
	return svc.client.CRMHandlerGetService(svc.entityName, "/"+id)
}

func (svc *APIGenericService) GetByFilter(filter string) ([]byte, error) {
	if filter == "" {
		return nil, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}

	return svc.client.CRMHandlerGetService(svc.entityName, "?"+filter)
}

func (svc *APIGenericService) Post(model interface{}) ([]byte, error) {
	payload, err := json.Marshal(model)

	if err != nil {
		return nil, err
	}

	return svc.client.CRMHandlerPostService(svc.entityName, payload)
}

func (svc *APIGenericService) Put(id string, model interface{}) ([]byte, error) {
	payload, err := json.Marshal(model)

	if err != nil {
		return nil, err
	}

	return svc.client.CRMHandlerPutService(svc.entityName+"/"+id, payload)
}

func (svc *APIGenericService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}

	_, err := svc.client.CRMHandlerDeleteService(svc.entityName, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}
