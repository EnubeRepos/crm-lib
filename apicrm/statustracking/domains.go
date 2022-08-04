package statustracking

const EntityStatusTracking string = "StatusTracking"

// ResponseStatusTracking .
type ResponseStatusTracking struct {
	Total          int                    `json:"total,omitempty"`
	StatusTracking []DomainStatusTracking `json:"list,omitempty"`
}

// estrutura usada para o put de registration, evitando erro de campos vazios
type DomainStatusTrackingBase struct {
	ID         string `json:"id,omitempty"`
	StatusType string `json:"statusType,omitempty"`
	Code       int    `json:"code,omitempty"`
}

// Account
type DomainStatusTracking struct {
	ID               string `json:"id,omitempty"`
	Code             int    `json:"code,omitempty"`
	AssignedUser     string `json:"assignedUser,omitempty"`
	AssignedUserID   string `json:"assignedUserId,omitempty"`
	AssignedUserName string `json:"assignedUserName,omitempty"`
}
