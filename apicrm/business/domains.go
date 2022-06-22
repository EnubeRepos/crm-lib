package business

const EntityBusiness string = "BusinessRegistration"

type ResponseBusiness struct {
	Total   int              `json:"total,omitempty"`
	Account []DomainBusiness `json:"list,omitempty"`
}

// Registration
type DomainBusiness struct {
	ID                         string `json:"id,omitempty"` // I created
	DocumentNumber             string `json:"documentNumber,omitempty"`
	BillingAddressStreet       string `json:"billingAddressStreet,omitempty"`
	BillingAddressCity         string `json:"billingAddressCity,omitempty"`
	BillingAddressState        string `json:"billingAddressState,omitempty"`
	BillingAddressCountry      string `json:"billingAddressCountry,omitempty"`
	BillingAddressPostalCode   string `json:"billingAddressPostalCode,omitempty"`
	BillingAddressNumber       string `json:"billingAddressNumber,omitempty"`
	BillingAddressComplement   string `json:"billingAddressComplement,omitempty"`
	BillingAddressNeighborhood string `json:"billingAddressNeighborhood,omitempty"`
	EmailAddress               string `json:"emailAddress,omitempty"`
	TradingName                string `json:"tradingName,omitempty"`
	Name                       string `json:"name,omitempty"`
	PhoneNumber                string `json:"phoneNumber,omitempty"`
	RegistrationID             string `json:"registrationId,omitempty"`
}
