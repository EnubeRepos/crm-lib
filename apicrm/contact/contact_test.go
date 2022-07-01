package contact

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetContact(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvContact := New(client)
	res, err := srvContact.Get()

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}

	fmt.Println(res)
}

func TestGetContactById(t *testing.T) {
	expectedId := "6189a36020f1f2fab"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvContact := New(client)
	res, err := srvContact.GetById(expectedId)

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestGetContactByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=startsWith&where%5B0%5D%5Battribute%5D=emailAddress&where%5B0%5D%5Bvalue%5D=T"
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

func TestPostContact(t *testing.T) {
	expectedName := "Thomas Test1"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainContact{
		Name:              expectedName,
		FirstName:         "Thomas",
		LastName:          "Test1",
		SalutationName:    "Thomas",
		AccountName:       "Thomas",
		CampaignName:      "Thomas",
		CreatedByName:     "Thomas",
		AssignedUserName:  "Thomas",
		EmailAddress:      "thomas@enube.me",
		AssignedUserID:    "1",
		SicCode:           "0",
		PhoneNumber:       "14981288851",
		AddressStreet:     "Rua 123 Sao Paulo",
		AddressCity:       "SP",
		AddressState:      "SP",
		AddressCountry:    "BR",
		AddressPostalCode: "33331",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestPutCommissions(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Put(DomainContactBase{
		ID:                "62bdf7991cc8b3aff",
		Name:              "Thomas",
		FirstName:         "Thomas",
		LastName:          "Test2",
		SalutationName:    "Thomas",
		AccountName:       "Thomas",
		CreatedByName:     "Thomas",
		AssignedUserName:  "Thomas",
		EmailAddress:      "thomaswwww@enube.me",
		AssignedUserID:    "1",
		SicCode:           "0",
		PhoneNumber:       "14981288851",
		AddressStreet:     "Rua 123 Sao Paulo",
		AddressCity:       "SP",
		AddressState:      "SP",
		AddressCountry:    "BR",
		AddressPostalCode: "33331",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestDeleteContact(t *testing.T) {
	id := "62a1f8057e9250661"
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
