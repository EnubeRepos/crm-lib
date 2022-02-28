package user

const EntityUser string = "User"

// ResponseUser .
type ResponseUser struct {
	Total        int                  `json:"total,omitempty"`
	User []DomainUser `json:"list,omitempty"`
}

// Account
type DomainUser struct{}
