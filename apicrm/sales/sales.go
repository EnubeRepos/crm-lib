package sales

import "encoding/json"

func (svc *APISalesService) Get() (ResponseSales, error) {
	response, err := svc.client.CRMHandlerGetService(EntitySales, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseSales{}, err
	}

	return convertMarchalResponseSales(response)
}

func (svc *APISalesService) GetById(id string) (DomainSales, error) {
	response, err := svc.client.CRMHandlerGetService(EntitySales, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainSales{}, err
	}

	return convertMarchalSales(response)
}

func (svc *APISalesService) GetByFilter(filter string) (ResponseSales, error) {
	response, err := svc.client.CRMHandlerGetService(EntitySales, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseSales{}, err
	}

	return convertMarchalResponseSales(response)
}

func (svc *APISalesService) Post(Sales DomainSales) (DomainSales, error) {
	payload, err := json.Marshal(Sales)

	if err != nil {
		return DomainSales{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntitySales, payload)
	if err != nil {
		return DomainSales{}, err
	}

	return convertMarchalSales(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseSales(response []byte) (ResponseSales, error) {
	var result ResponseSales

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseSales{}, err
	}

	return result, err
}

func convertMarchalSales(response []byte) (DomainSales, error) {
	var result DomainSales

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainSales{}, err
	}

	return result, err
}
