package blocklist

import "encoding/json"

func (svc *APIBlockListService) Get() (ResponseBlockList, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBlockList, "")
	if err != nil {
		return ResponseBlockList{}, err
	}

	return convertMarchalResponseBlockList(response)
}

func (svc *APIBlockListService) GetById(id string) (DomainBlockList, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBlockList, "/"+id)
	if err != nil {
		return DomainBlockList{}, err
	}

	return convertMarchalBlockList(response)
}

func (svc *APIBlockListService) GetByFilter(filter string) (ResponseBlockList, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBlockList, "?"+filter)
	if err != nil {
		return ResponseBlockList{}, err
	}

	return convertMarchalResponseBlockList(response)
}

func (svc *APIBlockListService) Post(BlockList DomainBlockList) (DomainBlockList, error) {
	payload, err := json.Marshal(BlockList)

	if err != nil {
		return DomainBlockList{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityBlockList, payload)
	if err != nil {
		return DomainBlockList{}, err
	}

	return convertMarchalBlockList(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseBlockList(response []byte) (ResponseBlockList, error) {
	var result ResponseBlockList

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseBlockList{}, err
	}

	return result, nil
}

func convertMarchalBlockList(response []byte) (DomainBlockList, error) {
	var result DomainBlockList

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainBlockList{}, err
	}

	return result, nil
}
