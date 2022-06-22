package accreditedaccounts

const EntityAccreditedAccounts string = "AccreditedAccounts"

// ResponseAccreditedAccounts .
type ResponseAccreditedAccounts struct {
	Total              int                        `json:"total,omitempty"`
	AccreditedAccounts []DomainAccreditedAccounts `json:"list,omitempty"`
}

// Account
type DomainAccreditedAccounts struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	AssignedUser     int    `json:"assignedUser"`
	AssignedUserId   string `json:"assignedUserId"`
	AssignedUserName string `json:"assignedUserName"`
}
