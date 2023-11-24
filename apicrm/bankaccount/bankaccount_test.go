package bankaccount

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
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
	res, err := srvAccount.Post(DomainBankAccountCreateRequest{
		Agency:         "00007",
		AccountNumber:  "00011",
		AssignedUserID: "62151eeecd1187df9",
		SicCode:        "40760176846",
		LastName:       "Test3",
		FirstName:      "Thomas",
	})

	if err != nil {
		t.Errorf("Error POST Image:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestPut(t *testing.T) {

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Put(DomainBankAccountCreateRequest{
		ID:             "62b5f5aa128a56e4f",
		Name:           "Thomas NovoNovo",
		LastName:       "Thomas",
		FirstName:      "Novo3",
		AccountNumber:  "00088",
		AssignedUserID: "1",
		SicCode:        "0",
		Agency:         "00008",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

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
