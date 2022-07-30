package personalcontact

const EntityPersonalContacts string = "PersonalContacts"

// ResponsePersonalContacts .
type ResponsePersonalContacts struct {
	Total            int                      `json:"total,omitempty"`
	PersonalContacts []DomainPersonalContacts `json:"list,omitempty"`
}

// PersonalContacts
type DomainPersonalContacts struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Deleted              bool   `json:"deleted,omitempty"`
	SalutationName       string `json:"salutationName,omitempty"`
	FirstName            string `json:"firstName,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	Description          string `json:"description,omitempty"`
	EmailAddress         string `json:"emailAddress,omitempty"`
	PhoneNumber          string `json:"phoneNumber,omitempty"`
	AddressStreet        string `json:"addressStreet,omitempty"`
	AddressCity          string `json:"addressCity,omitempty"`
	AddressState         string `json:"addressState,omitempty"`
	AddressCountry       string `json:"addressCountry,omitempty"`
	AddressPostalCode    string `json:"addressPostalCode,omitempty"`
	CreatedAt            string `json:"createdAt,omitempty"`
	ModifiedAt           string `json:"modifiedAt,omitempty"`
	ContactAccountNumber string `json:"contactAccountNumber,omitempty"`
	ContactAgency        string `json:"contactAgency,omitempty"`
	AccountType          string `json:"accountType,omitempty"`
	SicCode              string `json:"sicCode,omitempty"`
	MiddleName           string `json:"middleName,omitempty"`
	//EmailAddressIsOptedOut string   `json:"emailAddressIsOptedOut,omitempty"`
	//PhoneNumberIsOptedOut  string   `json:"phoneNumberIsOptedOut,omitempty"`
	//EmailAddressData       []string `json:"emailAddressData,omitempty"`
	//PhoneNumberData        []string `json:"phoneNumberData,omitempty"`
	CreatedByID            string   `json:"createdById,omitempty"`
	CreatedByName          string   `json:"createdByName,omitempty"`
	ModifiedByID           string   `json:"modifiedById,omitempty"`
	ModifiedByName         string   `json:"modifiedByName,omitempty"`
	AssignedUserID         string   `json:"assignedUserId,omitempty"`
	AssignedUserName       string   `json:"assignedUserName,omitempty"`
	TeamsIds               []string `json:"teamsIds,omitempty"`
	InstitutionFinancialID string   `json:"institutionFinancialId,omitempty"`
	PixKeiesIds            []string `json:"pixKeiesIds,omitempty"`
}
