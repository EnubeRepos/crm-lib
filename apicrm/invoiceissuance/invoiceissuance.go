package invoiceissuance

import (
	"encoding/json"
	"fmt"
)

func (svc *APIInvoiceIssuanceService) Get() (ResponseInvoiceIssuance, error) {
	response, err := svc.client.CRMHandlerGetService(EntityInvoiceIssuance, "")
	if err != nil {
		return ResponseInvoiceIssuance{}, err
	}

	return convertMarchalResponseInvoiceIssuance(response)
}

func (svc *APIInvoiceIssuanceService) GetById(id string) (DomainInvoiceIssuance, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainInvoiceIssuance{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityInvoiceIssuance, "/"+id)
	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	return convertMarchalInvoiceIssuance(response)
}

func (svc *APIInvoiceIssuanceService) GetByFilter(filter string) (ResponseInvoiceIssuance, error) {
	if filter == "" {
		return ResponseInvoiceIssuance{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityInvoiceIssuance, "?"+filter)
	if err != nil {
		return ResponseInvoiceIssuance{}, err
	}

	return convertMarchalResponseInvoiceIssuance(response)
}

func (svc *APIInvoiceIssuanceService) Post(InvoiceIssuance DomainInvoiceIssuance) (DomainInvoiceIssuance, error) {
	payload, err := json.Marshal(InvoiceIssuance)

	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityInvoiceIssuance, payload)
	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	return convertMarchalInvoiceIssuance(response)
}

func (svc *APIInvoiceIssuanceService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityInvoiceIssuance, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (svc *APIInvoiceIssuanceService) PutAuthCode(ModelPutAuthCode DomainInvoiceIssuancePutAuthCode) (DomainInvoiceIssuance, error) {
	payload, err := json.Marshal(ModelPutAuthCode)

	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityInvoiceIssuance+"/"+ModelPutAuthCode.ID, payload)
	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	return convertMarchalInvoiceIssuance(response)
}

func (svc *APIInvoiceIssuanceService) PutStatus(ModelPutStatus DomainInvoiceIssuancePutStatus) (DomainInvoiceIssuance, error) {
	payload, err := json.Marshal(ModelPutStatus)

	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityInvoiceIssuance+"/"+ModelPutStatus.ID, payload)
	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	return convertMarchalInvoiceIssuance(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseInvoiceIssuance(response []byte) (ResponseInvoiceIssuance, error) {
	var result ResponseInvoiceIssuance

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseInvoiceIssuance{}, err
	}

	return result, nil
}

func convertMarchalInvoiceIssuance(response []byte) (DomainInvoiceIssuance, error) {
	var result DomainInvoiceIssuance

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainInvoiceIssuance{}, err
	}

	return result, nil
}
