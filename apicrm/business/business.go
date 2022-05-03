package business

import "encoding/json"

func (svc *APIRegistrationService) GetById(id string) (DomainBusiness, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBusiness, "/"+id)
	if err != nil {
		return DomainBusiness{}, err
	}

	return convertMarchalBusiness(response)
}

func (svc *APIRegistrationService) Post(request DomainBusiness) (DomainBusiness, error) {
	payload, err := json.Marshal(request)

	if err != nil {
		return DomainBusiness{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityBusiness, payload)
	if err != nil {
		return DomainBusiness{}, err
	}

	return convertMarchalBusiness(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalBusiness(response []byte) (DomainBusiness, error) {
	var result DomainBusiness

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainBusiness{}, err
	}

	return result, nil
}
