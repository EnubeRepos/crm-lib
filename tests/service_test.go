package tests

import (
	"testing"

	"github.com/EnubeRepos/CRM_framework/internal/core/services/accountsrv"
	"github.com/EnubeRepos/CRM_framework/internal/crmapi"
)

const (
	HOST  = "https://cloudsolvssa-col.cloudanalytics.me/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetAccount(t *testing.T) {

	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srv := accountsrv.NewAPIAccountService(client)
	res, err := srv.GetAccount()

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}
}
