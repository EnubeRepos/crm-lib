package institutionfinancial

const EntityInstitutionFinancial string = "InstitutionFinancial"

// ResponseInstitutionFinancial .
type ResponseInstitutionFinancial struct {
	Total int         `json:"total,omitempty"`
	InstitutionFinancial   []DomainInstitutionFinancial `json:"list,omitempty"`
}

// Account
type DomainInstitutionFinancial struct {
	ID                     string                 `json:"id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	//Deleted                bool                   `json:"deleted,omitempty"`
	Description            string            `json:"description,omitempty"`
	//CreatedAt              string                 `json:"createdAt,omitempty"`
	//ModifiedAt             string                 `json:"modifiedAt,omitempty"`
	Ispb                   string                 `json:"ispb,omitempty"`
	Code                   string                 `json:"code,omitempty"`
	ShortName              string                 `json:"shortName,omitempty"`
	IsSPIDirect            bool                   `json:"isSPIDirect,omitempty"`
	Products               []string               `json:"products,omitempty"`
}