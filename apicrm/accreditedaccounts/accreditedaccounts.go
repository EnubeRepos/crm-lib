package accreditedaccounts

import "encoding/json"

func (svc *APIAccreditedAccountsService) Get() (ResponseAccreditedAccounts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccreditedAccounts, "")
	if err != nil {
		return ResponseAccreditedAccounts{}, err
	}

	return convertMarchalResponseAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) GetById(id string) (DomainAccreditedAccounts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccreditedAccounts, "/"+id)
	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	return convertMarchalAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) GetByFilter(filter string) (ResponseAccreditedAccounts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAccreditedAccounts, "?"+filter)
	if err != nil {
		return ResponseAccreditedAccounts{}, err
	}

	return convertMarchalResponseAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) Post(AccreditedAccounts DomainAccreditedAccounts) (DomainAccreditedAccounts, error) {
	payload, err := json.Marshal(AccreditedAccounts)

	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityAccreditedAccounts, payload)
	if err != nil {
		return DomainAccreditedAccounts{}, err
	}

	return convertMarchalAccreditedAccounts(response)
}

func (svc *APIAccreditedAccountsService) Delete(id string) (bool, error) {
	_, err := svc.client.CRMHandlerDeleteService(EntityAccreditedAccounts, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil

}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseAccreditedAccounts(response []byte) (ResponseAccreditedAccounts, error) {
	var result ResponseAccreditedAccounts

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseAccreditedAccounts{}, err
	}

	return result, nil
}

func convertMarchalAccreditedAccounts(response []byte) (DomainAccreditedAccounts, error) {
	var result DomainAccreditedAccounts

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainAccreditedAccounts{}, err
	}

	return result, nil
}
