package domains

type Webhook struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
	ID     string `json:"id,omitempty"`
}
