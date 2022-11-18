package forms

import (
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	var formData url.Values = make(url.Values)
	formData.Add("a", "a")

	form := New(formData)
	if !form.Has("a") {
		t.Error("form expected to have a specific field")
	}

	formData.Del("a")
	form = New(formData)
	if form.Has("a") {
		t.Error("form should not be valid")
	}
}

func TestForm_Valid(t *testing.T) {
	var formData url.Values

	form := New(formData)
	if !form.Valid() {
		t.Error("form should be valid")
	}

	form.Errors.Add("dummy", "error")
	if form.Valid() {
		t.Error("form should not be valid")
	}
}

func TestForm_Required(t *testing.T) {
	var formData url.Values = make(url.Values)
	formData.Add("a", "a")
	formData.Add("b", "b")
	formData.Add("c", "c")

	form := New(formData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form fields are required but not found")
	}

	formData.Del("c")
	form = New(formData)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form should not be valid")
	}
}

func TestForm_MinLenght(t *testing.T) {
	var formData url.Values = make(url.Values)
	formData.Add("a", "1234")

	form := New(formData)
	form.MinLength("a", 3)
	if !form.Valid() {
		t.Error("field minlength is not valid")
	}

	form.MinLength("a", 5)
	if form.Valid() {
		t.Error("form should not be valid")
	}
}

func TestForm_Email(t *testing.T) {
	var formData url.Values = make(url.Values)
	formData.Add("a", "a@a.gmail.com")

	form := New(formData)
	form.IsEmail("a")
	if !form.Valid() {
		t.Error("form field should be an email")
	}

	formData.Del("a")
	formData.Add("a", "lala")
	form = New(formData)
	form.IsEmail("a")
	if form.Valid() {
		t.Error("form should not be valid")
	}
}
