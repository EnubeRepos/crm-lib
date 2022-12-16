package sales

const EntitySales string = "Sales"

// ResponseSales .
type ResponseSales struct {
	Total int           `json:"total,omitempty"`
	Sales []DomainSales `json:"list,omitempty"`
}

// Account
type DomainSales struct {
	ID                     string   `json:"id"`
	Name                   string   `json:"name"`
	Commissionstotalamount float64  `json:"commissionsTotalAmount"`
	Status                 string   `json:"status"`
	Commissionsids         []string `json:"commissionsIds"`
	Accountid              string   `json:"accountId"`
	Parcelsids             []string `json:"parcelsIds"`
	AssignedUser           string   `json:"assignedUser,omitempty"`
}
