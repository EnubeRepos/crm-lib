package sales

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetSale(t *testing.T) {
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

func TestGetSaleById(t *testing.T) {
	expectedId := "628d13566ca7368b9"
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

// test it
func TestGetSaleByFilter(t *testing.T) {
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

	//fmt.Print(res)

}

func TestPostSale(t *testing.T) {
	expectedUserName := "Thomas Test"

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainSales{
		ID:                    "123",
		AssignedUserID:        "1",
		AssignedUser:          "Thomas",
		AssignedUserName:      expectedUserName,
		SalesNumber:           "27",
		SalesDate:             "20.07.2002",
		CommissionTotalAmount: 500,
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)

	// if res.AssignedUserName != expectedUserName {
	// 	t.Errorf("Error POST Account %q, wanted %q", res.ID, expectedUserName)
	// }
}

func TestDeleteSale(t *testing.T) {
	id := "62a9fdddcc860b8ac"
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
