package limitmanagement

const EntityLimitManagement string = "LimitManagement"

// ResponseLimitManagement .
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
	//TedInternalPerTransactionCurrency  string         `json:"tedInternalPerTransactionCurrency,omitempty"`
	//TedInternalPerDayCurrency          string    `json:"tedInternalPerDayCurrency,omitempty"`
	//TedInternalPerMonthCurrency        string    `json:"tedInternalPerMonthCurrency,omitempty"`
	//PixPerTransactionCurrency          string         `json:"pixPerTransactionCurrency,omitempty"`
	//PixPerDayCurrency                  string         `json:"pixPerDayCurrency,omitempty"`
	//CardPerTransactionCurrency         string         `json:"cardPerTransactionCurrency,omitempty"`
	//CardPerDayCurrency                 string         `json:"cardPerDayCurrency,omitempty"`
	//BillPaymentPerTransactionCurrency  string         `json:"billPaymentPerTransactionCurrency,omitempty"`
	//BillPaymentPerDayCurrency          string         `json:"billPaymentPerDayCurrency,omitempty"`
	//InvoicePerTransactionCurrency      string         `json:"invoicePerTransactionCurrency,omitempty"`
	//InvoicePerDayCurrency              string         `json:"invoicePerDayCurrency,omitempty"`
	//InvoicePerMonthCurrency            string         `json:"invoicePerMonthCurrency,omitempty"`
	//PixPerDayTotalCurrency             string         `json:"pixPerDayTotalCurrency,omitempty"`
	//CardPerDayTotalCurrency            string    `json:"cardPerDayTotalCurrency,omitempty"`
	//BillPaymentPerDayTotalCurrency     string    `json:"billPaymentPerDayTotalCurrency,omitempty"`
	//InvoicePerDayTotalCurrency         string         `json:"invoicePerDayTotalCurrency,omitempty"`
	//InvoicePerMonthTotalCurrency       string         `json:"invoicePerMonthTotalCurrency,omitempty"`
	//CreatedByID                        string         `json:"createdById,omitempty"`
	//CreatedByName                      string         `json:"createdByName,omitempty"`
	//ModifiedByID                       string         `json:"modifiedById,omitempty"`
	//ModifiedByName                     string         `json:"modifiedByName,omitempty"`
	AssignedUserID                     string         `json:"assignedUserId,omitempty"`
	AssignedUserName                   string         `json:"assignedUserName,omitempty"`
	//TeamsIds                           []string  `json:"teamsIds,omitempty"`
	//TeamsNames                         TeamsNames     `json:"teamsNames,omitempty"`
	//TedInternalPerTransactionConverted int            `json:"tedInternalPerTransactionConverted,omitempty"`
	//TedInternalPerDayConverted         string    `json:"tedInternalPerDayConverted,omitempty"`
	//TedInternalPerMonthConverted       string    `json:"tedInternalPerMonthConverted,omitempty"`
	//PixPerTransactionConverted         int            `json:"pixPerTransactionConverted,omitempty"`
	//PixPerDayConverted                 int            `json:"pixPerDayConverted,omitempty"`
	//CardPerTransactionConverted        int            `json:"cardPerTransactionConverted,omitempty"`
	//CardPerDayConverted                int            `json:"cardPerDayConverted,omitempty"`
	//BillPaymentPerTransactionConverted int            `json:"billPaymentPerTransactionConverted,omitempty"`
	//BillPaymentPerDayConverted         int            `json:"billPaymentPerDayConverted,omitempty"`
	//InvoicePerTransactionConverted     int            `json:"invoicePerTransactionConverted,omitempty"`
	//InvoicePerDayConverted             int            `json:"invoicePerDayConverted,omitempty"`
	//InvoicePerMonthConverted           int            `json:"invoicePerMonthConverted,omitempty"`
	BankAccountID                      string         `json:"bankAccountId,omitempty"`
	BankAccountName                    string         `json:"bankAccountName,omitempty"`
	//PixPerDayTotalConverted            int            `json:"pixPerDayTotalConverted,omitempty"`
	//CardPerDayTotalConverted           string    `json:"cardPerDayTotalConverted,omitempty"`
	//BillPaymentPerDayTotalConverted    string    `json:"billPaymentPerDayTotalConverted,omitempty"`
	//InvoicePerDayTotalConverted        int            `json:"invoicePerDayTotalConverted,omitempty"`
	//InvoicePerMonthTotalConverted      int            `json:"invoicePerMonthTotalConverted,omitempty"`
	//IsFollowed                         bool           `json:"isFollowed,omitempty"`
	//FollowersIds                       []string  `json:"followersIds,omitempty"`
	//FollowersNames                     FollowersNames `json:"followersNames,omitempty"`
}
