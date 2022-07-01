package attachment

import (
	"encoding/json"
	"fmt"
)

func (svc *APIAttachmentService) Get() (ResponseAttachment, error) {
	responseHttp, err := svc.client.CRMHandlerGetService(EntityAttachment, "")

	if err != nil {
		return ResponseAttachment{}, err
	}

	return convertMarchalResponse(responseHttp)
}

func (svc *APIAttachmentService) GetById(attachmentId string) (DomainAttachment, error) {
	expectedLen := 17
	idLen := len([]rune(attachmentId))

	if idLen != expectedLen {
		return DomainAttachment{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	responseHttp, err := svc.client.CRMHandlerGetService(EntityAttachment, "/"+attachmentId)

	if err != nil {
		return DomainAttachment{}, err
	}

	return convertMarchalResponseSingle(responseHttp)
}

func (svc *APIAttachmentService) GetByFilter(filter string) (ResponseAttachment, error) {
	if filter == "" {
		return ResponseAttachment{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}

	responseHttp, err := svc.client.CRMHandlerGetService(EntityAttachment, "?"+filter)
	if err != nil {
		return ResponseAttachment{}, err
	}

	return convertMarchalResponse(responseHttp)
}

func (svc *APIAttachmentService) Post(model DomainAttachment) (DomainAttachment, error) {
	payload, err := json.Marshal(model)

	if err != nil {
		return DomainAttachment{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPostService(EntityAttachment, payload)
	if err != nil {
		return DomainAttachment{}, err
	}

	return convertMarchalResponseSingle(responseHttp)
}

func (svc *APIAttachmentService) Put(account DomainAttachment) (DomainAttachment, error) {

	payload, err := json.Marshal(account)

	if err != nil {
		return DomainAttachment{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityAttachment+"/"+account.ID, payload)
	if err != nil {
		return DomainAttachment{}, err
	}

	return convertMarchalResponseSingle(response)

}

func (svc *APIAttachmentService) Delete(id string) (bool, error) {

	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityAttachment, "/"+id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func convertMarchalResponse(response []byte) (ResponseAttachment, error) {
	var result ResponseAttachment

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseAttachment{}, err
	}

	return result, nil
}

func convertMarchalResponseSingle(response []byte) (DomainAttachment, error) {
	var result DomainAttachment

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainAttachment{}, err
	}

	return result, nil
}
