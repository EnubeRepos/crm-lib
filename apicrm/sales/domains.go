package sales

const EntitySales string = "Sales"

// ResponseSales .
type ResponseSales struct {
	Total        int                  `json:"total,omitempty"`
	Sales []DomainSales `json:"list,omitempty"`
}

// Account
type DomainSales struct{}
