package freshservice

import (
	"fmt"
	"strings"
	"time"
)

const (
	// TicketOpen is the value required to indicate a ticket status is open
	TicketOpen = 2
	// TicketPending is the value required to indicate a ticket status is pending
	TicketPending = 3
	// TicketResolved is the value required to indicate a ticket status is resolved
	TicketResolved = 4
	// TicketClosed is the value required to indicate a ticket status is closed
	TicketClosed = 5
	// LowPriority is the value to set a ticket priority to low
	LowPriority = 1
	// MediumPriority is the value to set a ticket priority to medium
	MediumPriority = 2
	// HighPriority is the value to set a ticket priority to high
	HighPriority = 3
	// UrgentPriority is the value to set a ticket priority to urgent
	UrgentPriority = 4
	// SourceEmail is the value to specify a ticket was opened via email
	SourceEmail = 1
	// SourcePortal is the value to specify a ticket was opened via portal
	SourcePortal = 2
	// SourcePhone is the value to specify a ticket was opened via phone
	SourcePhone = 3
	// SourceChat is the value to specify a ticket was opened via chat
	SourceChat = 4
	// SourceFeedbackWidget is the value to specify a ticket was opened via a Feedback Widget
	SourceFeedbackWidget = 5
	// SourceYammer is the value to specify a ticket was opened via Yammer
	SourceYammer = 6
	// SourceAWSCloudwatch is the value to specify a ticket was opened by AWS Cloudwatch
	SourceAWSCloudwatch = 7
	// SourcePagerduty is the value to specify a ticket was opened by Pagerduty
	SourcePagerduty = 8
	// SourceWalkup is the value to specify a ticket was opened via Walkup
	SourceWalkup = 9
	// SourceSlack is the value to specify a ticket was opened via Slack
	SourceSlack = 10
)

// TicketList holds a list of tickets returned from the Freshservice API
type TicketList struct {
	Tickets []TicketDetails `json:"tickets"`
}

// Ticket represents a Freshservice ticket object
type Ticket struct {
	Details TicketDetails `json:"ticket,omitempty"`
}

// TicketDetails contains the specific ticket details
type TicketDetails struct {
	CcEmails        []string     `json:"cc_emails"`
	FwdEmails       []string     `json:"fwd_emails"`
	ReplyCcEmails   []string     `json:"reply_cc_emails"`
	FrEscalated     bool         `json:"fr_escalated"`
	Spam            bool         `json:"spam"`
	EmailConfigID   int          `json:"email_config_id"`
	GroupID         int          `json:"group_id"`
	Priority        int          `json:"priority"`
	RequesterID     int          `json:"requester_id"`
	ResponderID     int          `json:"responder_id"`
	Source          int          `json:"source"`
	Status          int          `json:"status"`
	Subject         string       `json:"subject"`
	ToEmails        []string     `json:"to_emails"`
	SLAPolicyID     int          `json:"sla_policy_id"`
	DepartmentID    int          `json:"department_id"`
	ID              int          `json:"id"`
	Type            string       `json:"type"`
	DueBy           time.Time    `json:"due_by"`
	FrDueBy         time.Time    `json:"fr_due_by"`
	IsEscalated     bool         `json:"is_escalated"`
	Description     string       `json:"description"`
	DescriptionText string       `json:"description_text"`
	CustomFields    CustomFields `json:"custom_fields"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	Urgency         int          `json:"urgency"`
	Impact          int          `json:"impact"`
	Category        string       `json:"category"`
	SubCategory     string       `json:"sub_category"`
	ItemCategory    string       `json:"item_category"`
	Deleted         bool         `json:"deleted"`
	Attachments     []Attachment `json:"attachments"`
}

// CarbonCopy manages the emails to be copied in on a ticket
type CarbonCopy struct {
	CcEmails  []string `json:"cc_emails"`
	FwdEmails []string `json:"fwd_emails"`
	ReplyCc   []string `json:"reply_cc"`
	TktCc     []string `json:"tkt_cc"`
}

// TicketNote represents a note added to a Freshservice ticket
type TicketNote struct {
	Details TicketNoteDetails `json:"note"`
}

// TicketNoteDetails holds the details of a note added to a Freshservice ticket
type TicketNoteDetails struct {
	ID           int64        `json:"id"` // Read-Only
	UserID       int64        `json:"user_id"`
	Source       int          `json:"source"`
	Incoming     bool         `json:"incoming"`
	Private      bool         `json:"private"`
	CreatedAt    string       `json:"created_at"`
	UpdatedAt    string       `json:"updated_at"`
	Deleted      bool         `json:"deleted"`
	Body         string       `json:"body"`
	BodyHTML     string       `json:"body_html"`   // Mandatory
	Attachments  []Attachment `json:"attachments"` // Read-Only
	SupportEmail interface{}  `json:"support_email"`
}

// Attachment represents a ticket attachment
type Attachment struct {
	ContentType   string    `json:"content_type"`
	Size          int       `json:"size"`
	Name          string    `json:"name"`
	AttachmentURL string    `json:"attachment_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CustomFields holds a mapping of custom ticket fields
type CustomFields map[string]interface{}

// TicketListOptions holds the available options that can be
// passed when requesting a list of Freshservice ticketsx
type TicketListOptions struct {
	FilterBy *TicketFilterOptions
	SortBy   *SortOptions
	Embed    *TicketEmbedOptions
}

// TicketEmbedOptions will optonally embed desired metadata in a ticket list response
// Each include will consume an additional 2 credits. For example if you embed the stats
// information you will be charged a total of 3 API credits (1 credit for the API call, and 2 credits for the additional stats embedding).
type TicketEmbedOptions struct {
	Stats         bool
	RequesterInfo bool
}

// SortOptions will opitionally sort the ticket list results
type SortOptions struct {
	Ascending  bool
	Descending bool
}

// TicketFilterOptions are optional filters that can be enabled when querying a ticket list
type TicketFilterOptions struct {
	NewAndMyOpen   bool
	Watching       bool
	Spam           bool
	Deleted        bool
	RequesterID    *int
	RequesterEmail *string
	UpdatedSince   *time.Time
	Type           *string
}

// QueryString allows us to pass TicketListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *TicketListOptions) QueryString() string {
	var qs []string

	if opts.FilterBy != nil {
		switch {
		case opts.FilterBy.NewAndMyOpen:
			qs = append(qs, "filter=new_and_my_open")
		case opts.FilterBy.Watching:
			qs = append(qs, "filter=watching")
		case opts.FilterBy.Spam:
			qs = append(qs, "filter=spam")
		case opts.FilterBy.Deleted:
			qs = append(qs, "filter=deleted")
		}

		switch {
		case opts.FilterBy.RequesterID != nil:
			qs = append(qs, fmt.Sprintf("requester_id=%d", *opts.FilterBy.RequesterID))
		case opts.FilterBy.RequesterEmail != nil:
			qs = append(qs, fmt.Sprintf("email=%v", *opts.FilterBy.RequesterEmail))
		}

		if opts.FilterBy.UpdatedSince != nil {
			qs = append(qs, fmt.Sprintf("updated_since=%v", *opts.FilterBy.UpdatedSince))
		}

		if opts.FilterBy.Type != nil {
			qs = append(qs, fmt.Sprintf("type=%v", *opts.FilterBy.Type))
		}
	}

	if opts.SortBy != nil {
		if opts.SortBy.Ascending {
			qs = append(qs, "order_type=asc")
		} else {
			qs = append(qs, "order_type=desc")
		}
	}

	if opts.Embed != nil {
		if opts.Embed.RequesterInfo {
			qs = append(qs, "include=requester")
		}
		if opts.Embed.Stats {
			qs = append(qs, "include=stats")
		}
	}

	return strings.Join(qs, "&")
}
