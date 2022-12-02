package commissions

const EntityCommissions string = "Commissions"

// ResponseCommissions .
type ResponseCommissions struct {
	Total       int                 `json:"total,omitempty"`
	Commissions []DomainCommissions `json:"list,omitempty"`
}

// Account
type DomainCommissions struct {
	ID                         string   `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Sales                      string   `json:"sales,omitempty"`
	Deleted                    bool     `json:"deleted,omitempty"`
	Description                string   `json:"description,omitempty"`
	CreatedAt                  string   `json:"createdAt,omitempty"`
	ModifiedAt                 string   `json:"modifiedAt,omitempty"`
	Status                     string   `json:"status,omitempty"`
	Amount                     float64  `json:"amount,omitempty"`
	ProductName                string   `json:"productName,omitempty"`
	SalesDate                  string   `json:"salesDate,omitempty"`
	Release                    bool     `json:"release,omitempty"`
	ParcelCode                 string   `json:"parcelCode,omitempty"`
	ParcelCreatedAt            string   `json:"parcelCreatedAt,omitempty"`
	ProductUnity               string   `json:"productUnity,omitempty"`
	ProductDivision            string   `json:"productDivision,omitempty"`
	ProductAddress             string   `json:"productAddress,omitempty"`
	Payment                    string   `json:"payment,omitempty"`
	DueDate                    string   `json:"dueDate,omitempty"`
	NetAmount                  float64  `json:"netAmount,omitempty"`
	InvoiceFee                 float64  `json:"invoiceFee,omitempty"`
	AdministrationFee          float64  `json:"administrationFee,omitempty"`
	ParcelNumber               int      `json:"parcelNumber,omitempty"`
	ExemptAccount              bool     `json:"exemptAccount,omitempty"`
	IntegratorFee              float64  `json:"integratorFee,omitempty"`
	AmountCurrency             string   `json:"amountCurrency,omitempty"`
	NetAmountCurrency          string   `json:"netAmountCurrency,omitempty"`
	InvoiceFeeCurrency         string   `json:"invoiceFeeCurrency,omitempty"`
	AdministrationFeeCurrency  string   `json:"administrationFeeCurrency,omitempty"`
	IntegratorFeeCurrency      string   `json:"integratorFeeCurrency,omitempty"`
	CreatedByID                string   `json:"createdById,omitempty"`
	CreatedByName              string   `json:"createdByName,omitempty"`
	ModifiedByID               string   `json:"modifiedById,omitempty"`
	ModifiedByName             string   `json:"modifiedByName,omitempty"`
	AssignedUserID             string   `json:"assignedUserId,omitempty"`
	AssignedUserName           string   `json:"assignedUserName,omitempty"`
	TeamsIds                   []string `json:"teamsIds,omitempty"`
	FinTransactionsIds         []string `json:"finTransactionsIds,omitempty"`
	AmountConverted            float64  `json:"amountConverted,omitempty"`
	SalesID                    string   `json:"salesId,omitempty"`
	SalesName                  string   `json:"salesName,omitempty"`
	BillingContactsIds         []string `json:"billingContactsIds,omitempty"`
	ParcelsID                  string   `json:"parcelsId,omitempty"`
	ParcelsName                string   `json:"parcelsName,omitempty"`
	NetAmountConverted         float64  `json:"netAmountConverted,omitempty"`
	InvoiceFeeConverted        float64  `json:"invoiceFeeConverted,omitempty"`
	AdministrationFeeConverted float64  `json:"administrationFeeConverted,omitempty"`
	AccountID                  string   `json:"accountId,omitempty"`
	AccountName                string   `json:"accountName,omitempty"`
	FeeApplicationsIds         []string `json:"feeApplicationsIds,omitempty"`
	IntegratorFeeConverted     float64  `json:"integratorFeeConverted,omitempty"`
	IsFollowed                 bool     `json:"isFollowed,omitempty"`
	FollowersIds               []string `json:"followersIds,omitempty"`
	RecipientDocumentNumber    string   `json:"recipientDocumentNumber,omitempty"`
}

// estrutura usada para o put de registration, evitando erro de campos vazios
type DomainCommissionsBase struct {
	ID                      string  `json:"id,omitempty"`
	Sales                   string  `json:"sales,omitempty"`
	ProductName             string  `json:"productName,omitempty"`
	Name                    string  `json:"name,omitempty"`
	Deleted                 bool    `json:"deleted,omitempty"`
	Description             string  `json:"description,omitempty"`
	CreatedAt               string  `json:"createdAt,omitempty"`
	ModifiedAt              string  `json:"modifiedAt,omitempty"`
	Status                  string  `json:"status,omitempty"`
	Amount                  float64 `json:"amount,omitempty"`
	SalesDate               string  `json:"salesDate,omitempty"`
	ParcelNumber            int     `json:"parcelNumber,omitempty"`
	ParcelsID               string  `json:"parcelsId,omitempty"`
	ParcelsName             string  `json:"parcelsName,omitempty"`
	AmountCurrency          string  `json:"amountCurrency,omitempty"`
	SalesID                 string  `json:"salesId,omitempty"`
	IsFollowed              bool    `json:"isFollowed,omitempty"`
	RecipientDocumentNumber string  `json:"recipientDocumentNumber,omitempty"`
}

type DomainComissionsPutStatus struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

type DomainReleaseCommission struct {
	ID      string `json:"id,omitempty"`
	Status  string `json:"status,omitempty"`
	Release bool   `json:"release,omitempty"`
}
