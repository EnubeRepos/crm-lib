package preregistration

import "encoding/json"

func (svc *APIPRERegistrationService) Get() (ResponseRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityRegistration, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseRegistration{}, err
	}

	return convertMarchalResponseRegistration(response)
}

func (svc *APIPRERegistrationService) GetById(id string) (DomainRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityRegistration, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainRegistration{}, err
	}

	return convertMarchalRegistration(response)
}

func (svc *APIPRERegistrationService) GetByFilter(filter string) (ResponseRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityRegistration, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseRegistration{}, err
	}

	return convertMarchalResponseRegistration(response)
}

func (svc *APIPRERegistrationService) Post(Registration DomainRegistration) (DomainRegistration, error) {
	payload, err := json.Marshal(Registration)

	if err != nil {
		return DomainRegistration{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityRegistration, payload)
	if err != nil {
		return DomainRegistration{}, err
	}

	return convertMarchalRegistration(response)
}

func (svc *APIPRERegistrationService) Put(Registration DomainRegistrationBase) (DomainRegistration, error) {
	payload, err := json.Marshal(Registration)

	if err != nil {
		return DomainRegistration{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityRegistration+"/"+Registration.ID, payload)
	if err != nil {
		return DomainRegistration{}, err
	}

	return convertMarchalRegistration(response)
}

func (svc *APIPRERegistrationService) PutRegistration(Registration DomainRegistration) (DomainRegistration, error) {
	payload, err := json.Marshal(Registration)

	if err != nil {
		return DomainRegistration{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityRegistration+"/"+Registration.ID, payload)
	if err != nil {
		return DomainRegistration{}, err
	}

	//

	return convertMarchalRegistration(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseRegistration(response []byte) (ResponseRegistration, error) {
	var result ResponseRegistration

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseRegistration{}, err
	}

	return result, nil
}

func convertMarchalRegistration(response []byte) (DomainRegistration, error) {
	var result DomainRegistration

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainRegistration{}, err
	}

	return result, nil
}
