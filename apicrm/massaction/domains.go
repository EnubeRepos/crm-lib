package massaction

type DomainRecalc struct {
	EntityType string `json:"entityType,omitempty"`
	Action     string `json:"action,omitempty"`
	Params     Params `json:"params,omitempty"`
}
type Params struct {
	Ids []string `json:"ids,omitempty"`
}

type ResponseRecalc struct {
	Count int      `json:"count,omitempty"`
	Ids   []string `json:"ids,omitempty"`
}