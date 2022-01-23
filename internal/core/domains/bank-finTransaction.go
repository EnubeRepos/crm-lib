package domains

type FinTransaction struct {
	Source string `json:"source,omitempty"`
	ID     string `json:"id,omitempty"`
}
