package eventtracking

import "encoding/json"

func (svc *APIEventTrackingService) Get() (ResponseEventTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTracking, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseEventTracking{}, err
	}

	return convertMarchalResponseEventTracking(response)
}

func (svc *APIEventTrackingService) GetById(id string) (DomainEventTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTracking, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainEventTracking{}, err
	}

	return convertMarchalEventTracking(response)
}

func (svc *APIEventTrackingService) GetByFilter(filter string) (ResponseEventTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityEventTracking, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseEventTracking{}, err
	}

	return convertMarchalResponseEventTracking(response)
}

func (svc *APIEventTrackingService) Post(EventTracking DomainEventTracking) (DomainEventTracking, error) {
	payload, err := json.Marshal(EventTracking)

	if err != nil {
		return DomainEventTracking{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityEventTracking, payload)
	if err != nil {
		return DomainEventTracking{}, err
	}

	return convertMarchalEventTracking(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseEventTracking(response []byte) (ResponseEventTracking, error) {
	var result ResponseEventTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseEventTracking{}, err
	}

	return result, err
}

func convertMarchalEventTracking(response []byte) (DomainEventTracking, error) {
	var result DomainEventTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainEventTracking{}, err
	}

	return result, err
}
