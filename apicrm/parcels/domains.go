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
	Productname          string   `json:"productName,omitempty"`
	Productunity         string   `json:"productUnity,omitempty"`
	Salesdate            string   `json:"salesDate,omitempty"`
	Salestatus           string   `json:"saleStatus,omitempty"`
	Parcelnumber         int      `json:"parcelNumber,omitempty"`
	Parcelcode           string   `json:"parcelCode,omitempty"`
	Productdivision      string   `json:"productDivision,omitempty"`
	Status               string   `json:"status,omitempty"`
	Duedate              string   `json:"dueDate,omitempty"`
	Amount               int      `json:"amount,omitempty"`
	Payment              string   `json:"payment,omitempty"`
	Parcelpaid           string   `json:"parcelPaid,omitempty"`
	Returnfee            string   `json:"returnFee,omitempty"`
	Approvereturntranfer bool     `json:"approveReturnTranfer,omitempty"`
	Invoicenumber        string   `json:"invoiceNumber,omitempty"`
	Authenticationcode   string   `json:"authenticationCode,omitempty"`
	Amountcurrency       string   `json:"amountCurrency,omitempty"`
	Returnfeecurrency    string   `json:"returnFeeCurrency,omitempty"`
	Salesid              string   `json:"salesId,omitempty"`
	Salesname            string   `json:"salesName,omitempty"`
	Billingcontactsids   []string `json:"billingContactsIds,omitempty"`
	Fintransactionsids   []string `json:"finTransactionsIds,omitempty"`
	Invoiceid            string   `json:"invoiceId,omitempty"`
	Invoicename          string   `json:"invoiceName,omitempty"`
	Amountconverted      int      `json:"amountConverted,omitempty"`
	Commissionssids      []string `json:"commissionssIds,omitempty"`
	Returnfeeconverted   string   `json:"returnFeeConverted,omitempty"`
	Billingcontactid     string   `json:"billingContactId,omitempty"`
	Billingcontactname   string   `json:"billingContactName,omitempty"`
	Parentid             string   `json:"parentId,omitempty"`
	Parenttype           string   `json:"parentType,omitempty"`
	Parentname           string   `json:"parentName,omitempty"`
	Isfollowed           bool     `json:"isFollowed,omitempty"`
	Followersids         []string `json:"followersIds,omitempty"`
	//Deleted              bool                 `json:"deleted,omitempty"`
	//Createdat            string               `json:"createdAt,omitempty"`
	//Modifiedat           string               `json:"modifiedAt,omitempty"`
	//Createdbyid          string               `json:"createdById,omitempty"`
	//Createdbyname        string               `json:"createdByName,omitempty"`
	//Modifiedbyid         string               `json:"modifiedById,omitempty"`
	//Modifiedbyname       string               `json:"modifiedByName,omitempty"`
	//Assigneduserid       string               `json:"assignedUserId,omitempty"`
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
}
