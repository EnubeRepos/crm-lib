package tests

import (
	"testing"

	"github.com/EnubeRepos/crm-lib/apicrm/account"
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

	srvAccount := account.NewAPIAccountService(client)
	res, err := srvAccount.Get()

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}
}

func TestAddNewAccount(t *testing.T) {
	generetedId := common.GenerateUUID()
	expected := generetedId + "Test Enube"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := account.NewAPIAccountService(client)
	res, err := srvAccount.Post(account.DomainAccount{
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

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Name != expected {
		t.Errorf("Error GET Account %q, wanted %q", res.ID, expected)
	}
}

func TestAddSameAccountLotofTime(t *testing.T) {

	generetedId := "17DF06D5-8AEA-9353-00D7-FB38D3E94D731"
	expected := generetedId + "Test Enube"

	for i := 1; i < 5; i++ {
		client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))
		srvAccount := account.NewAPIAccountService(client)
		res, err := srvAccount.Post(account.DomainAccount{
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
