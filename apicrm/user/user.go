package user

import "encoding/json"

func (svc *APIUserService) Get() (ResponseUser, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUser, "")
	if err != nil {
		return ResponseUser{}, err
	}

	return convertMarchalResponseUser(response)
}

func (svc *APIUserService) GetById(id string) (DomainUser, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUser, "/"+id)
	if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}

func (svc *APIUserService) GetByFilter(filter string) (ResponseUser, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUser, "?"+filter)
	if err != nil {
		return ResponseUser{}, err
	}

	return convertMarchalResponseUser(response)
}

func (svc *APIUserService) Post(User DomainUser) (DomainUser, error) {
	payload, err := json.Marshal(User)

	if err != nil {
		return DomainUser{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityUser, payload)
	if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}

func (svc *APIUserService) Delete(id string) (bool, error) {
	_, err := svc.client.CRMHandlerDeleteService(EntityUser, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseUser(response []byte) (ResponseUser, error) {
	var result ResponseUser

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseUser{}, err
	}

	return result, nil
}

func convertMarchalUser(response []byte) (DomainUser, error) {
	var result DomainUser

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainUser{}, err
	}

	return result, nil
}
