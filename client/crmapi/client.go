package crmapi

type CRMAPIClient struct {
	BaseURL         string
	UrlPrefix       string
	UseHttps        bool
	Cookie          string
	Headers         []Headers
	HeadersResponse map[string][]string

	Token string
}

type CRMAPIConfig struct {
	BaseURL, URLPrefix, Token string
	UseHttps                  bool
	Cookie                    string
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

func New(host, token string) CRMAPIClient {
	return CRMAPIClient{
		BaseURL: host,
		Token:   token,
	}
}

func NewBySession(host, token, Cookie string) CRMAPIClient {
	return CRMAPIClient{
		BaseURL: host,
		Token:   token,
		Cookie:  Cookie,
	}
}
