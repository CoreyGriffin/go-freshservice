package freshservice

import (
	"net/http"
	"net/url"
	"strings"
)

// Int is a built in utility function that will return a *int
func Int(i int) *int {
	return &i
}

// String is a built in utilty function that will return a *string
func String(s string) *string {
	return &s
}

// StringInSlice is a utility function that can be used to see if a string
// exists in a static list of strings
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// HasNextPage will take in an http response and check
// for the existence of the "link" header to determine whether or
// not there is another page returning the next page's URL
// <https://example.freshservice.com/api/v2/tickets?page=2>; rel="next"
func HasNextPage(resp *http.Response) string {
	link, ok := resp.Header["Link"]
	if !ok {
		return ""
	}

	// pull out raw url
	u := strings.Split(link[0], ";")[0]
	// drop the carrots < >
	u = u[1 : len(u)-1]

	return ParseNextPage(u)
}

// ParseNextPage will return the next page parameter parsed
// out of a raw URL string's "page=[:page_no]" parameter
func ParseNextPage(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.RawQuery
}
