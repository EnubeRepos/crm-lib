package bankaccount

const EntityBankAccount string = "BankAccount"

// ResponseBankAccount .
type ResponseBankAccount struct {
	Total        int                  `json:"total,omitempty"`
	BankAccount []DomainBankAccount `json:"list,omitempty"`
}

// Account
type DomainBankAccount struct{}
