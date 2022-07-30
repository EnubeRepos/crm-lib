package registration

const EntityRegistration string = "Registration"

// ResponseRegistration .
type ResponseRegistration struct {
	Total        int                  `json:"total,omitempty"`
	Registration []DomainRegistration `json:"list,omitempty"`
}

// estrutura usada para o put de registration, evitando erro de campos vazios //original
type DomainRegistrationBase struct {
	ID                 string `json:"id,omitempty"`
	StatusDetail       string `json:"statusDetail,omitempty"`
	StatusCid          string `json:"statusCid,omitempty"`
	StatusProcess      string `json:"statusProcess,omitempty"`
	StatusDatetime     string `json:"statusDatetime,omitempty"`
	RegistrationStatus string `json:"registrationStatus,omitempty"`
	AssignedUserID     string `json:"assignedUserId,omitempty"`
	BankAccountID      string `json:"bankAccountId,omitempty"`
	EmailAddress       string `json:"emailAddress,omitempty"`       //
	BillingAddressCity string `json:"billingAddressCity,omitempty"` //
}

// Registration
type DomainRegistration struct {
	DomainRegistrationBase
	ID                         string             `json:"id,omitempty"`
	Name                       string             `json:"name,omitempty"`
	Deleted                    bool               `json:"deleted"`
	Description                string             `json:"description,omitempty"`
	Website                    string             `json:"website,omitempty"`
	EmailAddress               string             `json:"emailAddress,omitempty"`
	PhoneNumber                string             `json:"phoneNumber,omitempty"`
	BillingAddressStreet       string             `json:"billingAddressStreet,omitempty"`
	BillingAddressCity         string             `json:"billingAddressCity,omitempty"`
	BillingAddressState        string             `json:"billingAddressState,omitempty"`
	BillingAddressCountry      string             `json:"billingAddressCountry,omitempty"`
	BillingAddressPostalCode   string             `json:"billingAddressPostalCode,omitempty"`
	BillingAddressNumber       string             `json:"billingAddressNumber,omitempty"`
	BillingAddressComplement   string             `json:"billingAddressComplement,omitempty"`
	BillingAddressNeighborhood string             `json:"billingAddressNeighborhood,omitempty"`
	ShippingAddressStreet      string             `json:"shippingAddressStreet,omitempty"`
	ShippingAddressCity        string             `json:"shippingAddressCity,omitempty"`
	ShippingAddressState       string             `json:"shippingAddressState,omitempty"`
	ShippingAddressCountry     string             `json:"shippingAddressCountry,omitempty"`
	ShippingAddressPostalCode  string             `json:"shippingAddressPostalCode,omitempty"`
	CreatedAt                  string             `json:"createdAt,omitempty"`
	ModifiedAt                 string             `json:"modifiedAt,omitempty"`
	DocumentNumber             string             `json:"documentNumber,omitempty"`
	RegistrationType           string             `json:"registrationType,omitempty"`
	SocialName                 string             `json:"socialName,omitempty"`
	BusinessType               string             `json:"businessType,omitempty"`
	BusinessSize               string             `json:"businessSize,omitempty"`
	BusinessRegistrationID     string             `json:"businessRegistrationId,omitempty"`
	BirthDate                  string             `json:"birthDate,omitempty"`
	MotherName                 string             `json:"motherName,omitempty"`
	IsPoliticallyExposedPerson bool               `json:"isPoliticallyExposedPerson"`
	RegistrationCompleted      bool               `json:"registrationCompleted"`
	BankPrivacyPolicyApproval  bool               `json:"bankPrivacyPolicyApproval"`
	BankPrivacyPolicyURL       string             `json:"bankPrivacyPolicyUrl,omitempty"`
	EmailAddressIsOptedOut     bool               `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut      bool               `json:"phoneNumberIsOptedOut,omitempty"`
	EmailAddressData           []EmailAddressData `json:"emailAddressData,omitempty"`
	PhoneNumberData            []PhoneNumberData  `json:"phoneNumberData,omitempty"`
	CreatedByID                string             `json:"createdById,omitempty"`
	CreatedByName              string             `json:"createdByName,omitempty"`
	ModifiedByID               string             `json:"modifiedById,omitempty"`
	ModifiedByName             string             `json:"modifiedByName,omitempty"`
	AssignedUserID             string             `json:"assignedUserId,omitempty"`
	AssignedUserName           string             `json:"assignedUserName,omitempty"`
	TeamsIds                   []string           `json:"teamsIds,omitempty"`
	PictureID                  string             `json:"pictureId,omitempty"`
	PictureName                string             `json:"pictureName,omitempty"`
	BankAccountID              string             `json:"bankAccountId,omitempty"`
	BankAccountName            string             `json:"bankAccountName,omitempty"`
	UserID                     string             `json:"userId,omitempty"`
	UserName                   string             `json:"userName,omitempty"`
	ContactsIds                []string           `json:"contactsIds,omitempty"`
	ContactsNames              ContactsNames      `json:"contactsNames,omitempty"`
	IDDocumentFrontID          string             `json:"idDocumentFrontId,omitempty"`
	IDDocumentFrontName        string             `json:"idDocumentFrontName,omitempty"`
	IDDocumentBackID           string             `json:"idDocumentBackId,omitempty"`
	IDDocumentBackName         string             `json:"idDocumentBackName,omitempty"`
	BillingAddressProofID      string             `json:"billingAddressProofId,omitempty"`
	BillingAddressProofName    string             `json:"billingAddressProofName,omitempty"`
	RegistrationRequestsID     []string           `json:"registrationRequestsIds,omitempty"`
	SocialContractID           string             `json:"socialContractId,omitempty"`
	SocialContractName         string             `json:"socialContractName,omitempty"`
	LegalRepresentativeID      string             `json:"legalRepresentativeId,omitempty"`
	LegalRepresentativeName    string             `json:"legalRepresentativeName,omitempty"`
	IsFollowed                 bool               `json:"isFollowed"`
	FollowersIds               []string           `json:"followersIds,omitempty"`
	FollowersNames             FollowersNames     `json:"followersNames,omitempty"`
	TemporaryPassword          string             `json:"temporaryPassword,omitempty"`
}

type EmailAddressData struct {
	EmailAddress string `json:"emailAddress,omitempty"`
	Lower        string `json:"lower,omitempty"`
	Primary      bool   `json:"primary"`
	OptOut       bool   `json:"optOut"`
	Invalid      bool   `json:"invalid"`
}
type PhoneNumberData struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Type        string `json:"type,omitempty"`
	Primary     bool   `json:"primary"`
	OptOut      bool   `json:"optOut"`
	Invalid     bool   `json:"invalid"`
}
type ContactsNames struct {
}
type FollowersNames struct {
	Num1 string `json:"1,omitempty"`
}
