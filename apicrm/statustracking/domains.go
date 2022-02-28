package statustracking

const EntityStatusTracking string = "StatusTracking"

// ResponseStatusTracking .
type ResponseStatusTracking struct {
	Total        int                  `json:"total,omitempty"`
	StatusTracking []DomainStatusTracking `json:"list,omitempty"`
}

// Account
type DomainStatusTracking struct{}
