package freshservice

import (
	"fmt"
	"time"
)

// Announcements represents a list of announcements in Freshservice
type Announcements struct {
	List []AnnouncementDetails `json:"announcements"`
}

// Announcement represents an announcment in Freshservice
type Announcement struct {
	Details AnnouncementDetails `json:"announcement"`
}

// AnnouncementDetails represents the specific details about a Freshservice announcement
type AnnouncementDetails struct {
	Title            string    `json:"title"`
	Body             string    `json:"body"`
	BodyHTML         string    `json:"body_html"`
	VisibleFrom      time.Time `json:"visible_from"`
	VisibleTill      time.Time `json:"visible_till"`
	Visibility       string    `json:"visibility"`
	Departments      []int     `json:"departments"`
	Groups           []int     `json:"groups"`
	State            string    `json:"state"`
	IsRead           bool      `json:"is_read"`
	SendEmail        bool      `json:"send_email"`
	AdditionalEmails []string  `json:"additional_emails"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedBy        int       `json:"created_by"`
}

// AnnouncementListFilter represents a filter that is available
// when listing Freshservice announcements
type AnnouncementListFilter struct {
	State string
}

// QueryString allows the available filter items to meet the QueryFilter interface
func (af *AnnouncementListFilter) QueryString() string {
	return fmt.Sprintf("state=%s", af.State)
}
