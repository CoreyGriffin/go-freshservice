package freshservice

import (
	"context"
	"net/http"
	"net/url"
)

const ticketURL = "/api/v2/tickets"

// TicketService is an interface for interacting with
// the ticket endpoints of the Freshservice API
type TicketService interface {
	List(context.Context, QueryFilter) ([]TicketDetails, error)
	Create() (*Ticket, error)
	CreateWithAttachment() (*Ticket, error)
	Get() (*Ticket, error)
	Update() (*Ticket, error)
	Delete() (*Ticket, error)
}

// TicketServiceClient facilitates requests with the TicketService methods
type TicketServiceClient struct {
	client *Client
}

// TODO: extract to private method to incorporate pagination

// List all Freshservice tickets
// All the below requests are paginated to return only 30 tickets per page.
// Append the parameter "page=[:page_no]" in the url to traverse through pages.
func (t *TicketServiceClient) List(ctx context.Context, filter QueryFilter) ([]TicketDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   ticketURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &TicketList{}
	if err := t.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.Tickets, nil
}

// Create a new Freshservice ticket
func (t *TicketServiceClient) Create() (*Ticket, error) {
	return nil, nil
}

// CreateWithAttachment creates new Freshservice ticket with attachment
func (t *TicketServiceClient) CreateWithAttachment() (*Ticket, error) {
	return nil, nil
}

// Get a specific Freshservice ticket
func (t *TicketServiceClient) Get() (*Ticket, error) {
	return nil, nil
}

// Update a Freshservice ticket
func (t *TicketServiceClient) Update() (*Ticket, error) {
	return nil, nil
}

// Delete Freshservice ticket
func (t *TicketServiceClient) Delete() (*Ticket, error) {
	return nil, nil
}
