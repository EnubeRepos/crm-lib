package generic

const EntityGeneric string = "Generic"

// ResponseGeneric .
type ResponseGeneric struct {
	Total   int             `json:"total,omitempty"`
	Generic []DomainGeneric `json:"list,omitempty"`
}

// estrutura usada para o put de registration, evitando erro de campos vazios
type DomainGenericBase struct {
	ID                    string `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	EmailAddress          string `json:"emailAddress,omitempty"`
	SicCode               string `json:"sicCode,omitempty"`
	PhoneNumber           string `json:"phoneNumber,omitempty"`
	BillingAddressStreet  string `json:"billingAddressStreet,omitempty"`
	BillingAddressCity    string `json:"billingAddressCity,omitempty"`
	BillingAddressState   string `json:"billingAddressState,omitempty"`
	BillingAddressCountry string `json:"billingAddressCountry,omitempty"`
	AssignedUserID        string `json:"assignedUserId,omitempty"`
}

// Generic
type DomainGeneric struct {
	ID                        string                   `json:"id,omitempty"`
	Name                      string                   `json:"name,omitempty"`
	Deleted                   bool                     `json:"deleted,omitempty"`
	Website                   string                   `json:"website,omitempty"`
	Domain                    string                   `json:"domain,omitempty"`
	EmailAddress              string                   `json:"emailAddress,omitempty"`
	PhoneNumber               string                   `json:"phoneNumber,omitempty"`
	Type                      string                   `json:"type,omitempty"`
	Industry                  string                   `json:"industry,omitempty"`
	SicCode                   string                   `json:"sicCode,omitempty"`
	BillingAddressStreet      string                   `json:"billingAddressStreet,omitempty"`
	BillingAddressCity        string                   `json:"billingAddressCity,omitempty"`
	BillingAddressState       string                   `json:"billingAddressState,omitempty"`
	BillingAddressCountry     string                   `json:"billingAddressCountry,omitempty"`
	BillingAddressPostalCode  string                   `json:"billingAddressPostalCode,omitempty"`
	ShippingAddressStreet     string                   `json:"shippingAddressStreet,omitempty"`
	ShippingAddressCity       string                   `json:"shippingAddressCity,omitempty"`
	ShippingAddressState      string                   `json:"shippingAddressState,omitempty"`
	ShippingAddressCountry    string                   `json:"shippingAddressCountry,omitempty"`
	ShippingAddressPostalCode string                   `json:"shippingAddressPostalCode,omitempty"`
	Description               string                   `json:"description,omitempty"`
	CreatedAt                 string                   `json:"createdAt,omitempty"`
	ModifiedAt                string                   `json:"modifiedAt,omitempty"`
	CustomerCode              string                   `json:"customerCode,omitempty"`
	Subscriptions             []string                 `json:"subscriptions,omitempty"`
	IsPartner                 bool                     `json:"isPartner,omitempty"`
	ExternalID                string                   `json:"externalID,omitempty"`
	EmailAddressIsOptedOut    bool                     `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut     bool                     `json:"phoneNumberIsOptedOut,omitempty"`
	EmailAddressData          []DomainEmailAddressData `json:"emailAddressData,omitempty"`
	PhoneNumberData           []DomainPhoneNumberData  `json:"phoneNumberData,omitempty"`
	CampaignID                string                   `json:"campaignId,omitempty"`
	CampaignName              string                   `json:"campaignName,omitempty"`
	CreatedByID               string                   `json:"createdById,omitempty"`
	CreatedByName             string                   `json:"createdByName,omitempty"`
	ModifiedByID              string                   `json:"modifiedById,omitempty"`
	ModifiedByName            string                   `json:"modifiedByName,omitempty"`
	AssignedUserID            string                   `json:"assignedUserId,omitempty"`
	AssignedUserName          string                   `json:"assignedUserName,omitempty"`
	TeamsIds                  []string                 `json:"teamsIds,omitempty"`
	OriginalLeadID            string                   `json:"originalLeadId,omitempty"`
	OriginalLeadName          string                   `json:"originalLeadName,omitempty"`
	ClientID                  string                   `json:"clientId,omitempty"`
	ClientSecret              string                   `json:"clientSecret,omitempty"`

	CodigoERP                 string `json:"erpCode,omitempty"`
	StatusExternal            string `json:"statusExternal,omitempty"`
	GenericParentId           string `json:"accountParentId,omitempty"`
	VendorGenericId           string `json:"vendorGenericId,omitempty"`
	CountryList               string `json:"countryList,omitempty"`
	MainContactName           string `json:"mainContactName,omitempty"`
	MainContactId             string `json:"mainContactId,omitempty"`
	BillingGenericContactId   string `json:"billingGenericContactId,omitempty"`
	BillingGenericContactName string `json:"billingGenericContactName,omitempty"`
}

// EmailAddressData
type DomainEmailAddressData struct {
	EmailAddress string `json:"emailAddress,omitempty"`
	Lower        string `json:"lower,omitempty"`
	Primary      bool   `json:"primary,omitempty"`
	OptOut       bool   `json:"optOut,omitempty"`
	Invalid      bool   `json:"invalid,omitempty"`
}

// PhoneNumberData
type DomainPhoneNumberData struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Primary     bool   `json:"primary,omitempty"`
	Type        string `json:"type,omitempty"`
	OptOut      bool   `json:"optOut,omitempty"`
	Invalid     bool   `json:"invalid,omitempty"`
}

//  StatusExternal .
type DomainStatusExternal struct {
	Status string `json:"statusExternal,omitempty"`
}

// GenericParentId
type DomainGenericParentId struct {
	GenericParentId   string   `json:"accountParentId,omitempty"`
	GenericParentName string   `json:"accountParentName,omitempty"`
	TeamsIds          []string `json:"teamsIds,omitempty"`
}
