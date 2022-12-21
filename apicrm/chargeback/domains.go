package chargeback

const EntityChargeback string = "Chargeback"

// ResponseChargeback .
type ResponseChargeback struct {
	Total      int                `json:"total,omitempty"`
	Chargeback []DomainChargeback `json:"list,omitempty"`
}

type DomainChargeback struct {
	ID                 string      `json:"id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Deleted            bool        `json:"deleted,omitempty"`
	Description        string      `json:"description,omitempty"`
	CreatedAt          string      `json:"createdAt,omitempty"`
	ModifiedAt         string      `json:"modifiedAt,omitempty"`
	Amount             float64     `json:"amount,omitempty"`
	SetValue           string      `json:"setValue,omitempty"`
	PartialAmount      float64     `json:"partialAmount,omitempty"`
	FullAmount         int         `json:"fullAmount,omitempty"`
	RecipientSicCode   string      `json:"recipientSicCode,omitempty"`
	Status             string      `json:"status,omitempty"`
	CreatedByID        string      `json:"createdById,omitempty"`
	CreatedByName      string      `json:"createdByName,omitempty"`
	AssignedUserID     string      `json:"assignedUserId,omitempty"`
	PixID              string      `json:"pixId,omitempty"`
	TedID              interface{} `json:"tedId,omitempty"`
	RecipientID        string      `json:"recipientId,omitempty"`
	RecipientName      string      `json:"recipientName,omitempty"`
	StatusDetail       string      `json:"statusDetail,omitempty"`
	AuthenticationCode string      `json:"authenticationCode,omitempty"`
}

type DomainChargebackPutStatus struct {
	ID                 string `json:"id,omitempty"`
	Status             string `json:"status,omitempty"`
	StatusDetail       string `json:"statusDetail,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
}
