package limitmanagement

import "encoding/json"

func (svc *APILimitManagementService) Get() (ResponseLimitManagement, error) {
	response, err := svc.client.CRMHandlerGetService(EntityLimitManagement, "")
	if err != nil {
		return ResponseLimitManagement{}, err
	}

	return convertMarchalResponseLimitManagement(response)
}

func (svc *APILimitManagementService) GetById(id string) (DomainLimitManagement, error) {
	response, err := svc.client.CRMHandlerGetService(EntityLimitManagement, "/"+id)
	if err != nil {
		return DomainLimitManagement{}, err
	}

	return convertMarchalLimitManagement(response)
}

func (svc *APILimitManagementService) GetByFilter(filter string) (ResponseLimitManagement, error) {
	response, err := svc.client.CRMHandlerGetService(EntityLimitManagement, "?"+filter)
	if err != nil {
		return ResponseLimitManagement{}, err
	}

	return convertMarchalResponseLimitManagement(response)
}

func (svc *APILimitManagementService) Post(limitManagement DomainLimitManagement) (DomainLimitManagement, error) {
	payload, err := json.Marshal(limitManagement)

	if err != nil {
		return DomainLimitManagement{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityLimitManagement, payload)
	if err != nil {
		return DomainLimitManagement{}, err
	}

	return convertMarchalLimitManagement(response)
}

func (svc *APILimitManagementService) Put(limitManagement DomainLimitManagement) (DomainLimitManagement, error) {

	payload, err := json.Marshal(limitManagement)

	if err != nil {
		return DomainLimitManagement{}, err
	}

	responseHttp, err := svc.client.CRMHandlerPutService(EntityLimitManagement+"/"+limitManagement.ID, payload)
	if err != nil {
		return DomainLimitManagement{}, err
	}

	return convertMarchalLimitManagement(responseHttp)
}

func convertMarchalResponseLimitManagement(response []byte) (ResponseLimitManagement, error) {
	var result ResponseLimitManagement

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseLimitManagement{}, err
	}

	return result, nil
}

func convertMarchalLimitManagement(response []byte) (DomainLimitManagement, error) {
	var result DomainLimitManagement

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainLimitManagement{}, err
	}

	return result, nil
}
