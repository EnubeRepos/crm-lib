package bankaccountbalances

const EntityBankAccountBalance string = "BankAccountBalance"

// ResponseBankAccountBalance .
type ResponseBankAccountBalance struct {
	Total              int                        `json:"total,omitempty"`
	BankAccountBalance []DomainBankAccountBalance `json:"list,omitempty"`
}

// Account
type DomainBankAccountBalance struct {
	ID                      string  `json:"id"`
	Name                    string  `json:"name"`
	Deleted                 bool    `json:"deleted"`
	Description             string  `json:"description"`
	ValueAvailable          float64 `json:"valueAvailable"`
	ValueInProcess          float64 `json:"valueInProcess"`
	ValueBlocked            float64 `json:"valueBlocked"`
	ValueSumVirtual         float64 `json:"valueSumVirtual"`
	ValueAvailableCurrency  string  `json:"valueAvailableCurrency"`
	ValueInProcessCurrency  string  `json:"valueInProcessCurrency"`
	ValueBlockedCurrency    string  `json:"valueBlockedCurrency"`
	ValueSumVirtualCurrency string  `json:"valueSumVirtualCurrency"`
	BankAccountId           string  `json:"bankAccountId"`
	BankAccountName         string  `json:"bankAccountName"`
	//AssignedUser            string  `json:"assignedUser"` // i added
	//AssignedUserId          string  `json:"assignedUserId"`
	//AssignedUserName        string  `json:"assignedUserName"`

	//CreatedAt                string        `json:"createdAt"`
	//ModifiedAt               string        `json:"modifiedAt"`
	//CreatedById              string        `json:"createdById"`
	//CreatedByName            string        `json:"createdByName"`
	//ModifiedById             string        `json:"modifiedById"`
	//ModifiedByName           string        `json:"modifiedByName"`

	//TeamsIds                 []string `json:"teamsIds"`
	//TeamsNames               Teamsnames    `json:"teamsNames"`
	//ValueAvailableConverted  float64           `json:"valueAvailableConverted"`
	//ValueInProcessConverted  float64           `json:"valueInProcessConverted"`
	//ValueBlockedConverted    float64           `json:"valueBlockedConverted"`
	//ValueSumVirtualConverted float64           `json:"valueSumVirtualConverted"`
}

type DomainBankAccountBalanceCreateRequest struct {
	ID               string  `json:"id"`
	BankAccountId    string  `json:"bankAccountId"`
	BankAccountName  string  `json:"bankAccountName"`
	ValueAvailable   float64 `json:"valueAvailable"`
	ValueInProcess   float64 `json:"valueInProcess"`
	ValueBlocked     float64 `json:"valueBlocked"`
	AssignedUser     string  `json:"assignedUser"`     // I added
	AssignedUserId   string  `json:"assignedUserId"`   // I
	AssignedUserName string  `json:"assignedUserName"` // I
}

type DomainBankAccountBalanceCreateResponse struct {
	ID             string  `json:"id"`
	BankAccountId  string  `json:"bankAccountId"`
	ValueAvailable float64 `json:"valueAvailable"`
	ValueInProcess float64 `json:"valueInProcess"`
	ValueBlocked   float64 `json:"valueBlocked"`
}
