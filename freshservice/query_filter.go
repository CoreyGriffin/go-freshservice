package freshservice

// QueryFilter is an interface that can be passed around
// to Freshservice API methods that can accept a query param filter
type QueryFilter interface {
	// QueryString should take return string with the query parameters attached
	QueryString() string
}
