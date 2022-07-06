package image

import (
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.triplicbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestNewAPIImageService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvImage := NewAPIImageService(client)
	_, err := srvImage.Get("62c34a3a955723e4a")

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}
}
