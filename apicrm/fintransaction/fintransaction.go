package fintransaction

import "encoding/json"

func (svc *APIFinTransactionService) Get() (ResponseFinTransaction, error) {
	response, err := svc.client.CRMHandlerGetService(EntityFinTransaction, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseFinTransaction{}, err
	}

	return convertMarchalResponseFinTransaction(response)
}

func (svc *APIFinTransactionService) GetById(id string) (DomainFinTransaction, error) {
	response, err := svc.client.CRMHandlerGetService(EntityFinTransaction, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainFinTransaction{}, err
	}

	return convertMarchalFinTransaction(response)
}

func (svc *APIFinTransactionService) GetByFilter(filter string) (ResponseFinTransaction, error) {
	response, err := svc.client.CRMHandlerGetService(EntityFinTransaction, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseFinTransaction{}, err
	}

	return convertMarchalResponseFinTransaction(response)
}

func (svc *APIFinTransactionService) Post(FinTransaction interface{}) (DomainFinTransaction, error) {
	payload, err := json.Marshal(FinTransaction)

	if err != nil {
		return DomainFinTransaction{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityFinTransaction, payload)
	if err != nil {
		return DomainFinTransaction{}, err
	}

	return convertMarchalFinTransaction(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseFinTransaction(response []byte) (ResponseFinTransaction, error) {
	var result ResponseFinTransaction

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseFinTransaction{}, err
	}

	return result, nil
}

func convertMarchalFinTransaction(response []byte) (DomainFinTransaction, error) {
	var result DomainFinTransaction

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainFinTransaction{}, err
	}

	return result, nil
}
