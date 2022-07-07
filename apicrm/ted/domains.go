package ted

const EntityTed string = "Ted"

// ResponseTed .
type ResponseTed struct {
	Total int         `json:"total,omitempty"`
	Ted   []DomainTed `json:"list,omitempty"`
}

// Account
type DomainTed struct {
	ID                        string   `json:"id,omitempty"`
	Name                      string   `json:"name,omitempty"`
	Deleted                   bool     `json:"deleted,omitempty"`
	Description               string   `json:"description,omitempty"`
	Amount                    int      `json:"amount,omitempty"`
	Agency                    string   `json:"agency,omitempty"`
	Account                   string   `json:"account,omitempty"`
	ConfirmTransaction        bool     `json:"confirmTransaction,omitempty"`
	OperationFeeAmount        int      `json:"operationFeeAmount,omitempty"`
	InstitutionFinancialID    string   `json:"institutionFinancialId,omitempty"`
	FinTransactionsIds        []string `json:"finTransactionsIds,omitempty"`
	OperationFeeApplicationID string   `json:"operationFeeApplicationId,omitempty"`
	PersonalContactsID        string   `json:"personalContactsId,omitempty"`
	ContactID                 string   `json:"contactId,omitempty"`
	AssignedUserID            string   `json:"assignedUserId,omitempty"`
	BankAccountId             string   `json:"bankAccountId,omitempty"`
	//CreatedAt                   string               `json:"createdAt,omitempty"`
	//ModifiedAt                  string               `json:"modifiedAt,omitempty"`
	//ContactAgency               string               `json:"contactAgency,omitempty"`
	//ContactAccountNumber        string               `json:"contactAccountNumber,omitempty"`
	//SicCode                     string               `json:"sicCode,omitempty"`
	//ScheduleTransaction         string               `json:"scheduleTransaction,omitempty"`
	//DateStart                   string               `json:"dateStart,omitempty"`
	//AmountCurrency              string               `json:"amountCurrency,omitempty"`
	//OperationFeeAmountCurrency  string               `json:"operationFeeAmountCurrency,omitempty"`
	//CreatedByID                 string               `json:"createdById,omitempty"`
	//CreatedByName               string               `json:"createdByName,omitempty"`
	//ModifiedByID                string               `json:"modifiedById,omitempty"`
	//ModifiedByName              string               `json:"modifiedByName,omitempty"`
	//AssignedUserName            string               `json:"assignedUserName,omitempty"`
	//TeamsIds                    []string        `json:"teamsIds,omitempty"`
	//AmountConverted             int                  `json:"amountConverted,omitempty"`
	//InstitutionFinancialName    string               `json:"institutionFinancialName,omitempty"`
	//OperationFeeAmountConverted int                  `json:"operationFeeAmountConverted,omitempty"`
	//ContactName                 string               `json:"contactName,omitempty"`
	//OperationFeeApplicationName string               `json:"operationFeeApplicationName,omitempty"`
	//PersonalContactsName        string          `json:"personalContactsName,omitempty"`
}

type DomainTedPutConfirm struct {
	ID                 string `json:"id,omitempty"`
	ConfirmTransaction bool   `json:"confirmTransaction,omitempty"`
}
