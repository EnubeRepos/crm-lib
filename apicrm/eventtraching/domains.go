package eventtraching

const EntityEventTraching string = "EventTraching"

// ResponseEventTraching .
type ResponseEventTraching struct {
	Total        int                  `json:"total,omitempty"`
	EventTraching []DomainEventTraching `json:"list,omitempty"`
}

// Account
type DomainEventTraching struct{}
