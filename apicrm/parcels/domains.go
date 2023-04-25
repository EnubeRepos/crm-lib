package parcels

const EntityParcels string = "Parcels"

// ResponseParcels .
type ResponseParcels struct {
	Total   int             `json:"total,omitempty"`
	Parcels []DomainParcels `json:"list,omitempty"`
}

// Parcels
type DomainParcels struct {
	ID                          string   `json:"id,omitempty"`
	Name                        string   `json:"name,omitempty"`
	Description                 string   `json:"description,omitempty"`
	ProductName                 string   `json:"productName,omitempty"`
	ProductUnity                string   `json:"productUnity,omitempty"`
	SalesDate                   string   `json:"salesDate,omitempty"`
	SaleStatus                  string   `json:"saleStatus,omitempty"`
	ParcelNumber                int      `json:"parcelNumber,omitempty"`
	ParcelCode                  string   `json:"parcelCode,omitempty"`
	ProductDivision             string   `json:"productDivision,omitempty"`
	Status                      string   `json:"status,omitempty"`
	DueDate                     string   `json:"dueDate,omitempty"`
	ClosePayment                string   `json:"closePayment,omitempty"`
	Amount                      float64  `json:"amount,omitempty"`
	Payment                     string   `json:"payment,omitempty"`
	PaymentDate                 string   `json:"paymentDate,omitempty"`
	ParcelPaid                  string   `json:"parcelPaid,omitempty"`
	ReturnFee                   float64  `json:"returnFee,omitempty"`
	ApproveReturnTranfer        bool     `json:"approveReturnTranfer,omitempty"`
	InvoiceNumber               string   `json:"invoiceNumber,omitempty"`
	AuthenticationCode          string   `json:"authenticationCode,omitempty"`
	AmountCurrency              string   `json:"amountCurrency,omitempty"`
	ReturnFeeCurrency           string   `json:"returnFeeCurrency,omitempty"`
	SalesId                     string   `json:"salesId,omitempty"`
	SalesName                   string   `json:"salesName,omitempty"`
	BillingContactsIds          []string `json:"billingContactsIds,omitempty"`
	FinTransactionsIds          []string `json:"finTransactionsIds,omitempty"`
	InvoiceId                   string   `json:"invoiceId,omitempty"`
	InvoiceName                 string   `json:"invoiceName,omitempty"`
	AmountConverted             float64  `json:"amountConverted,omitempty"`
	CommissionssIds             []string `json:"commissionssIds,omitempty"`
	ReturnFeeConverted          float64  `json:"returnFeeConverted,omitempty"`
	BillingContactId            string   `json:"billingContactId,omitempty"`
	BillingContactName          string   `json:"billingContactName,omitempty"`
	ParentId                    string   `json:"parentId,omitempty"`
	ParentType                  string   `json:"parentType,omitempty"`
	ParentName                  string   `json:"parentName,omitempty"`
	IsfFollowed                 bool     `json:"isFollowed,omitempty"`
	FollowersIds                []string `json:"followersIds,omitempty"`
	CreatedById                 string   `json:"createdById,omitempty"`
	AssignedUserId              string   `json:"assignedUserId,omitempty"`
	OurNumber                   string   `json:"ourNumber,omitempty"`
	Reprocess                   bool     `json:"reprocess,omitempty"`
	Deleted                     bool     `json:"deleted,omitempty"`
	RecipientDocumentNumber     string   `json:"recipientDocumentNumber,omitempty"`
	AmountCommission            float64  `json:"amountCommission,omitempty"`
	SicCode                     string   `json:"sicCode,omitempty"`
	SumCommissionsAmount        float64  `json:"sumCommissionsAmount,omitempty"`
	ChargebackRequest           bool     `json:"chargebackRequest,omitempty"`
	ConfirmChargeback           bool     `json:"confirmChargeback,omitempty"`
	TransactionType             string   `json:"transactionType,omitempty"`
	BillingContactAccountNumber string   `json:"billingContactAccountNumber,omitempty"`
	BillingContactAgency        string   `json:"billingContactAgency,omitempty"`
	BillingContactPixKey        string   `json:"billingContactPixKey,omitempty"`
	RegistrationRequestId       string   `json:"registrationRequestId,omitempty"`
	CommissionsAssignedUserId   string   `json:"commissionsAssignedUserId,omitempty"`
	AccountId                   string   `json:"accountId,omitempty"`
	AccountName                 string   `json:"accountName,omitempty"`
	InstitutionFinancialId      string   `json:"institutionFinancialId,omitempty"`
	InstitutionFinancialName    string   `json:"institutionFinancialName,omitempty"`
	ChargebackReceiptId         string   `json:"chargebackReceiptId,omitempty"`
	CanceledAt                  string   `json:"canceledAt,omitempty"`
}

type DomainParcelsPutAuthCode struct {
	ID                 string `json:"id,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
	InvoiceId          string `json:"invoiceId,omitempty"`
	InvoiceNumber      string `json:"invoiceNumber,omitempty"`
	Status             string `json:"status,omitempty"`
	StatusDetail       string `json:"statusDetail,omitempty"`
	ParcelPaid         string `json:"parcelPaid,omitempty"`
	Reprocess          bool   `json:"reprocess,omitempty"`
	ChargebackRequest  bool   `json:"chargebackRequest,omitempty"`
	ConfirmChargeback  bool   `json:"confirmChargeback,omitempty"`
	OurNumber          string `json:"ourNumber,omitempty"`
}

type DomainParcelsPutStatus struct {
	ID                string `json:"id,omitempty"`
	Status            string `json:"status,omitempty"`
	StatusDetail      string `json:"statusDetail,omitempty"`
	ParcelPaid        string `json:"parcelPaid,omitempty"`
	Reprocess         bool   `json:"reprocess,omitempty"`
	ChargebackRequest bool   `json:"chargebackRequest,omitempty"`
	ConfirmChargeback bool   `json:"confirmChargeback,omitempty"`
	PaymentDate       string `json:"paymentDate,omitempty"`
	CanceledAt        string `json:"canceledAt,omitempty"`
}
