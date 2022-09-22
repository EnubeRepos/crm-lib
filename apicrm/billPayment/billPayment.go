package billPayment

import "encoding/json"

func (svc *APIBillPaymentService) Get() (ResponseBillPayment, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBillPayment, "")
	if err != nil {
		return ResponseBillPayment{}, err
	}

	return convertMarchalResponseBillPayment(response)
}

func (svc *APIBillPaymentService) GetById(id string) (DomainBillPayment, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBillPayment, "/"+id)
	if err != nil {
		return DomainBillPayment{}, err
	}

	return convertMarchalBillPayment(response)
}

func (svc *APIBillPaymentService) GetByFilter(filter string) (ResponseBillPayment, error) {
	response, err := svc.client.CRMHandlerGetService(EntityBillPayment, "?"+filter)
	if err != nil {
		return ResponseBillPayment{}, err
	}

	return convertMarchalResponseBillPayment(response)
}

func (svc *APIBillPaymentService) Post(BillPayment DomainBillPayment) (DomainBillPayment, error) {
	payload, err := json.Marshal(BillPayment)

	if err != nil {
		return DomainBillPayment{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityBillPayment, payload)
	if err != nil {
		return DomainBillPayment{}, err
	}

	return convertMarchalBillPayment(response)
}


func (svc *APIBillPaymentService) PutStatus(ModelPut DomainBillPaymentPutStatus) (DomainBillPayment, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainBillPayment{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityBillPayment+"/"+ModelPut.ID, payload)
	if err != nil {
		return DomainBillPayment{}, err
	}

	return convertMarchalBillPayment(response)
}

func convertMarchalResponseBillPayment(response []byte) (ResponseBillPayment, error) {
	var result ResponseBillPayment

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseBillPayment{}, err
	}

	return result, nil
}

func convertMarchalBillPayment(response []byte) (DomainBillPayment, error) {
	var result DomainBillPayment

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainBillPayment{}, err
	}

	return result, nil
}
