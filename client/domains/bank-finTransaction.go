package domains

type ResponseFinTransaction struct {
	Total   int       `json:"total,omitempty"`
	FinTransaction []FinTransaction `json:"list,omitempty"`
}

type FinTransaction struct {
	Source              string   `json:"source,omitempty"`
	ID                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	Deleted             bool     `json:"deleted,omitempty"`
	Status              string   `json:"status,omitempty"`
	DateStart           string   `json:"dateStart,omitempty"`
	DateEnd             string   `json:"dateEnd,omitempty"`
	IsAllDay            bool     `json:"isAllDay,omitempty"`
	Duration            string   `json:"duration,omitempty"`
	Description         string   `json:"description,omitempty"`
	Reminders           []string `json:"reminders,omitempty"`
	CreatedAt           string   `json:"createdAt,omitempty"`
	ModifiedAt          string   `json:"modifiedAt,omitempty"`
	TransactionType     string   `json:"transactionType,omitempty"`
	Value               int      `json:"value,omitempty"`
	Commission          bool     `json:"commission,omitempty"`
	AccountTransaction  string   `json:"accountTransaction,omitempty"`
	TransactionCategory string   `json:"transactionCategory,omitempty"`
	Amount              int      `json:"amount,omitempty"`
	BankName            string   `json:"bankName,omitempty"`
	DateStartDate       string   `json:"dateStartDate,omitempty"`
	DateEndDate         string   `json:"dateEndDate,omitempty"`
	ValueCurrency       string   `json:"valueCurrency,omitempty"`
	ParentID            string   `json:"parentId,omitempty"`
	ParentType          string   `json:"parentType,omitempty"`
	ParentName          string   `json:"parentName,omitempty"`
	CreatedByID         string   `json:"createdById,omitempty"`
	CreatedByName       string   `json:"createdByName,omitempty"`
	ModifiedByID        string   `json:"modifiedById,omitempty"`
	ModifiedByName      string   `json:"modifiedByName,omitempty"`
	AssignedUserID      string   `json:"assignedUserId,omitempty"`
	AssignedUserName    string   `json:"assignedUserName,omitempty"`
	TeamsIds            []string `json:"teamsIds,omitempty"`
	SalesOrderItemsIds  []string `json:"salesOrderItemsIds,omitempty"`
	StatusTrackingID    string   `json:"statusTrackingId,omitempty"`
	StatusTrackingName  string   `json:"statusTrackingName,omitempty"`
	ValueConverted      int      `json:"valueConverted,omitempty"`
	SalesOrderID        string   `json:"salesOrderId,omitempty"`
	SalesOrderName      string   `json:"salesOrderName,omitempty"`
	BankAccountID       string   `json:"bankAccountId,omitempty"`
	BankAccountName     string   `json:"bankAccountName,omitempty"`
	SplitsIds           []string `json:"splitsIds,omitempty"`
	// SplitsNames          SplitsNames          `json:"splitsNames,omitempty"`
	CommissionsID   string `json:"commissionsId,omitempty"`
	CommissionsName string `json:"commissionsName,omitempty"`
	ContactID       string `json:"contactId,omitempty"`
	ContactName     string `json:"contactName,omitempty"`
	DocumentID      string `json:"documentId,omitempty"`
	DocumentName    string `json:"documentName,omitempty"`
}
