package user

import (
	"strings"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

func (svc *APIAuthService) SendCodeMFA() (DomainAuth, error) {

	svc.client.Headers = []crmapi.Headers{
		{
			Key:   "Espo-Authorization",
			Value: svc.client.Token,
		},
		{
			Key:   "Espo-Authorization-Create-Token-Secret",
			Value: "true",
		},
	}

	response, err := svc.client.CRMHandlerGetService(EntityAuth, "")
	if err != nil {
		return DomainAuth{}, err
	}

	return convertMarchalDomainAuth(response)
}

func (svc *APIAuthService) ConfirmCodeMFA(code string) (DomainAuth, error) {

	svc.client.Headers = []crmapi.Headers{
		{
			Key:   "Espo-Authorization",
			Value: svc.client.Token,
		},
		{
			Key:   "Espo-Authorization-Create-Token-Secret",
			Value: "true",
		},
		{
			Key:   "Espo-Authorization-Code",
			Value: code,
		},
	}

	response, err := svc.client.CRMHandlerGetService(EntityAuth, "?code="+code)
	if err != nil {
		return DomainAuth{}, err
	}

	secretCookie := svc.client.HeadersResponse["Set-Cookie"]

	res, err := convertMarchalDomainAuth(response)

	if err != nil {
		return DomainAuth{}, err
	}

	if len(secretCookie) > 0 {
		fraseSlice := strings.Split(secretCookie[0], ";")

		for _, item := range fraseSlice {
			if strings.Contains(item, "secret") {

				secret := strings.Split(item, "=")
				res.TokenSecret = secret[1]
			}
		}
	}

	return res, err
}
