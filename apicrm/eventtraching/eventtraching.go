package eventtraching

import "encoding/json"

func (svc *APIEventTrachingService) Get() (ResponseEventTraching, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTraching, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseEventTraching{}, err
	}

	return convertMarchalResponseEventTraching(response)
}

func (svc *APIEventTrachingService) GetById(id string) (DomainEventTraching, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTraching, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainEventTraching{}, err
	}

	return convertMarchalEventTraching(response)
}

func (svc *APIEventTrachingService) GetByFilter(filter string) (ResponseEventTraching, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTraching, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseEventTraching{}, err
	}

	return convertMarchalResponseEventTraching(response)
}

func (svc *APIEventTrachingService) Post(EventTraching DomainEventTraching) (DomainEventTraching, error) {
	payload, err := json.Marshal(EventTraching)

	if err != nil {
		return DomainEventTraching{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityEventTraching, payload)
	if err != nil {
		return DomainEventTraching{}, err
	}

	return convertMarchalEventTraching(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseEventTraching(response []byte) (ResponseEventTraching, error) {
	var result ResponseEventTraching

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseEventTraching{}, err
	}

	return result, err
}

func convertMarchalEventTraching(response []byte) (DomainEventTraching, error) {
	var result DomainEventTraching

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainEventTraching{}, err
	}

	return result, err
}
