package account

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/common"
	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetAccount(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAccountService(client)
	res, err := srvAccount.Get()

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total > 0 && len(res.Account) <=0{
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}
}

func TestGetAccountById(t *testing.T) {
	expectedId := "61899dfcc4f59bf6d"

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAccountService(client)
	res, err := srvAccount.GetById(expectedId)

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.ID != expectedId {
		t.Errorf("Error GetId Account %s, wanted %s", res.ID, expectedId)
	}

	fmt.Println(res)

}

func TestGetByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=linkedWith&where%5B0%5D%5Battribute%5D=teams&where%5B0%5D%5Bvalue%5D%5B%5D=62388f571a0bf1e48"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAccountService(client)
	res, err := srvAccount.GetByFilter(filter)

	if err != nil {
		t.Errorf("Error GETBYFILTER Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GETBYFILTER Account %q, wanted %q", res.Total, 1)
	}

}

func TestDelete(t *testing.T) {
	id := "629e198aa66904a4d"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAccountService(client)

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

func TestPost(t *testing.T) {
	generetedId := common.GenerateUUID()
	expected := generetedId + "Test Enube"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAccountService(client)
	res, err := srvAccount.Post(DomainAccount{
		Name:                  expected,
		EmailAddress:          generetedId + "_thomas@enube.me",
		AssignedUserID:        "1",
		SicCode:               "0",
		PhoneNumber:           "0",
		BillingAddressStreet:  "1",
		BillingAddressCity:    "SP",
		BillingAddressState:   "SP",
		BillingAddressCountry: "BR",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestPut(t *testing.T) {

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAccountService(client)
	res, err := srvAccount.Put(DomainAccountBase{
		ID:                    "62b4b4efcf6eddb2d",
		Name:                  "Thomas New",
		EmailAddress:          "test@enube.me",
		AssignedUserID:        "1",
		SicCode:               "0",
		PhoneNumber:           "0",
		BillingAddressStreet:  "1",
		BillingAddressCity:    "SP",
		BillingAddressState:   "SP",
		BillingAddressCountry: "BR",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestAddSameAccountLotofTime(t *testing.T) {

	generetedId := "17DF06D5-8AEA-9353-00D7-FB38D3E94D731"
	expected := generetedId + "Test Enube"

	for i := 1; i < 5; i++ {
		client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))
		srvAccount := NewAPIAccountService(client)
		res, err := srvAccount.Post(DomainAccount{
			Name:                  expected,
			EmailAddress:          generetedId + "_nicolas@enube.me",
			AssignedUserID:        "1",
			SicCode:               "0",
			PhoneNumber:           "0",
			BillingAddressStreet:  "1",
			BillingAddressCity:    "SP",
			BillingAddressState:   "SP",
			BillingAddressCountry: "BR",
		})

		if err == nil {
			t.Errorf("Error Insert Lot time the same Account %q, wanted %q", res.ID, expected)
		}

		if res.ID != "" {
			t.Errorf("Error Insert Lot time the same Account %q, wanted %q", res.ID, expected)
		}
	}
}
