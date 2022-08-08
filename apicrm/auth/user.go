package user

import "encoding/json"

func (svc *APIAuthService) Get() (DomainAuth, error) {
	response, err := svc.client.CRMHandlerGetService(EntityAuth, "")
	if err != nil {
		return DomainAuth{}, err
	}

	return convertMarchalDomainAuth(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalDomainAuth(response []byte) (DomainAuth, error) {
	var result DomainAuth

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainAuth{}, err
	}

	return result, nil
}
