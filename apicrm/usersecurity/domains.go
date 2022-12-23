package usersecurity

const EntityUserSecurity string = "UserSecurity"

// ResponseUser .
type ResponseUser struct {
	Total        int                   `json:"total,omitempty"`
	UserSecurity []DomainUserStatusMFA `json:"list,omitempty"`
}

type DomainUserStatusMFA struct {
	Auth2FA       bool   `json:"auth2FA"`
	Auth2FAMethod string `json:"auth2FAMethod,omitempty"`
	Password      string `json:"password,omitempty"`
}
