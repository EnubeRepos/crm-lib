package user

const EntityUser string = "User"

// ResponseUser .
type ResponseUser struct {
	Total int          `json:"total,omitempty"`
	User  []DomainUser `json:"list,omitempty"`
}

// Account
type DomainUser struct {
	ID                     string             `json:"id"`
	Name                   string             `json:"name"`
	Deleted                bool               `json:"deleted"`
	UserName               string             `json:"userName"`
	Type                   string             `json:"type"`
	AuthMethod             string             `json:"authMethod"`
	APIKey                 string             `json:"apiKey"`
	SalutationName         string             `json:"salutationName"`
	FirstName              string             `json:"firstName"`
	LastName               string             `json:"lastName"`
	IsActive               bool               `json:"isActive"`
	Title                  string             `json:"title"`
	EmailAddress           string             `json:"emailAddress"`
	PhoneNumber            string             `json:"phoneNumber"`
	Gender                 string             `json:"gender"`
	CreatedAt              string             `json:"createdAt"`
	ModifiedAt             string             `json:"modifiedAt"`
	Auth2FA                string             `json:"auth2FA"`
	LastAccess             string             `json:"lastAccess"`
	SicCode                string             `json:"sicCode"`
	ProfileType            string             `json:"profileType"`
	AddressCountry         string             `json:"addressCountry"`
	MiddleName             string             `json:"middleName"`
	EmailAddressIsOptedOut bool               `json:"emailAddressIsOptedOut"`
	PhoneNumberIsOptedOut  bool               `json:"phoneNumberIsOptedOut"`
	AddressStreet          string             `json:"addressStreet"`
	AddressCity            string             `json:"addressCity"`
	AddressState           string             `json:"addressState"`
	AddressPostalCode      string             `json:"addressPostalCode"`
	EmailAddressData       []EmailAddressData `json:"emailAddressData"`
	PhoneNumberData        []PhoneNumberData  `json:"phoneNumberData"`
	DefaultTeamID          string             `json:"defaultTeamId"`
	DefaultTeamName        string             `json:"defaultTeamName"`
	TeamsIds               []string           `json:"teamsIds"`
	RolesIds               []string           `json:"rolesIds"`
	PortalsIds             []string           `json:"portalsIds"`
	PortalRolesIds         []string           `json:"portalRolesIds"`
	ContactID              string             `json:"contactId"`
	ContactName            string             `json:"contactName"`
	AccountsIds            []string           `json:"accountsIds"`
	AvatarID               string             `json:"avatarId"`
	AvatarName             string             `json:"avatarName"`
	CreatedByID            string             `json:"createdById"`
	CreatedByName          string             `json:"createdByName"`
	DashboardTemplateID    string             `json:"dashboardTemplateId"`
	DashboardTemplateName  string             `json:"dashboardTemplateName"`
	RealEstateIds          []string           `json:"realEstateIds"`
	FinTransactionsIds     []string           `json:"finTransactionsIds"`
	AccreditedAccountsIds  []string           `json:"accreditedAccountsIds"`
	RegistrationID         string             `json:"registrationId"`
	RegistrationName       string             `json:"registrationName"`
}
type EmailAddressData struct {
	EmailAddress string `json:"emailAddress"`
	Lower        string `json:"lower"`
	Primary      bool   `json:"primary"`
	OptOut       bool   `json:"optOut"`
	Invalid      bool   `json:"invalid"`
}
type PhoneNumberData struct {
	PhoneNumber string `json:"phoneNumber"`
	Type        string `json:"type"`
	Primary     bool   `json:"primary"`
	OptOut      bool   `json:"optOut"`
	Invalid     bool   `json:"invalid"`
}
