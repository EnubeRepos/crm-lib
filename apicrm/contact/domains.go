package contact

const EntityContact string = "Contact"

// ResponseContact .
type ResponseContact struct {
	Total   int             `json:"total,omitempty"`
	Contact []DomainContact `json:"list,omitempty"`
}

// Account
type DomainContact struct{}
