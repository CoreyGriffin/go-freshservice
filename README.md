# go-freshservice

[![Go Reference](https://pkg.go.dev/badge/github.com/CoreyGriffin/go-freshservice/freshservice.svg)](https://pkg.go.dev/github.com/CoreyGriffin/go-freshservice/freshservice)

go-freshservice is an unofficial Golang API client that aims to provide access to most facets of the [Freshservice API](https://api.freshservice.com/v2/#introduction)

## Usage

```go
import fs "github.com/CoreyGriffin/go-freshservice/freshservice"

// You can optionally setup a custom HTTP client to use which can
// include any settings you desire. If you would like to use the
// default client configuration just pass nil. This will default
// to a client that is simply configured with a timeout of 1 minute
myCustomHTTPClient := &http.Client{
  Timeout: time.Minute,
}

ctx := context.Background()

// Create a new client instance for making API calls
api, err := fs.New(ctx, "example.com", "my-cool-API-key", myCustomHTTPClient)
if err != nil {
  log.Fatal(err)
}

// List all tickets
t, err := api.Tickets().List(ctx, nil)
if err != nil {
  log.Fatal(err)
}

// Add optional filter to list call
filter := &fs.TicketListOptions{
  FilterBy: &fs.TicketFilter{
    RequesterEmail: fs.String("test-account@example.com"),
  },
  SortBy: &fs.SortOptions{
    Descending: true,
  },
}

ft, err := api.Tickets().List(ctx, filter)
if err != nil {
  log.Fatal(err)
}
```

## Contributing

Refer to [CONTRIBUTING.md](./CONTRIBUTING.md)
