package usersecurity

import (
	"encoding/json"
)

func (svc *APIUserService) UpdateStatusMFA(user DomainUserStatusMFA, userID string) (DomainUserStatusMFA, error) {

	payload, err := json.Marshal(user)

	if err != nil {
		return DomainUserStatusMFA{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityUserSecurity+"/"+userID, payload)
	if err != nil {
		return DomainUserStatusMFA{}, err
	}

	return convertMarchalUser(response)
}

func convertMarchalUser(response []byte) (DomainUserStatusMFA, error) {
	var result DomainUserStatusMFA

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainUserStatusMFA{}, err
	}

	return result, nil
}
