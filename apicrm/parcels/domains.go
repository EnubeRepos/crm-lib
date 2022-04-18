package parcels

const EntityParcels string = "Parcels"

// ResponseParcels .
type ResponseParcels struct {
	Total   int             `json:"total,omitempty"`
	Parcels []DomainParcels `json:"list,omitempty"`
}

// Account
type DomainParcels struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	ProductName          string   `json:"productName,omitempty"`
	ProductUnity         string   `json:"productUnity,omitempty"`
	SalesDate            string   `json:"salesDate,omitempty"`
	SaleStatus           string   `json:"saleStatus,omitempty"`
	ParcelNumber         int      `json:"parcelNumber,omitempty"`
	ParcelCode           string   `json:"parcelCode,omitempty"`
	ProductDivision      string   `json:"productDivision,omitempty"`
	Status               string   `json:"status,omitempty"`
	DueDate              string   `json:"dueDate,omitempty"`
	Amount               int      `json:"amount,omitempty"`
	Payment              string   `json:"payment,omitempty"`
	ParcelPaid           string   `json:"parcelPaid,omitempty"`
	ReturnFee            string   `json:"returnFee,omitempty"`
	ApproveReturnTranfer bool     `json:"approveReturnTranfer,omitempty"`
	InvoiceNumber        string   `json:"invoiceNumber,omitempty"`
	AuthenticationCode   string   `json:"authenticationCode,omitempty"`
	AmountCurrency       string   `json:"amountCurrency,omitempty"`
	ReturnFeeCurrency    string   `json:"returnFeeCurrency,omitempty"`
	SalesId              string   `json:"salesId,omitempty"`
	SalesName            string   `json:"salesName,omitempty"`
	BillingContactsIds   []string `json:"billingContactsIds,omitempty"`
	FinTransactionsIds   []string `json:"finTransactionsIds,omitempty"`
	InvoiceId            string   `json:"invoiceId,omitempty"`
	InvoiceName          string   `json:"invoiceName,omitempty"`
	AmountConverted      int      `json:"amountConverted,omitempty"`
	CommissionssIds      []string `json:"commissionssIds,omitempty"`
	ReturnFeeConverted   string   `json:"returnFeeConverted,omitempty"`
	BillingContactId     string   `json:"billingContactId,omitempty"`
	BillingContactName   string   `json:"billingContactName,omitempty"`
	ParentId             string   `json:"parentId,omitempty"`
	ParentType           string   `json:"parentType,omitempty"`
	ParentName           string   `json:"parentName,omitempty"`
	IsfFollowed          bool     `json:"isFollowed,omitempty"`
	FollowersIds         []string `json:"followersIds,omitempty"`
	CreatedById          string   `json:"createdById,omitempty"`
	AssignedUserId       string   `json:"assignedUserId,omitempty"`
	//Deleted              bool                 `json:"deleted,omitempty"`
	//Createdat            string               `json:"createdAt,omitempty"`
	//Modifiedat           string               `json:"modifiedAt,omitempty"`
	//Createdbyname        string               `json:"createdByName,omitempty"`
	//Modifiedbyid         string               `json:"modifiedById,omitempty"`
	//Modifiedbyname       string               `json:"modifiedByName,omitempty"`
	//Assignedusername     string               `json:"assignedUserName,omitempty"`
	//Teamsids             []string             `json:"teamsIds,omitempty"`
	//Teamsnames           Teamsnames           `json:"teamsNames,omitempty"`
	//Billingcontactsnames Billingcontactsnames `json:"billingContactsNames,omitempty"`
	//Fintransactionsnames Fintransactionsnames `json:"finTransactionsNames,omitempty"`
	//Commissionssnames    Commissionssnames    `json:"commissionssNames,omitempty"`
	//Followersnames       Followersnames       `json:"followersNames,omitempty"`
}

type DomainParcelsPutAuthCode struct {
	ID                 string `json:"id,omitempty"`
	AuthenticationCode string `json:"authenticationCode,omitempty"`
	InvoiceId          string `json:"invoiceId,omitempty"`
}
