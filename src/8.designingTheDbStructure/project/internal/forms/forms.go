package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string) bool {
	has := f.Get(field)
	if has == "" {
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "field is required")
		}
	}
}

func (f *Form) MinLength(field string, lenght int) {
	if len(f.Get(field)) < lenght {
		f.Errors.Add(field, fmt.Sprintf("the field must be as least %d characters long", lenght))
	}
}

// uses go get github.com/asaskevich/govalidator
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
