package eventtracking

const EntityEventTracking string = "EventTracking"

// ResponseEventTracking .
type ResponseEventTracking struct {
	Total         int                   `json:"total,omitempty"`
	EventTracking []DomainEventTracking `json:"list,omitempty"`
}

type DomainEventTrackingBase struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AssignedUserID string `json:"assignedUserId"`
}

// Account
type DomainEventTracking struct {
	ID                        string   `json:"id"`
	Name                      string   `json:"name"`
	Deleted                   bool     `json:"deleted"`
	Description               string   `json:"description"`
	CreatedAt                 string   `json:"createdAt"`
	ModifiedAt                string   `json:"modifiedAt"`
	Action                    string   `json:"action"`
	StatusTrackingDescription string   `json:"statusTrackingDescription"`
	Status                    string   `json:"status"`
	QueueName                 string   `json:"queueName"`
	EntityName                string   `json:"entityName"`
	CreatedByID               string   `json:"createdById"`
	CreatedByName             string   `json:"createdByName"`
	ModifiedByID              string   `json:"modifiedById"`
	ModifiedByName            string   `json:"modifiedByName"`
	AssignedUserID            string   `json:"assignedUserId"`
	AssignedUserName          string   `json:"assignedUserName"`
	TeamsIds                  []string `json:"teamsIds"`
	StatusTrackingID          string   `json:"statusTrackingId"`
	StatusTrackingName        string   `json:"statusTrackingName"`
	RegistrationID            string   `json:"registrationId"`
	RegistrationName          string   `json:"registrationName"`
}
