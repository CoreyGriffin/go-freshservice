package freshservice

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
