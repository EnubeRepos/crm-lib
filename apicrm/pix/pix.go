package pix

import "encoding/json"

func (svc *APIPixService) Get() (ResponsePix, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPix, "")
	if err != nil {
		return ResponsePix{}, err
	}

	return convertMarchalResponsePix(response)
}

func (svc *APIPixService) GetById(id string) (DomainPix, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPix, "/"+id)
	if err != nil {
		return DomainPix{}, err
	}

	return convertMarchalPix(response)
}

func (svc *APIPixService) GetByFilter(filter string) (ResponsePix, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPix, "?"+filter)
	if err != nil {
		return ResponsePix{}, err
	}

	return convertMarchalResponsePix(response)
}

func (svc *APIPixService) Post(Pix DomainPix) (DomainPix, error) {
	payload, err := json.Marshal(Pix)

	if err != nil {
		return DomainPix{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityPix, payload)
	if err != nil {
		return DomainPix{}, err
	}

	return convertMarchalPix(response)
}

func (svc *APIPixService) PutStatus(ModelPut DomainPixPutStatus) (DomainPix, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainPix{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityPix+"/"+ModelPut.ID, payload)
	if err != nil {
		return DomainPix{}, err
	}

	return convertMarchalPix(response)
}

func (svc *APIPixService) PutPixKeyInfo(ModelPut DomainPixPutPixKeyInfo) (DomainPix, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainPix{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityPix+"/"+ModelPut.ID, payload)
	if err != nil {
		return DomainPix{}, err
	}

	return convertMarchalPix(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponsePix(response []byte) (ResponsePix, error) {
	var result ResponsePix

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponsePix{}, err
	}

	return result, nil
}

func convertMarchalPix(response []byte) (DomainPix, error) {
	var result DomainPix

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainPix{}, err
	}

	return result, nil
}
