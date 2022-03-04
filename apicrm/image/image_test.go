package image

import (
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestNewAPIImageService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvImage := NewAPIImageService(client)
	_, err := srvImage.Get("62153f3f7877ee35d")

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}
}
