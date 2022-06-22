package sales

const EntitySales string = "Sales"

// ResponseSales .
type ResponseSales struct {
	Total int           `json:"total,omitempty"`
	Sales []DomainSales `json:"list,omitempty"`
}

// Account //mising stuff? // Do I create it? How do I know what fields I have to inicialize

//didn't have any field

type DomainSales struct {
	ID string `json:"id,omitempty"`
	//Name                  string  `json:"name,omitempty"`
	AssignedUser          string  `json:"assignedUser,omitempty"`
	AssignedUserID        string  `json:"assignedUserId,omitempty"`
	AssignedUserName      string  `json:"assignedUserName,omitempty"`
	SalesNumber           string  `json:"salesNumber,omitempty"`
	SalesDate             string  `json:"salesDate,omitempty"`
	CommissionTotalAmount float32 `json:"commissionsTotalAmount,omitempty"`
}
