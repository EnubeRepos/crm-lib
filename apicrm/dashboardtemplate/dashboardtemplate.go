package dashboardtemplate

import "encoding/json"

func (svc *APIDashboardTemplateService) PutDashTemplateForUsers(ModelPut DomainPutDash4Users) (DomainDashboardTemplate, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainDashboardTemplate{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityDashboardTemplate+"/"+ModelPut.ID, payload)
	if err != nil {
		return DomainDashboardTemplate{}, err
	}
	
	return convertMarchalDashboardTemplate(response)
}


func convertMarchalDashboardTemplate(response []byte) (DomainDashboardTemplate, error) {
	var result DomainDashboardTemplate

	err := json.Unmarshal(response, &result)
	if err != nil {
		// ctx.Logger.WithField("Error:", err).Error("Error to make Unmarshal in Distributor")
		return DomainDashboardTemplate{}, err
	}

	return result, nil
}