package blocklist

const EntityBlockList string = "BlockList"

// ResponseBlockList .
type ResponseBlockList struct {
	Total     int               `json:"total,omitempty"`
	BlockList []DomainBlockList `json:"list,omitempty"`
}

// DomainBlockList
type DomainBlockList struct {
	ID                      string     `json:"id,omitempty"`
	Name                    string     `json:"name,omitempty"`
	Deleted                 bool       `json:"deleted,omitempty"`
	Description             string     `json:"description,omitempty"`
	CreatedAt               string     `json:"createdAt,omitempty"`
	ModifiedAt              string     `json:"modifiedAt,omitempty"`
	Context                 string     `json:"context,omitempty"`
	CorrelationID           string     `json:"correlationId,omitempty"`
	EntityID                string     `json:"entityId,omitempty"`
	CompanyKey              string     `json:"companyKey,omitempty"`
	IDEmpotencyKey          string     `json:"idEmpotencyKey,omitempty"`
	Status                  string     `json:"status,omitempty"`
	DataReason              string     `json:"dataReason,omitempty"`
	DataStatus              string     `json:"dataStatus,omitempty"`
	DataBranch              string     `json:"dataBranch,omitempty"`
	DataNumber              string     `json:"dataNumber,omitempty"`
	DataBankIspb            string     `json:"dataBankIspb,omitempty"`
	DataBankCode            string     `json:"dataBankCode,omitempty"`
	DataBankName            string     `json:"dataBankName,omitempty"`
	DataHolderDocumentValue string     `json:"dataHolderDocumentValue,omitempty"`
	DataHolderDocumentType  string     `json:"dataHolderDocumentType,omitempty"`
	DataHolderName          string     `json:"dataHolderName,omitempty"`
	DataHolderStatus        string     `json:"dataHolderStatus,omitempty"`
	DataHolderType          string     `json:"dataHolderType,omitempty"`
	CreatedByID             string     `json:"createdById,omitempty"`
	CreatedByName           string     `json:"createdByName,omitempty"`
	ModifiedByID            string     `json:"modifiedById,omitempty"`
	ModifiedByName          string     `json:"modifiedByName,omitempty"`
	AssignedUserID          string     `json:"assignedUserId,omitempty"`
	AssignedUserName        string     `json:"assignedUserName,omitempty"`
	BankAccountID           string     `json:"bankAccountId,omitempty"`
	BankAccountName         string     `json:"bankAccountName,omitempty"`
}
