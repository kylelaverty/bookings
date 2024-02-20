package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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

// Valid returns true if there are no errors, otherwise false.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct.
func New(data url.Values, method string) *Form {
	return &Form{
		data,
		method,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields in the form.
// ... is a variadic parameter, which allows us to pass in any number of strings.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty.
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	return x != ""
}

// MinLength checks for a minimum length of a form field.
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", length))
		return false
	}
	return true
}

// IsEmail checks for a valid email address.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
