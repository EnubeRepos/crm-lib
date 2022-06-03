package userregistration

import "encoding/json"

func (svc *APIUserRegistrationService) Get() (ResponseUserRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUserRegistration, "")
	if err != nil {
		return ResponseUserRegistration{}, err
	}

	return convertMarchalResponseUserRegistration(response)
}

func (svc *APIUserRegistrationService) GetById(id string) (DomainUserRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUserRegistration, "/"+id)
	if err != nil {
		return DomainUserRegistration{}, err
	}

	return convertMarchalUserRegistration(response)
}

func (svc *APIUserRegistrationService) GetByFilter(filter string) (ResponseUserRegistration, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUserRegistration, "?"+filter)
	if err != nil {
		return ResponseUserRegistration{}, err
	}

	return convertMarchalResponseUserRegistration(response)
}

func (svc *APIUserRegistrationService) Post(UserRegistration DomainUserRegistration) (DomainUserRegistration, error) {
	payload, err := json.Marshal(UserRegistration)

	if err != nil {
		return DomainUserRegistration{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityUserRegistration, payload)
	if err != nil {
		return DomainUserRegistration{}, err
	}

	return convertMarchalUserRegistration(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseUserRegistration(response []byte) (ResponseUserRegistration, error) {
	var result ResponseUserRegistration

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseUserRegistration{}, err
	}

	return result, nil
}

func convertMarchalUserRegistration(response []byte) (DomainUserRegistration, error) {
	var result DomainUserRegistration

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainUserRegistration{}, err
	}

	return result, nil
}
