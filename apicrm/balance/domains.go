package balance

const EntityBalance string = "Balance"

// ResponseBalance .
type ResponseBalance struct {
	Total   int             `json:"total,omitempty"`
	Balance []DomainBalance `json:"list,omitempty"`
}

// Account
type DomainBalance struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Deleted        bool    `json:"deleted"`
	CreatedAt      string  `json:"createdAt"`
	CurrentBalance float64 `json:"currentBalance"`
	BlockedBalance float64 `json:"blockedBalance"`
	AssignedUserID string  `json:"assignedUserId"`
	BankAccountID  string  `json:"bankAccountId"`
	ForceUpdate    bool    `json:"forceUpdate"`
}

type DomainBalanceCreateRequest struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	CurrentBalance float64 `json:"currentBalance"`
	BlockedBalance float64 `json:"blockedBalance"`
	AssignedUserID string  `json:"assignedUserId"`
	BankAccountID  string  `json:"bankAccountId"`
	ForceUpdate    bool    `json:"forceUpdate"`
}

type DomainBalanceCreateResponse struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Deleted        bool    `json:"deleted"`
	CreatedAt      string  `json:"createdAt"`
	CurrentBalance float64 `json:"currentBalance"`
	BlockedBalance float64 `json:"blockedBalance"`
	AssignedUserID string  `json:"assignedUserId"`
	BankAccountID  string  `json:"bankAccountId"`
	ForceUpdate    bool    `json:"forceUpdate"`
}
