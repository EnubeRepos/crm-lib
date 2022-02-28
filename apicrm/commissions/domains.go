package commissions

const EntityCommissions string = "Commissions"

// ResponseCommissions .
type ResponseCommissions struct {
	Total       int                 `json:"total,omitempty"`
	Commissions []DomainCommissions `json:"list,omitempty"`
}

// Account
type DomainCommissions struct{}
