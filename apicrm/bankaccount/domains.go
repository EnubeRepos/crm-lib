package bankaccount

const EntityBankAccount string = "BankAccount"

// ResponseBankAccount .
type ResponseBankAccount struct {
	Total       int                 `json:"total,omitempty"`
	BankAccount []DomainBankAccount `json:"list,omitempty"`
}

// Account
type DomainBankAccount struct {
	ID                       string   `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Deleted                  bool     `json:"deleted,omitempty"`
	SalutationName           string   `json:"salutationName,omitempty"`
	FirstName                string   `json:"firstName,omitempty"`
	LastName                 string   `json:"lastName,omitempty"`
	Description              string   `json:"description,omitempty"`
	EmailAddress             string   `json:"emailAddress,omitempty"`
	PhoneNumber              string   `json:"phoneNumber,omitempty"`
	AddressStreet            string   `json:"addressStreet,omitempty"`
	AddressCity              string   `json:"addressCity,omitempty"`
	AddressState             string   `json:"addressState,omitempty"`
	AddressCountry           string   `json:"addressCountry,omitempty"`
	AddressPostalCode        string   `json:"addressPostalCode,omitempty"`
	CreatedAt                string   `json:"createdAt,omitempty"`
	ModifiedAt               string   `json:"modifiedAt,omitempty"`
	SicCode                  string   `json:"sicCode,omitempty"`
	AccountNumber            string   `json:"accountNumber,omitempty"`
	ExpirationDate           string   `json:"expirationDate,omitempty"`
	Agency                   string   `json:"agency,omitempty"`
	RegistrationType         string   `json:"registrationType,omitempty"`
	MiddleName               string   `json:"middleName,omitempty"`
	EmailAddressIsOptedOut   string   `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut    string   `json:"phoneNumberIsOptedOut,omitempty"`
	EmailAddressData         []string `json:"emailAddressData,omitempty"`
	PhoneNumberData          []string `json:"phoneNumberData,omitempty"`
	CreatedByID              string   `json:"createdById,omitempty"`
	CreatedByName            string   `json:"createdByName,omitempty"`
	ModifiedByID             string   `json:"modifiedById,omitempty"`
	ModifiedByName           string   `json:"modifiedByName,omitempty"`
	AssignedUserID           string   `json:"assignedUserId,omitempty"`
	AssignedUserName         string   `json:"assignedUserName,omitempty"`
	TeamsIds                 []string `json:"teamsIds,omitempty"`
	FinTransactionsIds       []string `json:"finTransactionsIds,omitempty"`
	RegistrationID           string   `json:"registrationId,omitempty"`
	RegistrationName         string   `json:"registrationName,omitempty"`
	AccountID                string   `json:"accountId,omitempty"`
	AccountName              string   `json:"accountName,omitempty"`
	IntegratorID             string   `json:"integratorId,omitempty"`
	IntegratorName           string   `json:"integratorName,omitempty"`
	UserID                   string   `json:"userId,omitempty"`
	UserName                 string   `json:"userName,omitempty"`
	PixKeysIds               []string `json:"pixKeysIds,omitempty"`
	InstitutionFinancialID   string   `json:"institutionFinancialId,omitempty"`
	InstitutionFinancialName string   `json:"institutionFinancialName,omitempty"`
	BankAccountBalanceID     string   `json:"bankAccountBalanceId,omitempty"`
	BankAccountBalanceName   string   `json:"bankAccountBalanceName,omitempty"`
	UserRegistrationID       string   `json:"userRegistrationId,omitempty"`
	UserRegistrationName     string   `json:"userRegistrationName,omitempty"`
	BalanceID                string   `json:"balanceId,omitempty"`
	BalanceName              string   `json:"balanceName,omitempty"`
	LimitManagementID        string   `json:"limitManagementId,omitempty"`
	LimitManagementName      string   `json:"limitManagementName,omitempty"`
	IsFollowed               bool     `json:"isFollowed,omitempty"`
	FollowersIds             []string `json:"followersIds,omitempty"`
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
