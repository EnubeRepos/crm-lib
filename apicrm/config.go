package apicrm

type CRMAPIClient struct {
	BaseURL   string
	UrlPrefix string
	UseHttps  bool

	Token string
}

type CRMAPIConfig struct {
	BaseURL, URLPrefix, Token string
	UseHttps                  bool
}

func New(host, token string) CRMAPIClient {
	return CRMAPIClient{
		BaseURL: host,
		Token:   token,
	}
}
