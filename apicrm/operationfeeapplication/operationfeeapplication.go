package operationfeeapplication

import "encoding/json"

func (svc *APITedService) Get() (ResponseOperationFeeApplication, error) {
	response, err := svc.client.CRMHandlerGetService(EntityOperationFeeApplication, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseOperationFeeApplication{}, err
	}

	return convertMarchalResponseOperationFeeApplication(response)
}

func (svc *APITedService) GetById(id string) (DomainOperationFeeApplication, error) {
	response, err := svc.client.CRMHandlerGetService(EntityOperationFeeApplication, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainOperationFeeApplication{}, err
	}

	return convertMarchalTed(response)
}

func (svc *APITedService) GetByFilter(filter string) (ResponseOperationFeeApplication, error) {
	response, err := svc.client.CRMHandlerGetService(EntityOperationFeeApplication, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseOperationFeeApplication{}, err
	}

	return convertMarchalResponseOperationFeeApplication(response)
}

func (svc *APITedService) Post(Ted DomainOperationFeeApplication) (DomainOperationFeeApplication, error) {
	payload, err := json.Marshal(Ted)

	if err != nil {
		return DomainOperationFeeApplication{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityOperationFeeApplication, payload)
	if err != nil {
		return DomainOperationFeeApplication{}, err
	}

	return convertMarchalTed(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseOperationFeeApplication(response []byte) (ResponseOperationFeeApplication, error) {
	var result ResponseOperationFeeApplication

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseOperationFeeApplication{}, err
	}

	return result, nil
}

func convertMarchalTed(response []byte) (DomainOperationFeeApplication, error) {
	var result DomainOperationFeeApplication

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainOperationFeeApplication{}, err
	}

	return result, nil
}
