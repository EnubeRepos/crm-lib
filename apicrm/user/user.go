package user

import (
	"encoding/json"
	"fmt"
)

func (svc *APIUserService) Get() (ResponseUser, error) {
	response, err := svc.client.CRMHandlerGetService(EntityUser, "")
	if err != nil {
		return ResponseUser{}, err
	}

	return convertMarchalResponseUser(response)
}

func (svc *APIUserService) GetById(id string) (DomainUser, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return DomainUser{}, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	response, err := svc.client.CRMHandlerGetService(EntityUser, "/"+id)
	if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}

func (svc *APIUserService) GetByFilter(filter string) (ResponseUser, error) {
	if filter == "" {
		return ResponseUser{}, fmt.Errorf("error: invalid filter, filter cannot be empty")

	}
	response, err := svc.client.CRMHandlerGetService(EntityUser, "?"+filter)
	if err != nil {
		return ResponseUser{}, err
	}

	return convertMarchalResponseUser(response)
}

func (svc *APIUserService) Post(User DomainUser) (DomainUser, error) {
	payload, err := json.Marshal(User)

	if err != nil {
		return DomainUser{}, err
	}

	response, err := svc.client.CRMHandlerPostService(EntityUser, payload)
	if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}
func (svc *APIUserService) Put(user DomainUserBase) (DomainUser, error) {

	payload, err := json.Marshal(user)

	if err != nil {
		return DomainUser{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityUser+"/"+user.ID, payload)
	if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}

func (svc *APIUserService) PutAccessInfo(ModelPut DomainUserSendAccessInfoPut) (DomainUser, error) {
	payload, err := json.Marshal(ModelPut)

	if err != nil {
		return DomainUser{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityUser+"/"+ModelPut.ID, payload)

  if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}

func (svc *APIUserService) Delete(id string) (bool, error) {
	expectedLen := 17
	idLen := len([]rune(id))

	if idLen != expectedLen {
		return false, fmt.Errorf("error: invalid id expected length of %d but got length of %d", expectedLen, idLen)
	}
	_, err := svc.client.CRMHandlerDeleteService(EntityUser, "/"+id)
	if err != nil {
		return false, err
	}

	return true, nil
}


func (svc *APIUserService) PutDashboardTemplate(ModelPutDash DomainUserDashBoardTemplate) (DomainUser, error) {
	payload, err := json.Marshal(ModelPutDash)

	if err != nil {
		return DomainUser{}, err
	}

	response, err := svc.client.CRMHandlerPutService(EntityUser+"/"+ModelPutDash.ID, payload)
	if err != nil {
		return DomainUser{}, err
	}

	return convertMarchalUser(response)
}

// Todos os serviços deverão ter o seu próprio conversor de json para o domain
func convertMarchalResponseUser(response []byte) (ResponseUser, error) {
	var result ResponseUser

	err := json.Unmarshal(response, &result)
	if err != nil {
		return ResponseUser{}, err
	}

	return result, nil
}

func convertMarchalUser(response []byte) (DomainUser, error) {
	var result DomainUser

	err := json.Unmarshal(response, &result)
	if err != nil {
		return DomainUser{}, err
	}

	return result, nil
}
