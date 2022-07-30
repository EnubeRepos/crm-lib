package registration

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetRegistration(t *testing.T) {
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

func TestGetResgistartionById(t *testing.T) {
	expectedId := "628f9a22e77ea1451"
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

func TestGetResgistrationByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=startsWith&where%5B0%5D%5Battribute%5D=name&where%5B0%5D%5Bvalue%5D=Carolina"
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

func TestPostRegistration(t *testing.T) {

	expected := "Thomas Kuhn"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainRegistration{
		Name:                  expected,
		EmailAddress:          "thomas@enube.me",
		AssignedUserID:        "1",
		PhoneNumber:           "0",
		BillingAddressStreet:  "1",
		BillingAddressCity:    "SP",
		BillingAddressState:   "SP",
		BillingAddressCountry: "BR",
		DocumentNumber:        "40",
		BillingAddressNumber:  "Rua 123",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestPut(t *testing.T) {

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Put(DomainRegistrationBase{
		ID:                 "629f68d88acbf0228",
		StatusDetail:       "Aprovado",
		StatusCid:          "3331",
		StatusProcess:      "Aprovado",
		StatusDatetime:     "2020-07-07",
		RegistrationStatus: "Aprovado",
		AssignedUserID:     "1",
		BankAccountID:      "56",
		EmailAddress:       "test@enube.me",
		BillingAddressCity: "POA",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestDeleteRegistration(t *testing.T) {
	id := "629f67ee202b17175"
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
