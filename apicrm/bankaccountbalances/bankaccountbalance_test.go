package bankaccountbalances

import (
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestPost(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	v, err := srvAccount.Post(DomainBankAccountBalanceCreateRequest{
		ValueAvailable: 20,
		ValueInProcess: 0,
		ValueBlocked: 0,
		BankAccountId: "61b3a3c6b640f7674"})

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	if v.ID == "" {
		t.Errorf("Error On create balance:: error: %v", err)
		return
	}

}
