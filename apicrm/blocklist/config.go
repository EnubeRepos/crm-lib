package blocklist

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIBlockListService struct {
	client crmapi.CRMAPIClient
}

func New(client crmapi.CRMAPIClient) *APIBlockListService {
	return &APIBlockListService{
		client: client,
	}
}
