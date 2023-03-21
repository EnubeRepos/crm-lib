package blocklistexternal

import "encoding/json"

func (svc *APIBlockListExternalService) Get() (ResponseBlockListExternal, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBlockListExternal, "")
	if err != nil {
		return ResponseBlockListExternal{}, err
	}

	return convertMarchalResponseBlockListExternal(response)
}

func (svc *APIBlockListExternalService) GetById(id string) (DomainBlockListExternal, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBlockListExternal, "/"+id)
	if err != nil {
		return DomainBlockListExternal{}, err
	}

	return convertMarchalBlockListExternal(response)
}

func (svc *APIBlockListExternalService) GetByFilter(filter string) (ResponseBlockListExternal, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBlockListExternal, "?"+filter)
	if err != nil {
		return ResponseBlockListExternal{}, err
	}

	return convertMarchalResponseBlockListExternal(response)
}

func (svc *APIBlockListExternalService) Post(BlockListExternal DomainBlockListExternal) (DomainBlockListExternal, error) {
	payload, err := json.Marshal(BlockListExternal)

	if err != nil {
		return DomainBlockListExternal{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityBlockListExternal, payload)
	if err != nil {
		return DomainBlockListExternal{}, err
	}

	return convertMarchalBlockListExternal(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseBlockListExternal(response []byte) (ResponseBlockListExternal, error) {
	var result ResponseBlockListExternal

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseBlockListExternal{}, err
	}

	return result, nil
}

func convertMarchalBlockListExternal(response []byte) (DomainBlockListExternal, error) {
	var result DomainBlockListExternal

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainBlockListExternal{}, err
	}

	return result, nil
}
