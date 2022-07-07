package ted

import "encoding/json"

func (svc *APITedService) Get() (ResponseTed, error) {
	response, err := svc.client.CRMHandlerGetService(EntityTed, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseTed{}, err
	}

	return convertMarchalResponseTed(response)
}

func (svc *APITedService) GetById(id string) (DomainTed, error) {
	response, err := svc.client.CRMHandlerGetService(EntityTed, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainTed{}, err
	}

	return convertMarchalTed(response)
}

func (svc *APITedService) GetByFilter(filter string) (ResponseTed, error) {
	response, err := svc.client.CRMHandlerGetService(EntityTed, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseTed{}, err
	}

	return convertMarchalResponseTed(response)
}

func (svc *APITedService) Post(Ted DomainTed) (DomainTed, error) {
	payload, err := json.Marshal(Ted)

	if err != nil {
		return DomainTed{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityTed, payload)
	if err != nil {
		return DomainTed{}, err
	}

	return convertMarchalTed(response)
}

func (svc *APITedService) PutStatus(ModelPut DomainTedPutConfirm) (DomainTed, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainTed{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityTed+"/"+ModelPut.ID, payload)
	if err != nil {
		return DomainTed{}, err
	}

	return convertMarchalTed(response)
}


// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseTed(response []byte) (ResponseTed, error) {
	var result ResponseTed

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseTed{}, err
	}

	return result, nil
}

func convertMarchalTed(response []byte) (DomainTed, error) {
	var result DomainTed

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainTed{}, err
	}

	return result, nil
}
