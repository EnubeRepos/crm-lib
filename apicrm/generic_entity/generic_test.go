package account

import (
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetGeneric(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvGeneric := NewAPIGenericService(client, "")
	res, err := srvGeneric.Get()

	if err != nil {
		t.Errorf("Error GET Generic:: error: %v", err)
		return
	}

	if res != nil {
		t.Errorf("Error GET Generic:: error: %v", err)
		return
	}
}
