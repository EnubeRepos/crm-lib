package dashboardtemplate

const EntityDashboardTemplate string = "DashboardTemplate"

// ResponseDashboardTemplate .
type ResponseDashboardTemplate struct {
	Total             int                      `json:"total,omitempty,omitempty"`
	DashboardTemplates []DomainDashboardTemplate `json:"list,omitempty,omitempty"`
}

type DomainDashboardTemplate struct {
	ID                     string `json:"id,omitempty"`
}

type DomainPutDash4Users struct {
	ID         string   `json:"id,omitempty"`
	UserIdList []string `json:"userIdList,omitempty"`
	Append     bool     `json:"append,omitempty"`
}