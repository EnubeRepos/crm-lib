package invoiceissuance

const EntityInvoiceIssuance string = "InvoiceIssuance"

// ResponseInvoiceIssuance .
type ResponseInvoiceIssuance struct {
	Total           int                     `json:"total,omitempty"`
	InvoiceIssuance []DomainInvoiceIssuance `json:"list,omitempty"`
}

// Account
type DomainInvoiceIssuance struct {
	ID                         string   `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Deleted                    bool     `json:"deleted,omitempty"`
	Description                string   `json:"description,omitempty"`
	CreatedAt                  string   `json:"createdAt,omitempty"`
	ModifiedAt                 string   `json:"modifiedAt,omitempty"`
	DueDate                    string   `json:"dueDate,omitempty"`
	Amount                     float64  `json:"amount,omitempty"`
	Type                       string   `json:"type,omitempty"`
	DocumentNumber             string   `json:"documentNumber,omitempty"`
	OperationFeeAmount         float64  `json:"operationFeeAmount,omitempty"`
	PayerName                  string   `json:"payerName,omitempty"`
	PayerSicCode               string   `json:"payerSicCode,omitempty"`
	Status                     string   `json:"status,omitempty"`
	StatusDetail               string   `json:"statusDetail,omitempty"`
	AuthenticationCode         string   `json:"authenticationCode,omitempty"`
	ConfirmEmission            bool     `json:"confirmEmission,omitempty"`
	AmountCurrency             string   `json:"amountCurrency,omitempty"`
	OperationFeeAmountCurrency string   `json:"operationFeeAmountCurrency,omitempty"`
	CreatedByID                string   `json:"createdById,omitempty"`
	CreatedByName              string   `json:"createdByName,omitempty"`
	ModifiedByID               string   `json:"modifiedById,omitempty"`
	ModifiedByName             string   `json:"modifiedByName,omitempty"`
	AssignedUserID             string   `json:"assignedUserId,omitempty"`
	AssignedUserName           string   `json:"assignedUserName,omitempty"`
	TeamsIds                   []string `json:"teamsIds,omitempty"`
	InvoiceID                  string   `json:"invoiceId,omitempty"`
	InvoiceName                string   `json:"invoiceName,omitempty"`
	FinTransactionsIds         []string `json:"finTransactionsIds,omitempty"`
	InstitutionFinancialID     string   `json:"institutionFinancialId,omitempty"`
	InstitutionFinancialName   string   `json:"institutionFinancialName,omitempty"`

	PaymentDate string `json:"paymentDate,omitempty"`
}

type DomainInvoiceIssuancePutAuthCode struct {
	ID                 string `json:"id,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
	InvoiceId          string `json:"invoiceId,omitempty"`
	InvoiceNumber      string `json:"invoiceNumber,omitempty"`
	Status             string `json:"status,omitempty"`
	StatusDetail       string `json:"statusDetail,omitempty"`
}

type DomainInvoiceIssuancePutStatus struct {
	ID           string `json:"id,omitempty"`
	Status       string `json:"status,omitempty"`
	StatusDetail string `json:"statusDetail,omitempty"`
	PaymentDate  string `json:"paymentDate,omitempty"`
}
