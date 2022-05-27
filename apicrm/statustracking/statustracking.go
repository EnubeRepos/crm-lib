package statustracking

import "encoding/json"

func (svc *APIStatusTrackingService) Get() (ResponseStatusTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityStatusTracking, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseStatusTracking{}, err
	}

	return convertMarchalResponseStatusTracking(response)
}

func (svc *APIStatusTrackingService) GetById(id string) (DomainStatusTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityStatusTracking, "/"+id)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainStatusTracking{}, err
	}

	return convertMarchalStatusTracking(response)
}

func (svc *APIStatusTrackingService) GetByFilter(filter string) (ResponseStatusTracking, error) {
	response, err := svc.client.CRMHandlerGetService(EntityStatusTracking, "?"+filter)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseStatusTracking{}, err
	}

	return convertMarchalResponseStatusTracking(response)
}

func (svc *APIStatusTrackingService) Post(StatusTracking DomainStatusTracking) (DomainStatusTracking, error) {
	payload, err := json.Marshal(StatusTracking)

	if err != nil {
		return DomainStatusTracking{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityStatusTracking, payload)
	if err != nil {
		return DomainStatusTracking{}, err
	}

	return convertMarchalStatusTracking(response)
}

func (svc *APIStatusTrackingService) Delete() (ResponseStatusTracking, error) {
	response, err := svc.client.CRMHandlerDeleteService(EntityStatusTracking, "")
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseStatusTracking{}, err
	}

	return convertMarchalResponseStatusTracking(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseStatusTracking(response []byte) (ResponseStatusTracking, error) {
	var result ResponseStatusTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return ResponseStatusTracking{}, err
	}

	return result, nil
}

func convertMarchalStatusTracking(response []byte) (DomainStatusTracking, error) {
	var result DomainStatusTracking

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainStatusTracking{}, err
	}

	return result, nil
}
