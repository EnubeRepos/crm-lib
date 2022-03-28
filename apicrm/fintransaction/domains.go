package fintransaction

const EntityFinTransaction string = "FinTransaction"

// ResponseFinTransaction .
type ResponseFinTransaction struct {
	Total          int                    `json:"total,omitempty"`
	FinTransaction []DomainFinTransaction `json:"list,omitempty"`
}

// FinTransaction
type DomainFinTransaction struct {
	ID                                string   `json:"id"`
	Name                              string   `json:"name"`
	Deleted                           bool     `json:"deleted"`
	Status                            string   `json:"status"`
	DateStart                         string   `json:"dateStart"`
	DateEnd                           string   `json:"dateEnd"`
	IsAllDay                          bool     `json:"isAllDay"`
	Duration                          int      `json:"duration"`
	Description                       string   `json:"description"`
	CreatedAt                         string   `json:"createdAt"`
	ModifiedAt                        string   `json:"modifiedAt"`
	TransactionType                   string   `json:"transactionType"`
	Value                             string   `json:"value"`
	AccountTransaction                string   `json:"accountTransaction"`
	Amount                            float64  `json:"amount"`
	ConfirmTransaction                bool     `json:"confirmTransaction"`
	ScheduleTransaction               string   `json:"scheduleTransaction"`
	ExternalOrigin                    string   `json:"externalOrigin"`
	EntityID                          string   `json:"entityId"`
	IdEmpotencyKey                    string   `json:"idEmpotencyKey"`
	CompanyKey                        string   `json:"companyKey"`
	Context                           string   `json:"context"`
	Timestamp                         string   `json:"timestamp"`
	CorrelationID                     string   `json:"correlationId"`
	Version                           string   `json:"version"`
	AuthenticationCode                string   `json:"authenticationCode"`
	RecipientDocumentValue            string   `json:"recipientDocumentValue"`
	RecipientDocumentType             string   `json:"recipientDocumentType"`
	RecipientType                     string   `json:"recipientType"`
	RecipientName                     string   `json:"recipientName"`
	RecipientAccountBranch            string   `json:"recipientAccountBranch"`
	RecipientAccountNumber            string   `json:"recipientAccountNumber"`
	RecipientAccountBankIspb          string   `json:"recipientAccountBankIspb"`
	RecipientAccountBankCode          string   `json:"recipientAccountBankCode"`
	RecipientAccountBankName          string   `json:"recipientAccountBankName"`
	RecipientAccountBalanceValue      string   `json:"recipientAccountBalanceValue"`
	RecipientAccountBalanceCurrrency  string   `json:"recipientAccountBalanceCurrrency"`
	RecipientStatus                   string   `json:"recipientStatus"`
	ChannelName                       string   `json:"channelName"`
	ChannelControlNumber              string   `json:"channelControlNumber"`
	ChannelControlNumberOriginal      string   `json:"channelControlNumberOriginal"`
	ChannelOwnerNumber                string   `json:"channelOwnerNumber"`
	SenderDocumentValue               string   `json:"senderDocumentValue"`
	SenderDocumentType                string   `json:"senderDocumentType"`
	SenderType                        string   `json:"senderType"`
	SenderName                        string   `json:"senderName"`
	SenderAccountBranch               string   `json:"senderAccountBranch"`
	SenderAccountNumber               string   `json:"senderAccountNumber"`
	SenderAccountBankIspb             string   `json:"senderAccountBankIspb"`
	SenderAccountBankCode             string   `json:"senderAccountBankCode"`
	SenderAccountBankName             string   `json:"senderAccountBankName"`
	SenderStatus                      string   `json:"senderStatus"`
	DateStartDate                     string   `json:"dateStartDate"`
	DateEndDate                       string   `json:"dateEndDate"`
	ValueCurrency                     string   `json:"valueCurrency"`
	ParentType                        string   `json:"parentType"`
	ParentName                        string   `json:"parentName"`
	Commission                        bool     `json:"commission"`
	TransactionCategory               string   `json:"transactionCategory"`
	ParentID                          string   `json:"parentId"`
	AssignedUserID                    string   `json:"assignedUserId"`
	AssignedUserName                  string   `json:"assignedUserName"`
	TeamsIds                          []string `json:"teamsIds"`
	StatusTrackingID                  string   `json:"statusTrackingId"`
	StatusTrackingName                string   `json:"statusTrackingName"`
	ValueConverted                    string   `json:"valueConverted"`
	SalesOrderID                      string   `json:"salesOrderId"`
	SalesOrderName                    string   `json:"salesOrderName"`
	BankAccountID                     string   `json:"bankAccountId"`
	BankAccountName                   string   `json:"bankAccountName"`
	CommissionsID                     string   `json:"commissionsId"`
	CommissionsName                   string   `json:"commissionsName"`
	DocumentID                        string   `json:"documentId"`
	DocumentName                      string   `json:"documentName"`
	ParcelID                          string   `json:"parcelId"`
	ParcelName                        string   `json:"parcelName"`
	ContactID                         string   `json:"contactId"`
	ContactName                       string   `json:"contactName"`
	TedID                             string   `json:"tedId"`
	TedName                           string   `json:"tedName"`
	UserID                            string   `json:"userId"`
	UserName                          string   `json:"userName"`
	PixID                             string   `json:"pixId"`
	PixName                           string   `json:"pixName"`
	PixKeyID                          string   `json:"pixKeyId"`
	PixKeyName                        string   `json:"pixKeyName"`
	InstitutionFinancialRecipientID   string   `json:"institutionFinancialRecipientId"`
	InstitutionFinancialRecipientName string   `json:"institutionFinancialRecipientName"`
	InstitutionFinancialSenderID      string   `json:"institutionFinancialSenderId"`
	InstitutionFinancialSenderName    string   `json:"institutionFinancialSenderName"`
}
