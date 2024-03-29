package user

const EntityUser string = "User"

// ResponseUser .
type ResponseUser struct {
	Total int          `json:"total,omitempty"`
	User  []DomainUser `json:"list,omitempty"`
}

// estrutura usada para o put de registration, evitando erro de campos vazios
type DomainUserBase struct {
	ID           string `json:"id,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

// Account
type DomainUser struct {
	ID                     string             `json:"id,omitempty"`
	Name                   string             `json:"name,omitempty"`
	Deleted                bool               `json:"deleted,omitempty"`
	UserName               string             `json:"userName,omitempty"`
	Type                   string             `json:"type,omitempty"`
	AuthMethod             string             `json:"authMethod,omitempty"`
	APIKey                 string             `json:"apiKey,omitempty"`
	SalutationName         string             `json:"salutationName,omitempty"`
	FirstName              string             `json:"firstName,omitempty"`
	LastName               string             `json:"lastName,omitempty"`
	IsActive               bool               `json:"isActive,omitempty"`
	Title                  string             `json:"title,omitempty"`
	EmailAddress           string             `json:"emailAddress,omitempty"`
	PhoneNumber            string             `json:"phoneNumber,omitempty"`
	Gender                 string             `json:"gender,omitempty"`
	CreatedAt              string             `json:"createdAt,omitempty"`
	ModifiedAt             string             `json:"modifiedAt,omitempty"`
	Auth2FA                bool               `json:"auth2FA,omitempty"`
	LastAccess             string             `json:"lastAccess,omitempty"`
	SicCode                string             `json:"sicCode,omitempty"`
	ProfileType            string             `json:"profileType,omitempty"`
	AddressCountry         string             `json:"addressCountry,omitempty"`
	MiddleName             string             `json:"middleName,omitempty"`
	EmailAddressIsOptedOut bool               `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut  bool               `json:"phoneNumberIsOptedOut,omitempty"`
	AddressStreet          string             `json:"addressStreet,omitempty"`
	AddressCity            string             `json:"addressCity,omitempty"`
	AddressState           string             `json:"addressState,omitempty"`
	AddressPostalCode      string             `json:"addressPostalCode,omitempty"`
	EmailAddressData       []EmailAddressData `json:"emailAddressData,omitempty"`
	PhoneNumberData        []PhoneNumberData  `json:"phoneNumberData,omitempty"`
	DefaultTeamID          string             `json:"defaultTeamId,omitempty"`
	DefaultTeamName        string             `json:"defaultTeamName,omitempty"`
	TeamsIds               []string           `json:"teamsIds,omitempty"`
	RolesIds               []string           `json:"rolesIds,omitempty"`
	PortalsIds             []string           `json:"portalsIds,omitempty"`
	PortalRolesIds         []string           `json:"portalRolesIds,omitempty"`
	ContactID              string             `json:"contactId,omitempty"`
	ContactName            string             `json:"contactName,omitempty"`
	AccountsIds            []string           `json:"accountsIds,omitempty"`
	AvatarID               string             `json:"avatarId,omitempty"`
	AvatarName             string             `json:"avatarName,omitempty"`
	CreatedByID            string             `json:"createdById,omitempty"`
	CreatedByName          string             `json:"createdByName,omitempty"`
	DashboardTemplateID    string             `json:"dashboardTemplateId,omitempty"`
	DashboardTemplateName  string             `json:"dashboardTemplateName,omitempty"`
	RealEstateIds          []string           `json:"realEstateIds,omitempty"`
	FinTransactionsIds     []string           `json:"finTransactionsIds,omitempty"`
	AccreditedAccountsIds  []string           `json:"accreditedAccountsIds,omitempty"`
	RegistrationID         string             `json:"registrationId,omitempty"`
	RegistrationName       string             `json:"registrationName,omitempty"`
	PasswordPreview        string             `json:"passwordPreview,omitempty"`
	Password               string             `json:"password,omitempty"`
	PasswordConfirm        string             `json:"passwordConfirm,omitempty"`
	SendAccessInfo         bool               `json:"sendAccessInfo,omitempty"`
	BankAccountId          string             `json:"bankAccountId,omitempty"`
	ExternalIdSign         string             `json:"externalIdSign,omitempty"`
}

type EmailAddressData struct {
	EmailAddress string `json:"emailAddress,omitempty"`
	Lower        string `json:"lower,omitempty"`
	Primary      bool   `json:"primary,omitempty"`
	OptOut       bool   `json:"optOut,omitempty"`
	Invalid      bool   `json:"invalid,omitempty"`
}
type PhoneNumberData struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Type        string `json:"type,omitempty"`
	Primary     bool   `json:"primary,omitempty"`
	OptOut      bool   `json:"optOut,omitempty"`
	Invalid     bool   `json:"invalid,omitempty"`
}

type DomainUserSendAccessInfoPut struct {
	ID              string `json:"id,omitempty"`
	PasswordPreview string `json:"passwordPreview,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"passwordConfirm,omitempty"`
	SendAccessInfo  bool   `json:"sendAccessInfo,omitempty"`
}

type DomainUserDashBoardTemplate struct {
	ID                  string `json:"id,omitempty"`
	DashboardTemplateID string `json:"dashboardTemplateId,omitempty"`
}

type DomainUserStatusMFA struct {
	Auth2FA       string `json:"auth2FA,omitempty"`
	Auth2FAMethod string `json:"auth2FAMethod,omitempty"`
	Password      string `json:"password,omitempty"`
}
