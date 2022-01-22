package domains

type Webhook struct {
	Source string `json:"source"`
	Target string `json:"target"`
	ID     string `json:"id"`
}
