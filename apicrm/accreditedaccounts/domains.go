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
	User             string `json:"user,omitempty"`
	Account          string `json:"account,omitempty"`
	Name             string `json:"name,omitempty"`
	AssignedUser     int    `json:"assignedUser,omitempty"`
	AssignedUserId   string `json:"assignedUserId,omitempty"`
	AssignedUserName string `json:"assignedUserName,omitempty"`
}
