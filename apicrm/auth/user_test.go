package user

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestPost(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvUser := New(client)
	v, err := srvUser.Get()

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	fmt.Println(v)
}
