package generic

import (
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
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
