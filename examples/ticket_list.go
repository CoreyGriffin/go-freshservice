package main

import (
	"context"
	"fmt"
	"log"

	fs "github.com/CoreyGriffin/go-freshservice/freshservice"
)

func main() {
	APIKey := "testing-123"
	ctx := context.Background()
	api, err := fs.New(ctx, "example.com", APIKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	// List all tickets
	// https://example.com/api/v2/tickets
	t, err := api.Tickets().List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tList := []string{}
	for _, tick := range t {
		tList = append(tList, fmt.Sprintf("\n%d - %d", tick.ID, tick.ResponderID))
	}

	fmt.Printf("All Tickets:\nCount: %d\nResults: %v\n", len(t), tList)

	// List tickets using a built in filer query and sort by
	f := &fs.TicketListOptions{
		FilterBy: &fs.TicketFilterOptions{
			RequesterEmail: fs.String("test-account@example.com"),
		},
		SortBy: &fs.SortOptions{
			Descending: true,
		},
	}

	// https://example.com/api/v2/tickets?email=test-account@example.com&order_type=desc
	ftList := []string{}
	ft, err := api.Tickets().List(ctx, f)
	if err != nil {
		log.Fatal(err)
	}

	for _, ft := range ft {
		ftList = append(ftList, fmt.Sprintf("\n%d - %d - %v", ft.ID, ft.RequesterID, ft.CreatedAt))
	}

	fmt.Printf("Filtered Tickets:\nCount: %d\nResults: %v\n", len(ft), ftList)
}
