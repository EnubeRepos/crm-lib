package pixkey

const EntityPixKey string = "PixKey"

// ResponsePix .
type ResponsePixKey struct {
	Total  int            `json:"total,omitempty"`
	PixKey []DomainPixKey `json:"list,omitempty"`
}

type DomainPixKey struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Deleted              bool     `json:"deleted,omitempty"`
	Description          string   `json:"description,omitempty"`
	CreatedAt            string   `json:"createdAt,omitempty"`
	ModifiedAt           string   `json:"modifiedAt,omitempty"`
	Type                 string   `json:"type,omitempty"`
	Default              bool     `json:"default,omitempty"`
	CancelPixKey         bool     `json:"cancelPixKey,omitempty"`
	CreatedByID          string   `json:"createdById,omitempty"`
	CreatedByName        string   `json:"createdByName,omitempty"`
	ModifiedByID         string   `json:"modifiedById,omitempty"`
	ModifiedByName       string   `json:"modifiedByName,omitempty"`
	AssignedUserID       string   `json:"assignedUserId,omitempty"`
	AssignedUserName     string   `json:"assignedUserName,omitempty"`
	TeamsIds             []string `json:"teamsIds,omitempty"`
	PixsIds              []string `json:"pixsIds,omitempty"`
	BankAccountID        string   `json:"bankAccountId,omitempty"`
	BankAccountName      string   `json:"bankAccountName,omitempty"`
	PersonalContactsID   string   `json:"personalContactsId,omitempty"`
	PersonalContactsName string   `json:"personalContactsName,omitempty"`
	IsFollowed           bool     `json:"isFollowed,omitempty"`
	FollowersIds         []string `json:"followersIds,omitempty"`
	Status               string   `json:"status,omitempty"`
	StatusDetail         string   `json:"statusDetail,omitempty"`
}

type DomainPixKeyPutStatus struct {
	ID           string `json:"id,omitempty"`
	Status       string `json:"status,omitempty"`
	StatusDetail string `json:"statusDetail,omitempty"`
	Name         string `json:"name,omitempty"`
	CancelPixKey bool   `json:"cancelPixKey,omitempty"`
}
