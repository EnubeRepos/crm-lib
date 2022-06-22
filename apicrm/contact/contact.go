package contact

import "encoding/json"

func (svc *APIContactService) Get() (ResponseContact, error) {
	response, err := svc.client.CRMHandlerGetService(EntityContact, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseContact{}, err
	}

	return convertMarchalResponseContact(response)
}

func (svc *APIContactService) GetById(id string) (DomainContact, error) {
	response, err := svc.client.CRMHandlerGetService(EntityContact, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainContact{}, err
	}

	return convertMarchalContact(response)
}

func (svc *APIContactService) GetByFilter(filter string) (ResponseContact, error) {
	response, err := svc.client.CRMHandlerGetService(EntityContact, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseContact{}, err
	}

	return convertMarchalResponseContact(response)
}

func (svc *APIContactService) Post(Contact DomainContact) (DomainContact, error) {
	payload, err := json.Marshal(Contact)

	if err != nil {
		return DomainContact{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityContact, payload)
	if err != nil {
		return DomainContact{}, err
	}

	return convertMarchalContact(response)
}

func (svc *APIContactService) Delete(id string) (bool, error) {
	_, err := svc.client.CRMHandlerDeleteService(EntityContact, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseContact(response []byte) (ResponseContact, error) {
	var result ResponseContact

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseContact{}, err
	}

	return result, err
}

func convertMarchalContact(response []byte) (DomainContact, error) {
	var result DomainContact

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainContact{}, err
	}

	return result, err
}
