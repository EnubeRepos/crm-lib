package institutionfinancial

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
)

func TestPost(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	v, err := srvAccount.GetById("6232954e12f47d16a")

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	fmt.Println(v)
}
