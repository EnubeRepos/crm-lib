package usersecurity

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
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
