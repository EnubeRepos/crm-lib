package fintransaction

const EntityFinTransaction string = "FinTransaction"

// ResponseFinTransaction .
type ResponseFinTransaction struct {
	Total          int                    `json:"total,omitempty"`
	FinTransaction []DomainFinTransaction `json:"list,omitempty"`
}

// FinTransaction
type DomainFinTransaction struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Deleted             bool     `json:"deleted"`
	Status              string   `json:"status"`
	Datestart           string   `json:"dateStart"`
	Dateend             string   `json:"dateEnd"`
	Isallday            bool     `json:"isAllDay"`
	Duration            int      `json:"duration"`
	Description         string   `json:"description"`
	Reminders           []string `json:"reminders"`
	Createdat           string   `json:"createdAt"`
	Modifiedat          string   `json:"modifiedAt"`
	Transactiontype     string   `json:"transactionType"`
	Value               float64  `json:"value"`
	Commission          bool     `json:"commission"`
	Accounttransaction  string   `json:"accountTransaction"`
	Transactioncategory string   `json:"transactionCategory"`
	Amount              float64  `json:"amount"`
	Confirmtransaction  bool     `json:"confirmTransaction"`
	Scheduletransaction string   `json:"scheduleTransaction"`
	Externalorigin      string   `json:"externalOrigin"`
	Datestartdate       string   `json:"dateStartDate"`
	Dateenddate         string   `json:"dateEndDate"`
	Valuecurrency       string   `json:"valueCurrency"`
	Parentid            string   `json:"parentId"`
	Parenttype          string   `json:"parentType"`
	Parentname          string   `json:"parentName"`
	Createdbyid         string   `json:"createdById"`
	Createdbyname       string   `json:"createdByName"`
	Modifiedbyid        string   `json:"modifiedById"`
	Modifiedbyname      string   `json:"modifiedByName"`
	Assigneduserid      string   `json:"assignedUserId"`
	Assignedusername    string   `json:"assignedUserName"`
	Teamsids            []string `json:"teamsIds"`
	//Teamsnames                  Teamsnames           `json:"teamsNames"`
	Salesorderitemsids []string `json:"salesOrderItemsIds"`
	//Salesorderitemsnames        Salesorderitemsnames `json:"salesOrderItemsNames"`
	Statustrackingid            string   `json:"statusTrackingId"`
	Statustrackingname          string   `json:"statusTrackingName"`
	Valueconverted              string   `json:"valueConverted"`
	Salesorderid                string   `json:"salesOrderId"`
	Salesordername              string   `json:"salesOrderName"`
	Bankaccountid               string   `json:"bankAccountId"`
	Bankaccountname             string   `json:"bankAccountName"`
	Commissionsid               string   `json:"commissionsId"`
	Commissionsname             string   `json:"commissionsName"`
	Documentid                  string   `json:"documentId"`
	Documentname                string   `json:"documentName"`
	Parcelid                    string   `json:"parcelId"`
	Parcelname                  string   `json:"parcelName"`
	Operationfeeapplicationsids []string `json:"operationFeeApplicationsIds"`
	//Operationfeeapplicationsnames Operationfeeapplicationsnames `json:"operationFeeApplicationsNames"`
	Contactid                        string  `json:"contactId"`
	Contactname                      string  `json:"contactName"`
	Transactionchildrenid            string  `json:"transactionChildrenId"`
	Transactionchildrentype          string  `json:"transactionChildrenType"`
	Transactionchildrenname          string  `json:"transactionChildrenName"`
	Entityid                         string  `json:"entityid"`
	Idempotencykey                   string  `json:"idempotencykey"`
	Companykey                       string  `json:"companykey"`
	Context                          string  `json:"context"`
	Timestamp                        string  `json:"timestamp"`
	Correlationid                    string  `json:"correlationid"`
	Version                          string  `json:"version"`
	Tedid                            string  `json:"tedId"`
	Tedname                          string  `json:"tedName"`
	Userid                           string  `json:"userId"`
	Username                         string  `json:"userName"`
	Pixid                            string  `json:"pixId"`
	Pixname                          string  `json:"pixName"`
	Pixkeyid                         string  `json:"pixKeyId"`
	Pixkeyname                       string  `json:"pixKeyName"`
	InstitutionFinancialRecipientId  string  `json:"institutionFinancialRecipientId"`
	InstitutionFinancialSenderId     string  `json:"institutionFinancialSenderId"`
	Authenticationcode               string  `json:"authenticationCode"`
	Recipientdocumentvalue           string  `json:"recipientDocumentValue"`
	Recipientdocumenttype            string  `json:"recipientDocumentType"`
	Recipienttype                    string  `json:"recipientType"`
	Recipientname                    string  `json:"recipientName"`
	Recipientaccountbranch           string  `json:"recipientAccountBranch"`
	Recipientaccountnumber           string  `json:"recipientAccountNumber"`
	Recipientaccountbankispb         string  `json:"recipientAccountBankIspb"`
	Recipientaccountbankcode         string  `json:"recipientAccountBankCode"`
	Recipientaccountbankname         string  `json:"recipientAccountBankName"`
	Recipientaccountbalancevalue     float64 `json:"recipientAccountBalanceValue"`
	Recipientaccountbalancecurrrency string  `json:"recipientAccountBalanceCurrrency"`
	Recipientstatus                  string  `json:"recipientStatus"`
	Channelname                      string  `json:"channelName"`
	Channelcontrolnumber             string  `json:"channelControlNumber"`
	Channelcontrolnumberoriginal     string  `json:"channelControlNumberOriginal"`
	Channelownernumber               string  `json:"channelOwnerNumber"`
	Senderdocumentvalue              string  `json:"senderDocumentValue"`
	Senderdocumenttype               string  `json:"senderDocumentType"`
	Sendertype                       string  `json:"senderType"`
	Sendername                       string  `json:"senderName"`
	Senderaccountbranch              string  `json:"senderAccountBranch"`
	Senderaccountnumber              string  `json:"senderAccountNumber"`
	Senderaccountbankispb            string  `json:"senderAccountBankIspb"`
	Senderaccountbankcode            string  `json:"senderAccountBankCode"`
	Senderaccountbankname            string  `json:"senderAccountBankName"`
	Senderstatus                     string  `json:"senderStatus"`
}

//type Teamsnames struct {
//	Six18D2B386E52D46A6 string `json:"618d2b386e52d46a6"`
//}
//type Salesorderitemsnames struct {
//}
//type Operationfeeapplicationsnames struct {
//}
