package domains

// ResponseDistributor .
type ResponseDistributor struct {
	Total       int           `json:"total,omitempty"`
	Distributor []Distributor `json:"list,omitempty"`
}

type Distributor struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Deleted          bool   `json:"deleted,omitempty"`
	Description      string `json:"description,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	ModifiedAt       string `json:"modifiedAt,omitempty"`
	MpnID            string `json:"mpnId,omitempty"`
	AppID            string `json:"appId,omitempty"`
	AccountID        string `json:"accountId,omitempty"`
	Domain           string `json:"domain,omitempty"`
	ClientID         string `json:"clientId,omitempty"`
	Culture          string `json:"culture,omitempty"`
	ClientSecret     string `json:"clientSecret,omitempty"`
	RefreshToken     string `json:"refreshToken,omitempty"`
	ExpireToken      string `json:"expireToken,omitempty"`
	AccessToken      string `json:"accessToken,omitempty"`
	CreatedByID      string `json:"createdById,omitempty"`
	CreatedByName    string `json:"createdByName,omitempty"`
	ModifiedByID     string `json:"modifiedById,omitempty"`
	ModifiedByName   string `json:"modifiedByName,omitempty"`
	AssignedUserID   string `json:"assignedUserId,omitempty"`
	AssignedUserName string `json:"assignedUserName,omitempty"`
}
