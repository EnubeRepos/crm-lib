package fintransaction

import "time"

const EntityFinTransaction string = "FinTransaction"

// ResponseFinTransaction .
type ResponseFinTransaction struct {
	Total          int                    `json:"total,omitempty"`
	FinTransaction []DomainFinTransaction `json:"list,omitempty"`
}

type DomainFinTransactionBase struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Value         float32 `json:"value"`
	DateStartDate string  `json:"dateStartDate"`
	DateEndDate   string  `json:"dateEndDate"`
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
	Value                             float64  `json:"value"`
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
	RecipientAccountBalanceValue      float64  `json:"recipientAccountBalanceValue"`
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
	//ValueConverted                    string   `json:"valueConverted"`
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

type TransactionType string

const (
	TransactionTypeTED TransactionType = "TED"

	TransactionTypeDOC TransactionType = "DOC"

	TransactionTypePIX TransactionType = "PIX"

	TransactionTypeBoleto TransactionType = "Boleto"

	TransactionTypeTEDInternalBank TransactionType = "Triplic Bank" //Transferência entre contas 10+Bank
)

type AccountTransaction string

const (
	//Débito
	AccountTransactionSent AccountTransaction = "sent"

	//Crédito
	AccountTransactionReceived AccountTransaction = "received"
)

type StatusTransaction string

const (
	StatusWaiting StatusTransaction = "Waiting"

	StatusPlanned StatusTransaction = "Planned"

	StatusNotHeld StatusTransaction = "NotHeld"

	StatusHeld StatusTransaction = "Held"

	StatusCanceled StatusTransaction = "Canceled"
)

type RequestFinTrasactionTED struct {
	// nome genérico da transação
	Name string `json:"name"`

	//Waiting
	Status StatusTransaction `json:"status"`

	// Agendado para:
	DateStart time.Time `json:"dateStart"`

	// "Pagamento TED,DOC,PIX,BOLETO, TED INTERNAL"
	TransactionType TransactionType `json:"transactionType"`

	// valor convertido para negativo caso for débito -- tipo monetário
	Value float64 `json:"value"`

	AccountTransaction AccountTransaction `json:"accountTransaction"`

	// preenchimento apenas do valor, não podendo ser negativo
	Amount float64 `json:"amount"`

	ConfirmTransaction bool `json:"confirmTransaction"`

	// yes/no -- rever
	ScheduleTransaction string `json:"scheduleTransaction"`

	CorrelationID string `json:"correlationId"`

	AuthenticationCode string `json:"authenticationCode"`

	AssignedUserID string `json:"assignedUserId"`

	TeamsIds []string `json:"teamsIds"`

	// Emissor:
	BankAccountID string `json:"bankAccountId"`

	CommissionsID string `json:"commissionsId"`

	ParcelID string `json:"parcelId"`

	// Beneficiário externo do 10maisbank
	ContactID string `json:"contactId"`

	TedID string `json:"tedId"`

	// Beneficiário interno do 10+bank
	UserID string `json:"userId"`
}

type ResponseFinTransactionTED struct {
	// nome genérico da transação
	Id string `json:"id"`

	// nome genérico da transação
	Name string `json:"name"`

	//Waiting
	Status StatusTransaction `json:"status"`

	// Agendado para:
	DateStart time.Time `json:"dateStart"`

	// "Pagamento TED,DOC,PIX,BOLETO, TED INTERNAL"
	TransactionType TransactionType `json:"transactionType"`

	// valor convertido para negativo caso for débito -- tipo monetário
	Value float64 `json:"value"`

	AccountTransaction AccountTransaction `json:"accountTransaction"`

	// preenchimento apenas do valor, não podendo ser negativo
	Amount float64 `json:"amount"`

	ConfirmTransaction bool `json:"confirmTransaction"`

	// yes/no -- rever
	ScheduleTransaction string `json:"scheduleTransaction"`

	CorrelationID string `json:"correlationId"`

	AuthenticationCode string `json:"authenticationCode"`

	AssignedUserID string `json:"assignedUserId"`

	TeamsIds []string `json:"teamsIds"`

	// Emissor:
	BankAccountID string `json:"bankAccountId"`

	CommissionsID string `json:"commissionsId"`

	ParcelID string `json:"parcelId"`

	// Beneficiário externo do 10maisbank
	ContactID string `json:"contactId"`

	TedID string `json:"tedId"`

	// Beneficiário interno do 10+bank
	UserID string `json:"userId"`
}

type RequestFinTrasactionPIX struct {
	// nome genérico da transação
	Name string `json:"name"`

	//Waiting
	Status StatusTransaction `json:"status"`

	// Agendado para:
	DateStart time.Time `json:"dateStart"`

	// "Pagamento TED,DOC,PIX,BOLETO, TED INTERNAL"
	TransactionType TransactionType `json:"transactionType"`

	// valor convertido para negativo caso for débito -- tipo monetário
	Value float64 `json:"value"`

	AccountTransaction AccountTransaction `json:"accountTransaction"`

	// preenchimento apenas do valor, não podendo ser negativo
	Amount float64 `json:"amount"`

	ConfirmTransaction bool `json:"confirmTransaction"`

	// yes/no -- rever
	ScheduleTransaction string `json:"scheduleTransaction"`

	CorrelationID string `json:"correlationId"`

	AuthenticationCode string `json:"authenticationCode"`

	AssignedUserID string `json:"assignedUserId"`

	TeamsIds []string `json:"teamsIds"`

	// Emissor:
	BankAccountID string `json:"bankAccountId"`

	CommissionsID string `json:"commissionsId"`

	ParcelID string `json:"parcelId"`

	// Beneficiário externo do 10maisbank
	ContactID string `json:"contactId"`

	// Beneficiário interno do 10+bank
	UserID string `json:"userId"`

	PixID string `json:"pixId"`

	PixKeyID string `json:"pixKeyId"`
}

type RequestFinTrasactionTaxTariff struct {
	// nome genérico da transação
	Name string `json:"name"`

	//Waiting
	Status StatusTransaction `json:"status"`

	// Agendado para:
	DateStart time.Time `json:"dateStart"`

	// "Pagamento TED,DOC,PIX,BOLETO, TED INTERNAL"
	TransactionType TransactionType `json:"transactionType"`

	// valor convertido para negativo caso for débito -- tipo monetário
	Value float64 `json:"value"`

	AccountTransaction AccountTransaction `json:"accountTransaction"`

	// preenchimento apenas do valor, não podendo ser negativo
	Amount float64 `json:"amount"`

	ConfirmTransaction bool `json:"confirmTransaction"`

	CorrelationID string `json:"correlationId"`

	AuthenticationCode string `json:"authenticationCode"`

	AssignedUserID string `json:"assignedUserId"`

	TeamsIds []string `json:"teamsIds"`

	// Emissor:
	BankAccountID string `json:"bankAccountId"`

	CommissionsID string `json:"commissionsId"`

	ParcelID string `json:"parcelId"`

	// Beneficiário interno do 10+bank
	UserID string `json:"userId"`
}

type RequestFinTrasactionBoleto struct {
	// nome genérico da transação
	Name string `json:"name"`

	//Waiting
	Status StatusTransaction `json:"status"`

	// Agendado para:
	DateStart time.Time `json:"dateStart"`

	// "Pagamento TED,DOC,PIX,BOLETO, TED INTERNAL"
	TransactionType TransactionType `json:"transactionType"`

	// valor convertido para negativo caso for débito -- tipo monetário
	Value float64 `json:"value"`

	AccountTransaction AccountTransaction `json:"accountTransaction"`

	// preenchimento apenas do valor, não podendo ser negativo
	Amount float64 `json:"amount"`

	ConfirmTransaction bool `json:"confirmTransaction"`

	CorrelationID string `json:"correlationId"`

	AuthenticationCode string `json:"authenticationCode"`

	AssignedUserID string `json:"assignedUserId"`

	TeamsIds []string `json:"teamsIds"`

	// Emissor:
	BankAccountID string `json:"bankAccountId"`

	ParcelID string `json:"parcelId"`

	// Beneficiário interno do 10+bank
	UserID string `json:"userId"`
}
