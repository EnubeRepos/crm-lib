package attachment

import (
	"encoding/json"
	//"golang.org/x/sys/windows/svc"
)

func (svc *APIAttachmentService) Get() (ResponseAttachment, error) {
	responseHttp, err := svc.client.CRMHandlerGetService(EntityAttachment, "")

	if err != nil {
		return ResponseAttachment{}, err
	}

	return convertMarchalResponse(responseHttp)
}

func (svc *APIAttachmentService) GetById(attachmentId string) (ResponseAttachment, error) {
	responseHttp, err := svc.client.CRMHandlerGetService(EntityAttachment, "/"+attachmentId)

	if err != nil {
		return ResponseAttachment{}, err
	}

	return convertMarchalResponse(responseHttp)
}

func (svc *APIAttachmentService) GetByFilter(filter string) (ResponseAttachment, error) {
	responseHttp, err := svc.client.CRMHandlerGetService(EntityAttachment, "/"+filter)
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

func (svc *APIAttachmentService) Delete() (ResponseAttachment, error) {
	responseHttp, err := svc.client.CRMHandlerDeleteService(EntityAttachment, "")

	if err != nil {
		return ResponseAttachment{}, err
	}

	return convertMarchalResponse(responseHttp)
}

func convertMarchalResponse(response []byte) (ResponseAttachment, error) {
	var result ResponseAttachment

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseAttachment{}, err
	}

	return result, nil
}

func convertMarchalResponseSingle(response []byte) (DomainAttachment, error) {
	var result DomainAttachment

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainAttachment{}, err
	}

	return result, nil
}
