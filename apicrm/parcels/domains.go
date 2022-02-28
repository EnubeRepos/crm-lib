package parcels

const EntityParcels string = "Parcels"

// ResponseParcels .
type ResponseParcels struct {
	Total   int             `json:"total,omitempty"`
	Parcels []DomainParcels `json:"list,omitempty"`
}

// Account
type DomainParcels struct{}
