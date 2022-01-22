package crmapi

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

func NewCRMAPIConfig(host, token string) CRMAPIConfig {
	return CRMAPIConfig{
		BaseURL: host,
		Token:   token,
	}
}

func NewCRMAPIClient(config CRMAPIConfig) CRMAPIClient {
	return CRMAPIClient{
		BaseURL: config.BaseURL,
		Token:   config.Token,
	}
}
