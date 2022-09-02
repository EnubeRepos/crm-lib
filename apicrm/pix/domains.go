package pix

const EntityPix string = "Pix"

// ResponsePix .
type ResponsePix struct {
	Total int         `json:"total,omitempty"`
	Pix   []DomainPix `json:"list,omitempty"`
}

// Account
type DomainPix struct {
	ID                         string   `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Deleted                    bool     `json:"deleted,omitempty"`
	Description                string   `json:"description,omitempty"`
	CreatedAt                  string   `json:"createdAt,omitempty"`
	ModifiedAt                 string   `json:"modifiedAt,omitempty"`
	PixType                    string   `json:"pixType,omitempty"`
	Amount                     float64  `json:"amount,omitempty"`
	ScheduleTransaction        string   `json:"scheduleTransaction,omitempty"`
	DateStart                  string   `json:"dateStart,omitempty"`
	OperationFeeAmount         float64  `json:"operationFeeAmount,omitempty"`
	PixKeyTransfer             string   `json:"pixKeyTransfer,omitempty"`
	StatusDetail               string   `json:"statusDetail,omitempty"`
	Status                     string   `json:"status,omitempty"`
	Cancel                     bool     `json:"cancel,omitempty"`
	ScheduleDate               string   `json:"scheduleDate,omitempty"`
	AuthenticationCode         string   `json:"authenticationCode,omitempty"`
	AmountCurrency             string   `json:"amountCurrency,omitempty"`
	OperationFeeAmountCurrency string   `json:"operationFeeAmountCurrency,omitempty"`
	CreatedByID                string   `json:"createdById,omitempty"`
	CreatedByName              string   `json:"createdByName,omitempty"`
	ModifiedByID               string   `json:"modifiedById,omitempty"`
	ModifiedByName             string   `json:"modifiedByName,omitempty"`
	AssignedUserID             string   `json:"assignedUserId,omitempty"`
	AssignedUserName           string   `json:"assignedUserName,omitempty"`
	TeamsIds                   []string `json:"teamsIds,omitempty"`
	//TeamsNames                  TeamsNames           `json:"teamsNames,omitempty"`
	FinTransactionsIds []string `json:"finTransactionsIds,omitempty"`
	//FinTransactionsNames        FinTransactionsNames `json:"finTransactionsNames,omitempty"`
	InstitutionFinancialID      string `json:"institutionFinancialId,omitempty"`
	InstitutionFinancialName    string `json:"institutionFinancialName,omitempty"`
	OperationFeeApplicationID   string `json:"operationFeeApplicationId,omitempty"`
	OperationFeeApplicationName string `json:"operationFeeApplicationName,omitempty"`
	HolderDocumentType          string `json:"holderDocumentType,omitempty"`
	HolderDocumentValue         string `json:"holderDocumentValue,omitempty"`
	HolderName                  string `json:"holderName,omitempty"`
	HolderType                  string `json:"holderType,omitempty"`
	ContactID                   string `json:"contactId,omitempty"`
	ContactName                 string `json:"contactName,omitempty"`
	//AmountConverted             float64                  `json:"amountConverted,omitempty"`
	//OperationFeeAmountConverted float64                  `json:"operationFeeAmountConverted,omitempty"`
	PixKeyID             string `json:"pixKeyId,omitempty"`
	PixKeyName           string `json:"pixKeyName,omitempty"`
	PersonalContactsID   string `json:"personalContactsId,omitempty"`
	PersonalContactsName string `json:"personalContactsName,omitempty"`
	//IsFollowed                  bool                 `json:"isFollowed,omitempty"`
	//FollowersIds                []string        `json:"followersIds,omitempty"`
	//FollowersNames              FollowersNames       `json:"followersNames,omitempty"`
}

type DomainPixPutStatus struct {
	ID                 string `json:"id,omitempty"`
	Status             string `json:"status,omitempty"`
	StatusDetail       string `json:"statusDetail,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
}

type DomainPixPutPixKeyInfo struct {
	ID                  string `json:"id,omitempty"`
	HolderDocumentType  string `json:"holderDocumentType,omitempty"`
	HolderDocumentValue string `json:"holderDocumentValue,omitempty"`
	HolderName          string `json:"holderName,omitempty"`
	HolderType          string `json:"holderType,omitempty"`
}
