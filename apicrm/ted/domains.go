package ted

const EntityTed string = "Ted"

// ResponseTed .
type ResponseTed struct {
	Total int         `json:"total,omitempty"`
	Ted   []DomainTed `json:"list,omitempty"`
}

// Request Ted
type DomainTed struct {
	ID                        string   `json:"id,omitempty"`
	Name                      string   `json:"name,omitempty"`
	Deleted                   bool     `json:"deleted,omitempty"`
	Description               string   `json:"description,omitempty"`
	Amount                    float64  `json:"amount,omitempty"`
	Agency                    string   `json:"agency,omitempty"`
	Account                   string   `json:"account,omitempty"`
	ConfirmTransaction        bool     `json:"confirmTransaction,omitempty"`
	OperationFeeAmount        float64  `json:"operationFeeAmount,omitempty"`
	InstitutionFinancialID    string   `json:"institutionFinancialId,omitempty"`
	FinTransactionsIds        []string `json:"finTransactionsIds,omitempty"`
	OperationFeeApplicationID string   `json:"operationFeeApplicationId,omitempty"`
	PersonalContactsID        string   `json:"personalContactsId,omitempty"`
	AccountType               string   `json:"accountType,omitempty"`
	ContactID                 string   `json:"contactId,omitempty"`
	AuthenticationCode        string   `json:"authenticationCode,omitempty"`

	// Remetente-Sender
	AssignedUserID string `json:"assignedUserId,omitempty"`

	// Beneficiário
	RecipientID string `json:"recipientId,omitempty"`

	Status       string `json:"status,omitempty"`
	StatusDetail string `json:"statusDetail,omitempty"`
}

type DomainTedPutStatus struct {
	ID                 string `json:"id,omitempty"`
	ConfirmTransaction bool   `json:"confirmTransaction,omitempty"`
	Status             string `json:"status,omitempty"`
	StatusDetail       string `json:"statusDetail,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
}
