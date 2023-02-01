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
	ID                               string   `json:"id,omitempty"`
	Name                             string   `json:"name,omitempty"`
	Deleted                          bool     `json:"deleted,omitempty"`
	Status                           string   `json:"status,omitempty"`
	DateStart                        string   `json:"dateStart,omitempty"`
	DateEnd                          string   `json:"dateEnd,omitempty"`
	IsAllDay                         bool     `json:"isAllDay,omitempty"`
	Duration                         int      `json:"duration,omitempty"`
	Description                      string   `json:"description,omitempty"`
	CreatedAt                        string   `json:"createdAt,omitempty"`
	ModifiedAt                       string   `json:"modifiedAt,omitempty"`
	Category                         string   `json:"category,omitempty"`
	TransactionType                  string   `json:"transactionType,omitempty"`
	Value                            float64  `json:"value,omitempty"`
	AccountTransaction               string   `json:"accountTransaction,omitempty"`
	Amount                           float64  `json:"amount,omitempty"`
	ConfirmTransaction               bool     `json:"confirmTransaction,omitempty"`
	ScheduleTransaction              string   `json:"scheduleTransaction,omitempty"`
	ExternalOrigin                   string   `json:"externalOrigin,omitempty"`
	EntityID                         string   `json:"entityId,omitempty"`
	IdEmpotencyKey                   string   `json:"idEmpotencyKey,omitempty"`
	CompanyKey                       string   `json:"companyKey,omitempty"`
	Context                          string   `json:"context,omitempty"`
	Timestamp                        string   `json:"timestamp,omitempty"`
	CorrelationID                    string   `json:"correlationId,omitempty"`
	Version                          string   `json:"version,omitempty"`
	AuthenticationCode               string   `json:"authenticationCode,omitempty"`
	RecipientDocumentValue           string   `json:"recipientDocumentValue,omitempty"`
	RecipientDocumentType            string   `json:"recipientDocumentType,omitempty"`
	RecipientType                    string   `json:"recipientType,omitempty"`
	RecipientName                    string   `json:"recipientName,omitempty"`
	RecipientAccountBranch           string   `json:"recipientAccountBranch,omitempty"`
	RecipientAccountNumber           string   `json:"recipientAccountNumber,omitempty"`
	RecipientAccountBankIspb         string   `json:"recipientAccountBankIspb,omitempty"`
	RecipientAccountBankCode         string   `json:"recipientAccountBankCode,omitempty"`
	RecipientAccountBankName         string   `json:"recipientAccountBankName,omitempty"`
	RecipientAccountBalanceValue     float64  `json:"recipientAccountBalanceValue,omitempty"`
	RecipientAccountBalanceCurrrency string   `json:"recipientAccountBalanceCurrrency,omitempty"`
	RecipientStatus                  string   `json:"recipientStatus,omitempty"`
	ChannelName                      string   `json:"channelName,omitempty"`
	ChannelControlNumber             string   `json:"channelControlNumber,omitempty"`
	ChannelControlNumberOriginal     string   `json:"channelControlNumberOriginal,omitempty"`
	ChannelOwnerNumber               string   `json:"channelOwnerNumber,omitempty"`
	SenderDocumentValue              string   `json:"senderDocumentValue,omitempty"`
	SenderDocumentType               string   `json:"senderDocumentType,omitempty"`
	SenderType                       string   `json:"senderType,omitempty"`
	SenderName                       string   `json:"senderName,omitempty"`
	SenderAccountBranch              string   `json:"senderAccountBranch,omitempty"`
	SenderAccountNumber              string   `json:"senderAccountNumber,omitempty"`
	SenderAccountBankIspb            string   `json:"senderAccountBankIspb,omitempty"`
	SenderAccountBankCode            string   `json:"senderAccountBankCode,omitempty"`
	SenderAccountBankName            string   `json:"senderAccountBankName,omitempty"`
	SenderStatus                     string   `json:"senderStatus,omitempty"`
	DateStartDate                    string   `json:"dateStartDate,omitempty"`
	DateEndDate                      string   `json:"dateEndDate,omitempty"`
	ValueCurrency                    string   `json:"valueCurrency,omitempty"`
	ParentType                       string   `json:"parentType,omitempty"`
	ParentName                       string   `json:"parentName,omitempty"`
	Commission                       bool     `json:"commission,omitempty"`
	TransactionCategory              string   `json:"transactionCategory,omitempty"`
	ParentID                         string   `json:"parentId,omitempty"`
	AssignedUserID                   string   `json:"assignedUserId,omitempty"`
	AssignedUserName                 string   `json:"assignedUserName,omitempty"`
	TeamsIds                         []string `json:"teamsIds,omitempty"`
	StatusTrackingID                 string   `json:"statusTrackingId,omitempty"`
	StatusTrackingName               string   `json:"statusTrackingName,omitempty"`
	//ValueConverted                    string   `json:"valueConverted,omitempty"`
	SalesOrderID                      string `json:"salesOrderId,omitempty"`
	SalesOrderName                    string `json:"salesOrderName,omitempty"`
	BankAccountID                     string `json:"bankAccountId,omitempty"`
	BankAccountName                   string `json:"bankAccountName,omitempty"`
	CommissionsID                     string `json:"commissionsId,omitempty"`
	CommissionsName                   string `json:"commissionsName,omitempty"`
	DocumentID                        string `json:"documentId,omitempty"`
	DocumentName                      string `json:"documentName,omitempty"`
	ParcelID                          string `json:"parcelId,omitempty"`
	ParcelName                        string `json:"parcelName,omitempty"`
	ContactID                         string `json:"contactId,omitempty"`
	ContactName                       string `json:"contactName,omitempty"`
	TedID                             string `json:"tedId,omitempty"`
	BillPaymentId                     string `json:"billPaymentId,omitempty"`
	TedName                           string `json:"tedName,omitempty"`
	UserID                            string `json:"userId,omitempty"`
	UserName                          string `json:"userName,omitempty"`
	PixID                             string `json:"pixId,omitempty"`
	InvoiceIssuanceId                 string `json:"invoiceIssuanceId,omitempty"`
	PixName                           string `json:"pixName,omitempty"`
	PixKeyID                          string `json:"pixKeyId,omitempty"`
	PixKeyName                        string `json:"pixKeyName,omitempty"`
	InstitutionFinancialRecipientID   string `json:"institutionFinancialRecipientId,omitempty"`
	InstitutionFinancialRecipientName string `json:"institutionFinancialRecipientName,omitempty"`
	InstitutionFinancialSenderID      string `json:"institutionFinancialSenderId,omitempty"`
	InstitutionFinancialSenderName    string `json:"institutionFinancialSenderName,omitempty"`
	ChargebackId                      string `json:"chargebackId,omitempty"`
	CardAlias                         string `json:"cardAlias,omitempty"`
	CardFunction                      string `json:"cardFunction,omitempty"`
	CardProxy                         string `json:"cardProxy,omitempty"`
	CardType                          string `json:"cardType,omitempty"`
	AuthorizationTransactionId        string `json:"authorizationTransactionId,omitempty"`
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

	StatusDone StatusTransaction = "Done"

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

	ChargebackId string `json:"chargebackId,omitempty"`

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

	ChargebackId string `json:"chargebackId,omitempty"`

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
