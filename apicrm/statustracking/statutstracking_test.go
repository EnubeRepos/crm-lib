package statustracking

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
)

func TestGetStatus(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Get()

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}
}

func TestGetStatusById(t *testing.T) {
	expectedId := "6219196ac2d4aa230"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.GetById(expectedId)

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.ID != expectedId {
		t.Errorf("Error GetId Account %s, wanted %s", res.ID, expectedId)
	}
}

func TestGetStatusByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=linkedWith&where%5B0%5D%5Battribute%5D=teams&where%5B0%5D%5Bvalue%5D%5B%5D=62388f571a0bf1e48"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.GetByFilter(filter)

	if err != nil {
		t.Errorf("Error GETBYFILTER Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GETBYFILTER Account %q, wanted %q", res.Total, 1)
	}

}

func TestPostStatus(t *testing.T) {

	expectedCode := 890
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainStatusTracking{
		ID:               "123",
		Code:             expectedCode,
		AssignedUser:     "Test",
		AssignedUserID:   "123",
		AssignedUserName: "Thomas Test",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestPutStatus(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Put(DomainStatusTrackingBase{
		ID:         "6219196ac2d4aa230",
		StatusType: "Ativo",
		Code:       3,
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestDeleteStatus(t *testing.T) {
	id := "62a9f76b891c97477"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)

	res, err := srvAccount.Delete(id)

	if err != nil {
		t.Errorf("Error DELETE Account:: error: %v", err)
		return
	}

	if res == false {
		t.Errorf("Error DELETE: Account not deleted")
		return
	}
}
