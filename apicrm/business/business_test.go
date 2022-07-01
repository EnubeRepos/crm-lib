package business

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

func TestGetBusiness(t *testing.T) {
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

func TestGetBusinessById(t *testing.T) {
	expectedId := "6283c7851f2b610e9"
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

func TestGetBusinessByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=linkedWith&where%5B0%5D%5Battribute%5D=teams&where%5B0%5D%5Bvalue%5D%5B%5D=62388f571a0bf1e48"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.GetByFilter(filter)

	if err != nil {
		t.Errorf("Error GETBYFILTER Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestPostBusiness(t *testing.T) {
	generetedId := common.GenerateUUID()
	expected := generetedId + "Test Enube"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainBusiness{
		Name:                       expected,
		EmailAddress:               generetedId + "_thomas@enube.me",
		PhoneNumber:                "0",
		BillingAddressStreet:       "1",
		BillingAddressCity:         "SP",
		BillingAddressState:        "SP",
		BillingAddressCountry:      "BR",
		BillingAddressNumber:       "34",
		BillingAddressPostalCode:   "33333",
		BillingAddressNeighborhood: "Granja",
		DocumentNumber:             "1234",
		TradingName:                "Thomas Test",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestPutBusiness(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Put(DomainBusiness{
		ID:                         "62bde9c2c2beaa692",
		Name:                       "Thomas Test",
		PhoneNumber:                "0",
		BillingAddressStreet:       "1",
		BillingAddressCity:         "SP",
		BillingAddressState:        "SP",
		BillingAddressCountry:      "BR",
		BillingAddressNumber:       "34",
		BillingAddressPostalCode:   "33333",
		BillingAddressNeighborhood: "Granja",
		DocumentNumber:             "123456",
		TradingName:                "Thomas Test",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestDeleteBusiness(t *testing.T) {
	id := "62a91edaabfbe44df"
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
