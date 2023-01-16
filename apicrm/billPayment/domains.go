package billPayment

const EntityBillPayment string = "BillPayment"

// ResponseBillPayment .
type ResponseBillPayment struct {
	Total       int                 `json:"total,omitempty"`
	BillPayment []DomainBillPayment `json:"list,omitempty"`
}

type DomainBillPayment struct {
	ID                 string   `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	Deleted            bool     `json:"deleted,omitempty"`
	Description        string   `json:"description,omitempty"`
	CreatedAt          string   `json:"createdAt,omitempty"`
	ModifiedAt         string   `json:"modifiedAt,omitempty"`
	DocumentNumber     string   `json:"documentNumber,omitempty"`
	Amount             float64  `json:"amount,omitempty"`
	Status             string   `json:"status,omitempty"`
	AuthenticationCode string   `json:"authenticationCode,omitempty"`
	CreatedByID        string   `json:"createdById,omitempty"`
	CreatedByName      string   `json:"createdByName,omitempty"`
	ModifiedByID       string   `json:"modifiedById,omitempty"`
	ModifiedByName     string   `json:"modifiedByName,omitempty"`
	AssignedUserID     string   `json:"assignedUserId,omitempty"`
	AssignedUserName   string   `json:"assignedUserName,omitempty"`
	Assignor           string   `json:"assignor,omitempty"`
	TeamsIds           []string `json:"teamsIds,omitempty"`
	//TeamsNames         TeamsNames     `json:"teamsNames,omitempty"`
	IsFollowed   bool     `json:"isFollowed,omitempty"`
	FollowersIds []string `json:"followersIds,omitempty"`
	//FollowersNames     FollowersNames `json:"followersNames,omitempty"`
}

type DomainBillPaymentPutStatus struct {
	ID                 string `json:"id,omitempty"`
	Status             string `json:"status,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
}
