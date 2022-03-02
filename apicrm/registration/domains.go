package registration

const EntityRegistration string = "Registration"

// ResponseRegistration .
type ResponseRegistration struct {
	Total        int                  `json:"total,omitempty"`
	Registration []DomainRegistration `json:"list,omitempty"`
}

// Account
type DomainRegistration struct {
	ID                         string             `json:"id"`
	Name                       string             `json:"name"`
	Deleted                    bool               `json:"deleted"`
	Description                string             `json:"description"`
	Website                    string             `json:"website"`
	EmailAddress               string             `json:"emailAddress"`
	PhoneNumber                string             `json:"phoneNumber"`
	BillingAddressStreet       string             `json:"billingAddressStreet"`
	BillingAddressCity         string             `json:"billingAddressCity"`
	BillingAddressState        string             `json:"billingAddressState"`
	BillingAddressCountry      string             `json:"billingAddressCountry"`
	BillingAddressPostalCode   string             `json:"billingAddressPostalCode"`
	ShippingAddressStreet      string             `json:"shippingAddressStreet"`
	ShippingAddressCity        string             `json:"shippingAddressCity"`
	ShippingAddressState       string             `json:"shippingAddressState"`
	ShippingAddressCountry     string             `json:"shippingAddressCountry"`
	ShippingAddressPostalCode  string             `json:"shippingAddressPostalCode"`
	CreatedAt                  string             `json:"createdAt"`
	ModifiedAt                 string             `json:"modifiedAt"`
	DocumentNumber             string             `json:"documentNumber"`
	RegistrationStatus         string             `json:"registrationStatus"`
	RegistrationType           string             `json:"registrationType"`
	SocialName                 string             `json:"socialName"`
	BusinessType               string             `json:"businessType"`
	BusinessSize               string             `json:"businessSize"`
	BirthDate                  string             `json:"birthDate"`
	MotherName                 string             `json:"motherName"`
	IsPoliticallyExposedPerson bool               `json:"isPoliticallyExposedPerson"`
	RegistrationCompleted      bool               `json:"registrationCompleted"`
	BankPrivacyPolicyApproval  bool               `json:"bankPrivacyPolicyApproval"`
	BankPrivacyPolicyURL       string             `json:"bankPrivacyPolicyUrl"`
	EmailAddressIsOptedOut     bool               `json:"emailAddressIsOptedOut"`
	PhoneNumberIsOptedOut      bool               `json:"phoneNumberIsOptedOut"`
	EmailAddressData           []EmailAddressData `json:"emailAddressData"`
	PhoneNumberData            []PhoneNumberData  `json:"phoneNumberData"`
	CreatedByID                string             `json:"createdById"`
	CreatedByName              string             `json:"createdByName"`
	ModifiedByID               string             `json:"modifiedById"`
	ModifiedByName             string             `json:"modifiedByName"`
	AssignedUserID             string             `json:"assignedUserId"`
	AssignedUserName           string             `json:"assignedUserName"`
	TeamsIds                   []string           `json:"teamsIds"`
	PictureID                  string             `json:"pictureId"`
	PictureName                string             `json:"pictureName"`
	BankAccountID              string             `json:"bankAccountId"`
	BankAccountName            string             `json:"bankAccountName"`
	UserID                     string             `json:"userId"`
	UserName                   string             `json:"userName"`
	ContactsIds                []string           `json:"contactsIds"`
	ContactsNames              ContactsNames      `json:"contactsNames"`
	IDDocumentFrontID          string             `json:"idDocumentFrontId"`
	IDDocumentFrontName        string             `json:"idDocumentFrontName"`
	IDDocumentBackID           string             `json:"idDocumentBackId"`
	IDDocumentBackName         string             `json:"idDocumentBackName"`
	BillingAddressProofID      string             `json:"billingAddressProofId"`
	BillingAddressProofName    string             `json:"billingAddressProofName"`
	SocialContractID           string             `json:"socialContractId"`
	SocialContractName         string             `json:"socialContractName"`
	LegalRepresentativeID      string             `json:"legalRepresentativeId"`
	LegalRepresentativeName    string             `json:"legalRepresentativeName"`
	IsFollowed                 bool               `json:"isFollowed"`
	FollowersIds               []string           `json:"followersIds"`
	FollowersNames             FollowersNames     `json:"followersNames"`
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
type ContactsNames struct {
}
type FollowersNames struct {
	Num1 string `json:"1"`
}
