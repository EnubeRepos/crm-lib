package dashboardtemplate

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIDashboardTemplateService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIDashboardTemplateService {
	return &APIDashboardTemplateService{
		client: client,
	}
}
