package context

import "github.com/EnubeRepos/CRM_framework/internal/common"

// Config contains the configurations for application.
type Config struct {
	// Application version.
	Version string

	// AWSClientID client id.
	AWSClientID string

	// AWSClientSecret client secret.
	AWSClientSecret string

	// HTTPPrefix is http prefix for web endpoints.
	HTTPPrefix string

	// Webservice Westcon
	HOST_WS string

	// API Westcon
	HOST_WS_RESELLER_API string

	// HOST CRM
	HOST_API_CRM string

	HOST_CRM string

	// HOST CRM
	TOKEN_CRM string

	//API_PARTNER_CENTER
	API_PARTNER_CENTER string

	//API_ROCKET
	API_ROCKET string

	// TOKEN_PARTNER_CENTER
	TOKEN_PARTNER_CENTER string

	//TENANT_PARTNER_CENTER
	TENANT_PARTNER_CENTER string

	//CLIENT_ID_PARTNER_CENTER
	CLIENT_ID_PARTNER_CENTER string

	// SECRET_ID_PARTNER_CENTER
	SECRET_ID_PARTNER_CENTER string

	//API_RESOURCE_MANAGEMENT
	API_RESOURCE_MANAGEMENT string

	// TOKEN_RESOURCE_MANAGEMENT
	TOKEN_RESOURCE_MANAGEMENT string

	//TENANT_RESOURCE_MANAGEMENT
	TENANT_RESOURCE_MANAGEMENT string

	//CLIENT_ID_RESOURCE_MANAGEMENT
	CLIENT_ID_RESOURCE_MANAGEMENT string

	// SECRET_ID_RESOURCE_MANAGEMENT
	SECRET_ID_RESOURCE_MANAGEMENT string

	// PATH_DOWNLOAD
	PATH_DOWNLOAD string

	//API_PARTNER_REFERRALS
	API_PARTNER_REFERRALS string
}

// InitConfig load the configuration for application.
func NewConfig() (*Context, error) {
	config := &Config{
		Version: common.GetEnv("VERSION", ""),
		// AWSClientID:                   viper.GetString("AWS_CLIENT_ID"),
		// AWSClientSecret:               viper.GetString("AWS_CLIENT_SECRET"),
		// HTTPPrefix:                    viper.GetString("HTTP_PREFIX"),
		// HOST_WS:                       viper.GetString("HOST_WS"),
		// HOST_WS_RESELLER_API:          viper.GetString("HOST_WS_RESELLER_API"),
		HOST_CRM:     common.GetEnv("HOST_CRM", ""),
		HOST_API_CRM: common.GetEnv("HOST_API_CRM", ""),
		// TOKEN_CRM:                     viper.GetString("TOKEN_CRM"),
		API_PARTNER_CENTER: common.GetEnv("API_PARTNER_CENTER", ""),
		// API_ROCKET:                    viper.GetString("API_ROCKET"),
		// TOKEN_PARTNER_CENTER:          viper.GetString("TOKEN_PARTNER_CENTER"),
		// TENANT_PARTNER_CENTER:         viper.GetString("TENANT_PARTNER_CENTER"),
		// CLIENT_ID_PARTNER_CENTER:      viper.GetString("CLIENT_ID_PARTNER_CENTER"),
		// SECRET_ID_PARTNER_CENTER:      viper.GetString("SECRET_ID_PARTNER_CENTER"),
		// API_RESOURCE_MANAGEMENT:       viper.GetString("API_RESOURCE_MANAGEMENT"),
		// TOKEN_RESOURCE_MANAGEMENT:     viper.GetString("TOKEN_RESOURCE_MANAGEMENT"),
		// TENANT_RESOURCE_MANAGEMENT:    viper.GetString("TENANT_RESOURCE_MANAGEMENT"),
		// CLIENT_ID_RESOURCE_MANAGEMENT: viper.GetString("CLIENT_ID_RESOURCE_MANAGEMENT"),
		// SECRET_ID_RESOURCE_MANAGEMENT: viper.GetString("SECRET_ID_RESOURCE_MANAGEMENT"),
		// PATH_DOWNLOAD:                 viper.GetString("PATH_DOWNLOAD"),
		API_PARTNER_REFERRALS: common.GetEnv("API_PARTNER_REFERRALS", ""),
	}

	c := &Context{
		Config: *config,
	}

	return c, nil
}

// Context represent a context for application.
type Context struct {
	Config Config
}
