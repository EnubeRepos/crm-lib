package user

const EntityAuth string = "App/user"

// DomainAuth
type DomainAuth struct {
	User        User   `json:"user"`
	Token       string `json:"token"`
	TokenSecret string `json:"tokenSecret"`
	Language    string `json:"language"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

type User struct {
	ID       string `json:"id"`
	Deleted  bool   `json:"deleted"`
	UserName string `json:"userName"`
	IsActive bool   `json:"isActive"`
}
