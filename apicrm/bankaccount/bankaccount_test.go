package bankaccount

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
	expectedId := "629a268f46ce2442f"

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
	res, err := srvAccount.Post(DomainBankAccountCreateRequest{Agency: "00006", AccountNumber: "00009", AssignedUserID: "62151eeecd1187df0", SicCode: "40760176845", LastName: "Test", FirstName: "Thomas"})

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestDelete(t *testing.T) {
	id := "628cdf2613adbfdd0"
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
