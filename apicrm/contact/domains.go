package contact

const EntityContact string = "Contact"

// ResponseContact .
type ResponseContact struct {
	Total   int             `json:"total,omitempty"`
	Contact []DomainContact `json:"list,omitempty"`
}

// Account
type DomainContact struct {
	ID                       string             `json:"id,omitempty"`
	Name                     string             `json:"name,omitempty"`
	Deleted                  bool               `json:"deleted,omitempty"`
	SalutationName           string             `json:"salutationName,omitempty"`
	FirstName                string             `json:"firstName,omitempty"`
	LastName                 string             `json:"lastName,omitempty"`
	Title                    string             `json:"title,omitempty"`
	Description              string             `json:"description,omitempty"`
	EmailAddress             string             `json:"emailAddress,omitempty"`
	PhoneNumber              string             `json:"phoneNumber,omitempty"`
	DoNotCall                bool               `json:"doNotCall,omitempty"`
	AddressStreet            string             `json:"addressStreet,omitempty"`
	AddressCity              string             `json:"addressCity,omitempty"`
	AddressState             string             `json:"addressState,omitempty"`
	AddressCountry           string             `json:"addressCountry,omitempty"`
	AddressPostalCode        string             `json:"addressPostalCode,omitempty"`
	AccountIsInactive        bool               `json:"accountIsInactive,omitempty"`
	AccountType              string             `json:"accountType,omitempty"`
	CreatedAt                string             `json:"createdAt,omitempty"`
	ModifiedAt               string             `json:"modifiedAt,omitempty"`
	HasPortalUser            bool               `json:"hasPortalUser,omitempty"`
	SicCode                  string             `json:"sicCode,omitempty"`
	ContactAccountNumber     string             `json:"contactAccountNumber,omitempty"`
	ContactAgency            string             `json:"contactAgency,omitempty"`
	MiddleName               string             `json:"middleName,omitempty"`
	EmailAddressIsOptedOut   bool               `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut    bool               `json:"phoneNumberIsOptedOut,omitempty"`
	EmailAddressData         []Emailaddressdata `json:"emailAddressData,omitempty"`
	AccountID                string             `json:"accountId,omitempty"`
	AccountName              string             `json:"accountName,omitempty"`
	AccountsIds              []string           `json:"accountsIds,omitempty"`
	CampaignID               string             `json:"campaignId,omitempty"`
	CampaignName             string             `json:"campaignName,omitempty"`
	CreatedByID              string             `json:"createdById,omitempty"`
	CreatedByName            string             `json:"createdByName,omitempty"`
	ModifiedByID             string             `json:"modifiedById,omitempty"`
	ModifiedByName           string             `json:"modifiedByName,omitempty"`
	AssignedUserID           string             `json:"assignedUserId,omitempty"`
	AssignedUserName         string             `json:"assignedUserName,omitempty"`
	TeamsIds                 []string           `json:"teamsIds,omitempty"`
	PortalUserID             string             `json:"portalUserId,omitempty"`
	PortalUserName           string             `json:"portalUserName,omitempty"`
	OriginalLeadID           string             `json:"originalLeadId,omitempty"`
	OriginalLeadName         string             `json:"originalLeadName,omitempty"`
	SalesOrdersIds           []string           `json:"salesOrdersIds,omitempty"`
	SalesOrderItemsIds       []string           `json:"salesOrderItemsIds,omitempty"`
	SalesBillingContactsIds  []string           `json:"salesBillingContactsIds,omitempty"`
	SalesBuyerContactsIds    []string           `json:"salesBuyerContactsIds,omitempty"`
	CommissionsIds           []string           `json:"commissionsIds,omitempty"`
	ParcelsIds               []string           `json:"parcelsIds,omitempty"`
	ParcelssIds              []string           `json:"parcelssIds,omitempty"`
	FinTransactionsIds       []string           `json:"finTransactionsIds,omitempty"`
	InstitutionFinancialID   string             `json:"institutionFinancialId,omitempty"`
	InstitutionFinancialName string             `json:"institutionFinancialName,omitempty"`
	TedsIds                  []string           `json:"tedsIds,omitempty"`
	PixsIds                  []string           `json:"pixsIds,omitempty"`
	PixKeysIds               []string           `json:"pixKeysIds,omitempty"`
	IsFollowed               bool               `json:"isFollowed,omitempty"`
	FollowersIds             []string           `json:"followersIds,omitempty"`
	//PhoneNumberData          []string `json:"phoneNumberData,omitempty"`

}
type Emailaddressdata struct {
	Emailaddress string `json:"emailAddress"`
	Lower        string `json:"lower"`
	Primary      bool   `json:"primary"`
	Optout       bool   `json:"optOut"`
	Invalid      bool   `json:"invalid"`
}
