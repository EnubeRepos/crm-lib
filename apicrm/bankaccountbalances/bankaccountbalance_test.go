package bankaccountbalances

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
)

func TestGet(t *testing.T) {
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
	expectedId := "62a7883b667fc8df0"
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

func TestPost(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainBankAccountBalanceCreateRequest{
		ValueAvailable:   20,
		ValueInProcess:   13,
		ValueBlocked:     34,
		BankAccountId:    "6272dfb1d6499bae2",
		BankAccountName:  "Thomas Test",
		AssignedUser:     "",
		AssignedUserName: "thomas thomas",
		AssignedUserId:   "12345",
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
	res, err := srvAccount.Put(DomainBankAccountBalanceCreateRequest{
		ID:               "62a8da8a8f4c6faf9",
		ValueSumVirtual:  69,
		ValueAvailable:   20,
		ValueInProcess:   13,
		ValueBlocked:     34,
		BankAccountId:    "62b5f5aa128a56e4f",
		BankAccountName:  "Thomas Test",
		AssignedUser:     "",
		AssignedUserName: "thomas thomas",
		AssignedUserId:   "12345",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestDelete(t *testing.T) {
	id := "62acb1e9a311a02bb"
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
