package user

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "ODk1MTkzODIwNDY6MDkxOWE4YjdhOGM1ZDRjZTBiZTEzMGQyMjA4M2QyNDQ="
)

func TestPost(t *testing.T) {
	client := crmapi.NewBySession(HOST, TOKEN, "auth-token-secret=b5283ce327960c8c9723c69b614ecebb; auth-username=89519382046; auth-token=0919a8b7a8c5d4ce0be130d22083d244")

	srvUser := New(client)
	v, err := srvUser.Get()

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	fmt.Println(v)
}
