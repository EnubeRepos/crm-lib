package operationfeeapplication

const EntityOperationFeeApplication string = "OperationFeeApplication"

// ResponseOperationFeeApplication .
type ResponseOperationFeeApplication struct {
	Total                   int                             `json:"total,omitempty"`
	OperationFeeApplication []DomainOperationFeeApplication `json:"list,omitempty"`
}

// Request OperationFeeApplication
type DomainOperationFeeApplication struct {
	ID                          string  `json:"id,omitempty"`
	Name                        string  `json:"name,omitempty"`
	Deleted                     bool    `json:"deleted,omitempty"`
	Description                 string  `json:"description,omitempty"`
	CreatedAt                   string  `json:"createdAt,omitempty"`
	ModifiedAt                  string  `json:"modifiedAt,omitempty"`
	OperationFeeAmount          float64 `json:"operationFeeAmount,omitempty"`
	OperationFeeAmountCurrency  string  `json:"operationFeeAmountCurrency,omitempty"`
	CreatedByID                 string  `json:"createdById,omitempty"`
	AssignedUserID              string  `json:"assignedUserId,omitempty"`
	AssignedUserName            string  `json:"assignedUserName,omitempty"`
	OperationFeeAmountConverted float64 `json:"operationFeeAmountConverted,omitempty"`
}
