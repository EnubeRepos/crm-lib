package statustracking

const EntityStatusTracking string = "StatusTracking"

// ResponseStatusTracking .
type ResponseStatusTracking struct {
	Total          int                    `json:"total,omitempty"`
	StatusTracking []DomainStatusTracking `json:"list,omitempty"`
}

// Account
type DomainStatusTracking struct { //didn't have any fields
	ID               string `json:"id,omitempty"`
	Code             int    `json:"code,omitempty"`
	AssignedUser     string `json:"assignedUser,omitempty"`
	AssignedUserID   string `json:"assignedUserId,omitempty"`
	AssignedUserName string `json:"assignedUserName,omitempty"`
}
