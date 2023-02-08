package cardvirtual

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetAccount(t *testing.T) {

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

func TestGetById(t *testing.T) {
	expectedId := "62151eeecd1187df0"

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

func TestGetByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=contains&where%5B0%5D%5Battribute%5D=name&where%5B0%5D%5Bvalue%5D=Lucas"

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

	fmt.Print(res)

}

func TestPost(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainVirtualCard{
		Name:                       "Nicolas Carvalho",
		BillingAddressCity:         "São Paulo",
		BillingAddressNeighborhood: "Liberdade",
		BillingAddressStreet:       "Av nove de julho",
		BillingAddressComplement:   "Av.",
		BillingAddressCountry:      "BR",
		BillingAddressNumber:       "119",
		BillingAddressPostalCode:   "01513000",
		BillingAddressState:        "SP",
		AssignedUserID:             "6276b7d16bc4ae792",
		DocumentNumber:             "39234412826",
		CardName:                   "Nicolas Carvalho",
		Alias:                      "Nick",
		AccountNumber:              "261319",
		BankAgency:                 "0001",
		Password:                   "1234",
		BankAccountParentID:        "1234",
	})

	if err != nil {
		t.Errorf("Error POST Image:: error: %v", err)
		return
	}

	fmt.Println(res)
}

// func TestPut(t *testing.T) {

// 	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

// 	srvAccount := New(client)
// 	res, err := srvAccount.Put(DomainVirtualCardCreateRequest{
// 		ID:             "62b5f5aa128a56e4f",
// 		Name:           "Thomas NovoNovo",
// 		LastName:       "Thomas",
// 		FirstName:      "Novo3",
// 		AccountNumber:  "00088",
// 		AssignedUserID: "1",
// 		SicCode:        "0",
// 		Agency:         "00008",
// 	})

// 	if err != nil {
// 		t.Errorf("Error PUT Account:: error: %v", err)
// 		return
// 	}

// 	fmt.Println(res)

// }

func TestDelete(t *testing.T) {
	id := "6272dfb1d6499bae2"
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
