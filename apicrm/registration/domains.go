package registration

const EntityRegistration string = "Registration"

// ResponseRegistration .
type ResponseRegistration struct {
	Total        int                  `json:"total,omitempty"`
	Registration []DomainRegistration `json:"list,omitempty"`
}

// Account
type DomainRegistration struct{}
