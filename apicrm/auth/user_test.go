package user

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
	client := crmapi.NewBySession(HOST, TOKEN, "_")

	srvUser := New(client)
	v, err := srvUser.Get()

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}

	fmt.Println(v)
}

func TestGenerateCodeMFA(t *testing.T) {

	client := crmapi.NewBySession(HOST, "ODk1MTkzODIwNDY6MTIzNDU2", "")

	srvUser := New(client)
	v, err := srvUser.SendCodeMFA()

	if err != nil {
		t.Errorf("Error GET :: error: %v", err)
		return
	}

	fmt.Println(v)
}

func TestConfirmCodeMFA(t *testing.T) {

	client := crmapi.NewBySession(HOST, "ODk1MTkzODIwNDY6MTIzNDU2", "")

	srvUser := New(client)
	v, err := srvUser.ConfirmCodeMFA("9935665")

	if err != nil {
		t.Errorf("Error GET :: error: %v", err)
		return
	}

	fmt.Println(v)
}
