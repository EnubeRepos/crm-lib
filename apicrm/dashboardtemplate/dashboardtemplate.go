package dashboardtemplate

import "encoding/json"

func (svc *APIDashboardTemplateService) PutDashTemplateForUsers(ModelPut DomainPutDash4Users) (bool, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return false, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityDashboardTemplate+"/"+ModelPut.ID, payload)
	if err != nil {
		return false, err
	}

	var result bool

	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}

	return result == true, nil
}