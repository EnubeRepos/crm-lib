package dashboardtemplate

import "encoding/json"

func (svc *APIDashboardTemplateService) PostDeployToUsers(ModelPut DomainPutDash4Users) (DomainDashboardTemplate, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainDashboardTemplate{}, err
	}

	_, err = svc.client.CRMHandlerPostService(EntityDashboardTemplateDeplorToUsers, payload)
	if err != nil {
		return DomainDashboardTemplate{}, err
	}

	response,_ := svc.client.CRMHandlerGetService(EntityDashboardTemplate, ModelPut.ID)

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