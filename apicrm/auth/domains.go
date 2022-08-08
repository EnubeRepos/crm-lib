package user

const EntityAuth string = "App/user"

// DomainAuth
type DomainAuth struct {
	User     User   `json:"user"`
	Token    string `json:"token"`
	Language string `json:"language"`
}

type User struct {
	ID       string `json:"id"`
	Deleted  bool   `json:"deleted"`
	UserName string `json:"userName"`
	IsActive bool   `json:"isActive"`
}
