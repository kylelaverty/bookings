package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postData := url.Values{}
	form := New(postData, "POST")

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postData := url.Values{}
	form := New(postData, "POST")

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postData = url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	form = New(postData, "POST")
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("got invalid when required fields are present")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(postData, "POST")

	hasA := form.Has("a")
	if hasA {
		t.Error("form shows has field when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "a")
	form = New(postData, "POST")

	hasA = form.Has("a")
	if !hasA {
		t.Error("form shows no field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues, "POST")

	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues, "POST")

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("form shows min length for field that is too short")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues, "POST")

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("form shows field does not meet min length when it does")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues, "POST")

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "bob@test.com")
	form = New(postedValues, "POST")

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got invalid when should have been valid")
	}

}
