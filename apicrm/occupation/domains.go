package occupation

const EntityOccupation string = "Occupation"

// ResponseOccupation .
type ResponseOccupation struct {
	Total      int                `json:"total,omitempty"`
	Occupation []DomainOccupation `json:"list,omitempty"`
}

// estrutura usada para o put de registration, evitando erro de campos vazios
type DomainOccupationBase struct {
	ID         string `json:"id,omitempty"`
	StatusType string `json:"statusType,omitempty"`
	Code       int    `json:"code,omitempty"`
}

// Account
type DomainOccupation struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Code             string `json:"code,omitempty"`
	AssignedUser     string `json:"assignedUser,omitempty"`
	AssignedUserID   string `json:"assignedUserId,omitempty"`
	AssignedUserName string `json:"assignedUserName,omitempty"`
}
