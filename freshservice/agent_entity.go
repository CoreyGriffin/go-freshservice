package freshservice

import (
	"fmt"
	"strings"
	"time"
)

// Agents holds a list of Freshservice agents
type Agents struct {
	List []AgentDetails `json:"agents"`
}

// Agent holds the details of a specific Freshservice agent
type Agent struct {
	Details AgentDetails `json:"agent"`
}

// AgentDetails contains the details of a specific Freshservice agent
type AgentDetails struct {
	ID                    int         `json:"id"`
	FirstName             string      `json:"first_name"`
	LastName              string      `json:"last_name"`
	Occasional            bool        `json:"occasional"`
	Active                bool        `json:"active"`
	JobTitle              string      `json:"job_title"`
	Email                 string      `json:"email"`
	WorkPhoneNumber       string      `json:"work_phone_number"`
	MobilePhoneNumber     string      `json:"mobile_phone_number"`
	ReportingManagerID    int         `json:"reporting_manager_id"`
	Address               string      `json:"address"`
	TimeZone              string      `json:"time_zone"`
	TimeFormat            string      `json:"time_format"`
	Language              string      `json:"language"`
	LocationID            int         `json:"location_id"`
	BackgroundInformation string      `json:"background_information"`
	ScoreboardLevelID     int         `json:"scoreboard_level_id"`
	GroupIds              []int       `json:"group_ids"` // being deprecated by freshservice
	MemberOf              []int       `json:"member_of"`
	ObserverOf            []int       `json:"observer_of"`
	RoleIds               []int       `json:"role_ids"` // being deprecated by freshservice
	Roles                 []AgentRole `json:"roles"`
	LastLoginAt           time.Time   `json:"last_login_at"`
	LastActiveAt          time.Time   `json:"last_active_at"`
	CustomFields          struct {
		House string `json:"house"`
	} `json:"custom_fields"`
	HasLoggedIn bool `json:"has_logged_in"`
}

// AgentRole represents a Freshservice role that can be assigned to an agent
type AgentRole struct {
	RoleID          int    `json:"role_id"`
	AssignmentScope string `json:"assignment_scope"`
	Groups          []int  `json:"groups"`
}

// Validate will confirm that an agent role is valud
func (ar *AgentRole) Validate() error {
	validScopes := []string{
		"entire_helpdesk",
		"member_groups",
		"assigned_items",
		"speciﬁed_groups",
	}

	if !StringInSlice(ar.AssignmentScope, validScopes) {
		return fmt.Errorf("Agent assignment scope is invalid; choose from %s", strings.Join(validScopes, ","))
	}

	if len(ar.Groups) > 0 && ar.AssignmentScope != "specified_groups" {
		return fmt.Errorf("Agent role groups are only applicable only if speciﬁed_groups is selected not %s", ar.AssignmentScope)
	}

	return nil
}

// AgentListFilter holds the filters available when listing Freservice agents
type AgentListFilter struct {
	PageQuery   string
	Email       *string
	MobilePhone *int
	WorkPhone   *int
	Active      bool
	Fulltime    bool
	Occasional  bool
}

// QueryString allows the available filter items to meet the QueryFilter interface
func (af *AgentListFilter) QueryString() string {
	var qs []string
	if af.PageQuery != "" {
		qs = append(qs, af.PageQuery)
	}

	switch {
	case af.Email != nil:
		qs = append(qs, fmt.Sprintf("email=%s", *af.Email))
	case af.MobilePhone != nil:
		qs = append(qs, fmt.Sprintf("mobile_phone_number=%d", *af.MobilePhone))
	case af.WorkPhone != nil:
		qs = append(qs, fmt.Sprintf("work_phone_number=%d", *af.WorkPhone))
	case af.Active:
		qs = append(qs, fmt.Sprintf("active=%v", af.Active))
	case af.Fulltime:
		qs = append(qs, "state=fulltime")
	case af.Occasional:
		qs = append(qs, "state=occasional")
	}
	return strings.Join(qs, "&")
}
