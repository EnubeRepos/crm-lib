package chargeback

import "encoding/json"

func (svc *APIChargebackService) Get() (ResponseChargeback, error) {
	response, err := svc.client.CRMHandlerGetService(EntityChargeback, "")
	if err != nil {
		return ResponseChargeback{}, err
	}

	return convertMarchalResponseChargeback(response)
}

func (svc *APIChargebackService) GetById(id string) (DomainChargeback, error) {
	response, err := svc.client.CRMHandlerGetService(EntityChargeback, "/"+id)
	if err != nil {
		return DomainChargeback{}, err
	}

	return convertMarchalChargeback(response)
}

func (svc *APIChargebackService) GetByFilter(filter string) (ResponseChargeback, error) {
	response, err := svc.client.CRMHandlerGetService(EntityChargeback, "?"+filter)
	if err != nil {
		return ResponseChargeback{}, err
	}

	return convertMarchalResponseChargeback(response)
}

func (svc *APIChargebackService) Post(Chargeback DomainChargeback) (DomainChargeback, error) {
	payload, err := json.Marshal(Chargeback)

	if err != nil {
		return DomainChargeback{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityChargeback, payload)
	if err != nil {
		return DomainChargeback{}, err
	}

	return convertMarchalChargeback(response)
}

func (svc *APIChargebackService) PutStatus(ModelPut DomainChargebackPutStatus) (DomainChargeback, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainChargeback{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityChargeback+"/"+ModelPut.ID, payload)
	if err != nil {
		return DomainChargeback{}, err
	}

	return convertMarchalChargeback(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseChargeback(response []byte) (ResponseChargeback, error) {
	var result ResponseChargeback

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseChargeback{}, err
	}

	return result, nil
}

func convertMarchalChargeback(response []byte) (DomainChargeback, error) {
	var result DomainChargeback

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainChargeback{}, err
	}

	return result, nil
}
