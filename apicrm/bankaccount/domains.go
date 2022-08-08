package bankaccount

const EntityBankAccount string = "BankAccount"

// ResponseBankAccount .
type ResponseBankAccount struct {
	Total       int                 `json:"total,omitempty"`
	BankAccount []DomainBankAccount `json:"list,omitempty"`
}

// Account
type DomainBankAccount struct {
	ID                     string   `json:"id"`
	Name                   string   `json:"name"`
	Deleted                bool     `json:"deleted"`
	SalutationName         string   `json:"salutationName"`
	FirstName              string   `json:"firstName"`
	LastName               string   `json:"lastName"`
	Description            string   `json:"description"`
	EmailAddress           string   `json:"emailAddress"`
	PhoneNumber            string   `json:"phoneNumber"`
	AddressStreet          string   `json:"addressStreet"`
	AddressCity            string   `json:"addressCity"`
	AddressState           string   `json:"addressState"`
	AddressCountry         string   `json:"addressCountry"`
	AddressPostalCode      string   `json:"addressPostalCode"`
	CreatedAt              string   `json:"createdAt"`
	ModifiedAt             string   `json:"modifiedAt"`
	SicCode                string   `json:"sicCode"`
	AccountNumber          string   `json:"accountNumber"`
	ExpirationDate         string   `json:"expirationDate"`
	Agency                 string   `json:"agency"`
	PixKey                 string   `json:"pixKey"`
	BankCode               string   `json:"bankCode"`
	MiddleName             string   `json:"middleName"`
	EmailAddressIsOptedOut string   `json:"emailAddressIsOptedOut"`
	PhoneNumberIsOptedOut  string   `json:"phoneNumberIsOptedOut"`
	EmailAddressData       []string `json:"emailAddressData"`
	PhoneNumberData        []string `json:"phoneNumberData"`
	CreatedByID            string   `json:"createdById"`
	CreatedByName          string   `json:"createdByName"`
	ModifiedByID           string   `json:"modifiedById"`
	ModifiedByName         string   `json:"modifiedByName"`
	AssignedUserID         string   `json:"assignedUserId"`
	AssignedUserName       string   `json:"assignedUserName"`
	TeamsIds               []string `json:"teamsIds"`
	FinTransactionsIds     []string `json:"finTransactionsIds"`
	RegistrationID         string   `json:"registrationId"`
	RegistrationName       string   `json:"registrationName"`
	AccountID              string   `json:"accountId"`
	AccountName            string   `json:"accountName"`
	IntegratorID           string   `json:"integratorId"`
	IntegratorName         string   `json:"integratorName"`
	ParcelssIds            []string `json:"parcelssIds"`
	UserID                 string   `json:"userId"`
	UserName               string   `json:"userName"`
	InstitutionFinancialID string   `json:"institutionFinancialId"`
	BalanceId              string   `json:"balanceId"`
}

type DomainBankAccountCreateRequest struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	FirstName              string `json:"firstName"`
	LastName               string `json:"lastName"`
	SicCode                string `json:"sicCode"`
	AccountNumber          string `json:"accountNumber"`
	Agency                 string `json:"agency"`
	AssignedUserID         string `json:"assignedUserId"`
	InstitutionFinancialID string `json:"institutionFinancialId"`
}

type DomainBankAccountCreateResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
