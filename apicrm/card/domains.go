package card

const EntityDebitCard string = "DebitCard"
const EntityVirtualCard string = "VirtualCard"

// ResponseDomainDebitCard .
type ResponseDebitCard struct {
	Total     int               `json:"total,omitempty"`
	DebitCard []DomainDebitCard `json:"list,omitempty"`
}

// DomainDebitCard
type DomainDebitCard struct {
	ID                         string   `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Deleted                    bool     `json:"deleted,omitempty"`
	Description                string   `json:"description,omitempty"`
	CreatedAt                  string   `json:"createdAt,omitempty"`
	ModifiedAt                 string   `json:"modifiedAt,omitempty"`
	BillingAddressCity         string   `json:"billingAddressCity,omitempty"`
	BillingAddressComplement   string   `json:"billingAddressComplement,omitempty"`
	BillingAddressCountry      string   `json:"billingAddressCountry,omitempty"`
	BillingAddressNeighborhood string   `json:"billingAddressNeighborhood,omitempty"`
	BillingAddressNumber       string   `json:"billingAddressNumber,omitempty"`
	BillingAddressPostalCode   string   `json:"billingAddressPostalCode,omitempty"`
	BillingAddressState        string   `json:"billingAddressState,omitempty"`
	BillingAddressStreet       string   `json:"billingAddressStreet,omitempty"`
	DocumentNumber             string   `json:"documentNumber,omitempty"`
	Email                      string   `json:"email,omitempty"`
	PhoneNumber                string   `json:"phoneNumber,omitempty"`
	CardName                   string   `json:"cardName,omitempty"`
	Alias                      string   `json:"alias,omitempty"`
	ProgramID                  string   `json:"programId,omitempty"`
	RequisitionCompleted       bool     `json:"requisitionCompleted,omitempty"`
	AccountNumber              string   `json:"accountNumber,omitempty"`
	BankAgency                 string   `json:"bankAgency,omitempty"`
	OperationFeeAmount         int      `json:"operationFeeAmount,omitempty"`
	CardNumber                 string   `json:"cardNumber,omitempty"`
	ExpirationDate             string   `json:"expirationDate,omitempty"`
	CardVerificationValue      string   `json:"cardVerificationValue,omitempty"`
	Password                   string   `json:"password,omitempty"`
	Activate                   bool     `json:"activate,omitempty"`
	Disable                    bool     `json:"disable,omitempty"`
	ActivateCode               string   `json:"activateCode,omitempty"`
	OperationFeeAmountCurrency string   `json:"operationFeeAmountCurrency,omitempty"`
	CreatedByID                string   `json:"createdById,omitempty"`
	CreatedByName              string   `json:"createdByName,omitempty"`
	ModifiedByID               string   `json:"modifiedById,omitempty"`
	ModifiedByName             string   `json:"modifiedByName,omitempty"`
	AssignedUserID             string   `json:"assignedUserId,omitempty"`
	AssignedUserName           string   `json:"assignedUserName,omitempty"`
	TeamsIds                   []string `json:"teamsIds,omitempty"`
	BankAccountParentID        string   `json:"bankAccountParentId,omitempty"`
	BankAccountParentName      string   `json:"bankAccountParentName,omitempty"`
	OperationFeeApplicationID  string   `json:"operationFeeApplicationId,omitempty"`
	Status                     string   `json:"status,omitempty"`
	Proxy                      string   `json:"proxy,omitempty"`
	TedId                      string   `json:"tedId,omitempty"`
	FinTransactionsIds         []string `json:"finTransactionsIds,omitempty"`
	Cvv                        string   `json:"cvv,omitempty"`
}

// ResponseDomainDebitCard .
type ResponseVirtualCard struct {
	Total       int                 `json:"total,omitempty"`
	VirtualCard []DomainVirtualCard `json:"list,omitempty"`
}

// DomainVirtualCard
type DomainVirtualCard struct {
	ID                         string   `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Deleted                    bool     `json:"deleted,omitempty"`
	Description                string   `json:"description,omitempty"`
	CreatedAt                  string   `json:"createdAt,omitempty"`
	ModifiedAt                 string   `json:"modifiedAt,omitempty"`
	BillingAddressCity         string   `json:"billingAddressCity,omitempty"`
	BillingAddressComplement   string   `json:"billingAddressComplement,omitempty"`
	BillingAddressCountry      string   `json:"billingAddressCountry,omitempty"`
	BillingAddressNeighborhood string   `json:"billingAddressNeighborhood,omitempty"`
	BillingAddressNumber       string   `json:"billingAddressNumber,omitempty"`
	BillingAddressPostalCode   string   `json:"billingAddressPostalCode,omitempty"`
	BillingAddressState        string   `json:"billingAddressState,omitempty"`
	BillingAddressStreet       string   `json:"billingAddressStreet,omitempty"`
	DocumentNumber             string   `json:"documentNumber,omitempty"`
	Email                      string   `json:"email,omitempty"`
	PhoneNumber                string   `json:"phoneNumber,omitempty"`
	CardName                   string   `json:"cardName,omitempty"`
	Alias                      string   `json:"alias,omitempty"`
	ProgramID                  string   `json:"programId,omitempty"`
	RequisitionCompleted       bool     `json:"requisitionCompleted,omitempty"`
	AccountNumber              string   `json:"accountNumber,omitempty"`
	BankAgency                 string   `json:"bankAgency,omitempty"`
	OperationFeeAmount         int      `json:"operationFeeAmount,omitempty"`
	CardNumber                 string   `json:"cardNumber,omitempty"`
	ExpirationDate             string   `json:"expirationDate,omitempty"`
	CardVerificationValue      string   `json:"cardVerificationValue,omitempty"`
	Password                   string   `json:"password,omitempty"`
	Activate                   bool     `json:"activate,omitempty"`
	Disable                    bool     `json:"disable,omitempty"`
	ActivateCode               string   `json:"activateCode,omitempty"`
	OperationFeeAmountCurrency string   `json:"operationFeeAmountCurrency,omitempty"`
	CreatedByID                string   `json:"createdById,omitempty"`
	CreatedByName              string   `json:"createdByName,omitempty"`
	ModifiedByID               string   `json:"modifiedById,omitempty"`
	ModifiedByName             string   `json:"modifiedByName,omitempty"`
	AssignedUserID             string   `json:"assignedUserId,omitempty"`
	AssignedUserName           string   `json:"assignedUserName,omitempty"`
	TeamsIds                   []string `json:"teamsIds,omitempty"`
	BankAccountParentID        string   `json:"bankAccountParentId,omitempty"`
	BankAccountParentName      string   `json:"bankAccountParentName,omitempty"`
	OperationFeeApplicationID  string   `json:"operationFeeApplicationId,omitempty"`
	Status                     string   `json:"status,omitempty"`
	Proxy                      string   `json:"proxy,omitempty"`
	TedId                      string   `json:"tedId,omitempty"`
	Cvv                        string   `json:"cvv,omitempty"`
	FinTransactionsIds         []string `json:"finTransactionsIds,omitempty"`
}
