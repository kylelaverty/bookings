package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.Values object to hold form data, and adds a method and errors field to the struct.
type Form struct {
	// URL is the form action URL
	url.Values
	// Method is the form submission method
	Method string
	// Errors contains any validation errors for the form
	Errors errors
}

// New initializes a form struct.
func New(data url.Values) *Form {
	return &Form{
		data,
		"POST",
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	return x != ""
}
