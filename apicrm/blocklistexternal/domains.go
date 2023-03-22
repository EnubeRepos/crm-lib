package blocklistexternal

const EntityBlockListExternal string = "BlockListExternal"

// ResponseBlockListExternal .
type ResponseBlockListExternal struct {
	Total             int                       `json:"total,omitempty"`
	BlockListExternal []DomainBlockListExternal `json:"list,omitempty"`
}

// Account
type DomainBlockListExternal struct {
	ID                              string `json:"id,omitempty"`
	Name                            string `json:"name,omitempty"`
	Deleted                         bool   `json:"deleted,omitempty"`
	Description                     string `json:"description,omitempty"`
	CreatedAt                       string `json:"createdAt,omitempty"`
	ModifiedAt                      string `json:"modifiedAt,omitempty"`
	JudicialBlockID                 string `json:"judicialBlockId,omitempty"`
	Status                          string `json:"status,omitempty"`
	Details                         string `json:"details,omitempty"`
	BlockedAmount                   string `json:"blockedAmount,omitempty"`
	BlockedAmountvalue              float64 `json:"blockedAmountvalue,omitempty"`
	DetailsjudicialOrderAmount      string `json:"detailsjudicialOrderAmount,omitempty"`
	DetailsjudicialOrderAmountvalue float64 `json:"detailsjudicialOrderAmountvalue,omitempty"`
	Holder                          string `json:"holder,omitempty"`
	Holderdocument                  string `json:"holderdocument,omitempty"`
	Holderdocumentvalue             string `json:"holderdocumentvalue,omitempty"`
	Holderdocumenttype              string `json:"holderdocumenttype,omitempty"`
	Holdertype                      string `json:"holdertype,omitempty"`
	Holderaccount                   string `json:"holderaccount,omitempty"`
	Holderaccountbranch             string `json:"holderaccountbranch,omitempty"`
	Holderaccountnumber             string `json:"holderaccountnumber,omitempty"`
	Holderaccounttype               string `json:"holderaccounttype,omitempty"`
	Holderaccountstatus             string `json:"holderaccountstatus,omitempty"`
	Holderaccountbank               string `json:"holderaccountbank,omitempty"`
	Holderaccountbankispb           string `json:"holderaccountbankispb,omitempty"`
	Holderaccountbankcode           string `json:"holderaccountbankcode,omitempty"`
	Holderaccountbankname           string `json:"holderaccountbankname,omitempty"`
	Holderbalances                  string `json:"holderbalances,omitempty"`
	Holderbalancesavailable         string `json:"holderbalancesavailable,omitempty"`
	Holderbalancesvalue             float64 `json:"holderbalancesvalue,omitempty"`
	Holderbalancesblocked           string `json:"holderbalancesblocked,omitempty"`
	Holderbalancesblockedvalue      float64 `json:"holderbalancesblockedvalue,omitempty"`
	DetailslawsuitNumber            string `json:"detailslawsuitNumber,omitempty"`
	//BlockedAmountvalueCurrency               string `json:"blockedAmountvalueCurrency,omitempty"`
	//DetailsjudicialOrderAmountvalueCurrency  string `json:"detailsjudicialOrderAmountvalueCurrency,omitempty"`
	//HolderbalancesvalueCurrency              string `json:"holderbalancesvalueCurrency,omitempty"`
	//HolderbalancesblockedvalueCurrency       string `json:"holderbalancesblockedvalueCurrency,omitempty"`
	CreatedByID      string `json:"createdById,omitempty"`
	CreatedByName    string `json:"createdByName,omitempty"`
	ModifiedByID     string `json:"modifiedById,omitempty"`
	ModifiedByName   string `json:"modifiedByName,omitempty"`
	AssignedUserID   string `json:"assignedUserId,omitempty"`
	AssignedUserName string `json:"assignedUserName,omitempty"`
	//TeamsIds                                 []string    `json:"teamsIds,omitempty"`
	//TeamsNames                               TeamsNames  `json:"teamsNames,omitempty"`
	//BlockedAmountvalueConverted              string `json:"blockedAmountvalueConverted,omitempty"`
	//DetailsjudicialOrderAmountvalueConverted string `json:"detailsjudicialOrderAmountvalueConverted,omitempty"`
	//HolderbalancesvalueConverted             string `json:"holderbalancesvalueConverted,omitempty"`
	//HolderbalancesblockedvalueConverted      string `json:"holderbalancesblockedvalueConverted,omitempty"`
	BankAccountID   string `json:"bankAccountId,omitempty"`
	BankAccountName string `json:"bankAccountName,omitempty"`
}
