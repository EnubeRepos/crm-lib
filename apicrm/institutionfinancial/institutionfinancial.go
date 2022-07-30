package institutionfinancial

import "encoding/json"

func (svc *APIInstitutionFinancialService) Get() (ResponseInstitutionFinancial, error) {
	response, err := svc.client.CRMHandlerGetService(EntityInstitutionFinancial, "")
	if err != nil {
		return ResponseInstitutionFinancial{}, err
	}

	return convertMarchalResponseInstitutionFinancial(response)
}

func (svc *APIInstitutionFinancialService) GetById(id string) (DomainInstitutionFinancial, error) {
	response, err := svc.client.CRMHandlerGetService(EntityInstitutionFinancial, "/"+id)
	if err != nil {
		return DomainInstitutionFinancial{}, err
	}

	return convertMarchalInstitutionFinancial(response)
}

func (svc *APIInstitutionFinancialService) GetByFilter(filter string) (ResponseInstitutionFinancial, error) {
	response, err := svc.client.CRMHandlerGetService(EntityInstitutionFinancial, "?"+filter)
	if err != nil {
		return ResponseInstitutionFinancial{}, err
	}

	return convertMarchalResponseInstitutionFinancial(response)
}

func (svc *APIInstitutionFinancialService) Post(InstitutionFinancial DomainInstitutionFinancial) (DomainInstitutionFinancial, error) {
	payload, err := json.Marshal(InstitutionFinancial)

	if err != nil {
		return DomainInstitutionFinancial{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityInstitutionFinancial, payload)
	if err != nil {
		return DomainInstitutionFinancial{}, err
	}

	return convertMarchalInstitutionFinancial(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseInstitutionFinancial(response []byte) (ResponseInstitutionFinancial, error) {
	var result ResponseInstitutionFinancial

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseInstitutionFinancial{}, err
	}

	return result, nil
}

func convertMarchalInstitutionFinancial(response []byte) (DomainInstitutionFinancial, error) {
	var result DomainInstitutionFinancial

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainInstitutionFinancial{}, err
	}

	return result, nil
}
