package pixkey

import "encoding/json"

func (svc *APIPixKeyService) Get() (ResponsePixKey, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPixKey, "")
	if err != nil {
		return ResponsePixKey{}, err
	}

	return convertMarchalResponsePixKey(response)
}

func (svc *APIPixKeyService) GetById(id string) (DomainPixKey, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPixKey, "/"+id)
	if err != nil {
		return DomainPixKey{}, err
	}

	return convertMarchalPixKey(response)
}

func (svc *APIPixKeyService) GetByFilter(filter string) (ResponsePixKey, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPixKey, "?"+filter)
	if err != nil {
		return ResponsePixKey{}, err
	}

	return convertMarchalResponsePixKey(response)
}

func (svc *APIPixKeyService) Post(PixKey DomainPixKey) (DomainPixKey, error) {
	payload, err := json.Marshal(PixKey)

	if err != nil {
		return DomainPixKey{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityPixKey, payload)
	if err != nil {
		return DomainPixKey{}, err
	}

	return convertMarchalPixKey(response)
}

func convertMarchalResponsePixKey(response []byte) (ResponsePixKey, error) {
	var result ResponsePixKey

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponsePixKey{}, err
	}

	return result, nil
}

func convertMarchalPixKey(response []byte) (DomainPixKey, error) {
	var result DomainPixKey

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainPixKey{}, err
	}

	return result, nil
}
