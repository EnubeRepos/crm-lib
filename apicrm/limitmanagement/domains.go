package limitmanagement

const EntityLimitManagement string = "LimitManagement"

type ResponseLimitManagement struct {
	Total int         `json:"total,omitempty"`
	LimitManagement   []DomainLimitManagement `json:"list,omitempty"`
}

type DomainLimitManagement struct {
	ID                                 string         `json:"id,omitempty"`
	Name                               string         `json:"name,omitempty"`
	Deleted                            bool           `json:"deleted,omitempty"`
	Description                        string    `json:"description,omitempty"`
	CreatedAt                          string         `json:"createdAt,omitempty"`
	ModifiedAt                         string         `json:"modifiedAt,omitempty"`
	TedInternalPerTransaction          float64            `json:"tedInternalPerTransaction,omitempty"`
	TedInternalPerDay                  string    `json:"tedInternalPerDay,omitempty"`
	TedInternalPerMonth                string    `json:"tedInternalPerMonth,omitempty"`
	PixPerTransaction                  float64            `json:"pixPerTransaction,omitempty"`
	PixPerDay                          float64            `json:"pixPerDay,omitempty"`
	WithdrawCard                       string         `json:"withdrawCard,omitempty"`
	CardPerTransaction                 float64            `json:"cardPerTransaction,omitempty"`
	CardPerDay                         float64            `json:"cardPerDay,omitempty"`
	BillPaymentPerTransaction          float64            `json:"billPaymentPerTransaction,omitempty"`
	BillPaymentPerDay                  float64            `json:"billPaymentPerDay,omitempty"`
	InvoicePerTransaction              float64            `json:"invoicePerTransaction,omitempty"`
	InvoicePerDay                      float64            `json:"invoicePerDay,omitempty"`
	InvoicePerMonth                    float64            `json:"invoicePerMonth,omitempty"`
	BankAccountCustomer                string         `json:"bankAccountCustomer,omitempty"`
	PixPerDayTotal                     float64            `json:"pixPerDayTotal,omitempty"`
	CardPerDayTotal                    string    `json:"cardPerDayTotal,omitempty"`
	BillPaymentPerDayTotal             string    `json:"billPaymentPerDayTotal,omitempty"`
	InvoicePerDayTotal                 float64            `json:"invoicePerDayTotal,omitempty"`
	InvoicePerMonthTotal               float64            `json:"invoicePerMonthTotal,omitempty"`
	AssignedUserID                     string         `json:"assignedUserId,omitempty"`
	AssignedUserName                   string         `json:"assignedUserName,omitempty"`
	BankAccountID                      string         `json:"bankAccountId,omitempty"`
	BankAccountName                    string         `json:"bankAccountName,omitempty"`
}
