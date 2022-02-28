package fintransaction

const EntityFinTransaction string = "FinTransaction"

// ResponseFinTransaction .
type ResponseFinTransaction struct {
	Total        int                  `json:"total,omitempty"`
	FinTransaction []DomainFinTransaction `json:"list,omitempty"`
}

// Account
type DomainFinTransaction struct{}
