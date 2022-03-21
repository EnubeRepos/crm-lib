package attachment

const EntityAttachment string = "Attachment"

// ResponseAttachment .
type ResponseAttachment struct {
	Total      int                `json:"total,omitempty"`
	Attachment []DomainAttachment `json:"list,omitempty"`
}

// Account
type DomainAttachment struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Size        int    `json:"size,omitempty"`
	Role        string `json:"role,omitempty"`
	RelatedType string `json:"relatedType,omitempty"`
	File        string `json:"file,omitempty"`
	Field       string `json:"field,omitempty"`
}
