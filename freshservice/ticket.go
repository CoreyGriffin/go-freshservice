package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const ticketURL = "/api/v2/tickets"

// TicketService is an interface for interacting with
// the ticket endpoints of the Freshservice API
type TicketService interface {
	List(context.Context, QueryFilter) ([]TicketDetails, error)
	Create(context.Context, *TicketDetails) (*TicketDetails, error)
	CreateWithAttachment() (*Ticket, error)
	Get(context.Context, int, QueryFilter) (*TicketDetails, error)
	Update(context.Context, int, *TicketDetails) (*TicketDetails, error)
	Delete() (*Ticket, error)
}

// TicketServiceClient facilitates requests with the TicketService methods
type TicketServiceClient struct {
	client *Client
}

// TODO: extract to private method to incorporate pagination using "link" header

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
func (t *TicketServiceClient) Create(ctx context.Context, td *TicketDetails) (*TicketDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   ticketURL,
	}

	ticketContent, err := json.Marshal(td)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(ticketContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Ticket{}
	if err := t.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// CreateWithAttachment creates new Freshservice ticket with attachment
func (t *TicketServiceClient) CreateWithAttachment() (*Ticket, error) {
	return nil, nil
}

// Get a specific Freshservice ticket by Ticket ID. By default, certain
// fields such as conversations, tags and requester email will not be included
// in the response. They can be retrieved via the embedding functionality.
func (t *TicketServiceClient) Get(ctx context.Context, id int, filter QueryFilter) (*TicketDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   fmt.Sprintf("%s/%d", ticketURL, id),
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Ticket{}
	if err := t.client.makeRequest(req, res); err != nil {
		return nil, err
	}
	return &res.Details, nil
}

// Update a Freshservice ticket
func (t *TicketServiceClient) Update(ctx context.Context, id int, details *TicketDetails) (*TicketDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   fmt.Sprintf("%s/%d", ticketURL, id),
	}

	ticketContent, err := json.Marshal(details)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(ticketContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Ticket{}
	if err := t.client.makeRequest(req, res); err != nil {
		return nil, err
	}
	return &res.Details, nil
}

// Delete Freshservice ticket
func (t *TicketServiceClient) Delete() (*Ticket, error) {
	return nil, nil
}
