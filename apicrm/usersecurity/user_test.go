package usersecurity

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestPutUser(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	repo := New(client)
	res, err := repo.UpdateStatusMFA(DomainUserStatusMFA{
		Auth2FA:       false,
		Auth2FAMethod: "Email",
		Password:      "123",
	}, "6276b7d16bc4ae792")

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}
