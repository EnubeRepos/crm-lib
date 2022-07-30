package personalcontact

import "encoding/json"

func (svc *APIPersonalContactService) Get() (ResponsePersonalContacts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPersonalContacts, "")
	if err != nil {
		return ResponsePersonalContacts{}, err
	}

	return convertMarchalResponsePersonalContact(response)
}

func (svc *APIPersonalContactService) GetById(id string) (DomainPersonalContacts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPersonalContacts, "/"+id)
	if err != nil {
		return DomainPersonalContacts{}, err
	}

	return convertMarchalPersonalContact(response)
}

func (svc *APIPersonalContactService) GetByFilter(filter string) (ResponsePersonalContacts, error) {
	response, err := svc.client.CRMHandlerGetService(EntityPersonalContacts, "?"+filter)
	if err != nil {
		return ResponsePersonalContacts{}, err
	}

	return convertMarchalResponsePersonalContact(response)
}

func (svc *APIPersonalContactService) Post(PersonalContact DomainPersonalContacts) (DomainPersonalContacts, error) {
	payload, err := json.Marshal(PersonalContact)

	if err != nil {
		return DomainPersonalContacts{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityPersonalContacts, payload)
	if err != nil {
		return DomainPersonalContacts{}, err
	}

	return convertMarchalPersonalContact(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponsePersonalContact(response []byte) (ResponsePersonalContacts, error) {
	var result ResponsePersonalContacts

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponsePersonalContacts{}, err
	}

	return result, err
}

func convertMarchalPersonalContact(response []byte) (DomainPersonalContacts, error) {
	var result DomainPersonalContacts

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainPersonalContacts{}, err
	}

	return result, err
}
