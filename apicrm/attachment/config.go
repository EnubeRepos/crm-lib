package attachment

import "github.com/EnubeRepos/crm-lib/client/crmapi"

type APIAttachmentService struct {
	client crmapi.CRMAPIClient
}

func NewAPIAttachmentService(client crmapi.CRMAPIClient) *APIAttachmentService {
	return &APIAttachmentService{
		client: client,
	}
}
