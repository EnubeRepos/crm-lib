package accreditedaccounts

const EntityAccreditedAccounts string = "AccreditedAccounts"

// ResponseAccreditedAccounts .
type ResponseAccreditedAccounts struct {
	Total              int                        `json:"total,omitempty"`
	AccreditedAccounts []DomainAccreditedAccounts `json:"list,omitempty"`
}

// Account
type DomainAccreditedAccounts struct{}
