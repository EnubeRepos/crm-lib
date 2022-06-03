package userregistration

const EntityUserRegistration string = "UserRegistration"

// ResponseUserRegistration .
type ResponseUserRegistration struct {
	Total             int                      `json:"total,omitempty,omitempty"`
	UserRegistrations []DomainUserRegistration `json:"list,omitempty,omitempty"`
}

type DomainUserRegistration struct {
	ID                     string `json:"id,omitempty"`
	Name                   string `json:"name,omitempty"`
	Deleted                bool   `json:"deleted,omitempty"`
	Salutationname         string `json:"salutationName,omitempty"`
	Firstname              string `json:"firstName,omitempty"`
	Lastname               string `json:"lastName,omitempty"`
	Description            string `json:"description,omitempty"`
	Emailaddress           string `json:"emailAddress,omitempty"`
	Phonenumber            string `json:"phoneNumber,omitempty"`
	Addressstreet          string `json:"addressStreet,omitempty"`
	Addresscity            string `json:"addressCity,omitempty"`
	Addressstate           string `json:"addressState,omitempty"`
	Addresscountry         string `json:"addressCountry,omitempty"`
	Addresspostalcode      string `json:"addressPostalCode,omitempty"`
	Createdat              string `json:"createdAt,omitempty"`
	Modifiedat             string `json:"modifiedAt,omitempty"`
	Profiletype            string `json:"profileType,omitempty"`
	Siccode                string `json:"sicCode,omitempty"`
	Username               string `json:"userName,omitempty"`
	Isactive               bool   `json:"isActive,omitempty"`
	Disable                bool   `json:"disable,omitempty"`
	Temporarypassword      string `json:"temporaryPassword,omitempty"`
	Middlename             string `json:"middleName,omitempty"`
	Emailaddressisoptedout bool   `json:"emailAddressIsOptedOut,omitempty"`
	Phonenumberisoptedout  bool   `json:"phoneNumberIsOptedOut,omitempty"`
	Createdbyid            string `json:"createdById,omitempty"`
	Createdbyname          string `json:"createdByName,omitempty"`
	Modifiedbyid           string `json:"modifiedById,omitempty"`
	Modifiedbyname         string `json:"modifiedByName,omitempty"`
	Assigneduserid         string `json:"assignedUserId,omitempty"`
	Assignedusername       string `json:"assignedUserName,omitempty"`
	Createuserid           string `json:"createUserId,omitempty"`
	Createusername         string `json:"createUserName,omitempty"`
	Accountid              string `json:"accountId,omitempty"`
	Accountname            string `json:"accountName,omitempty"`
}
