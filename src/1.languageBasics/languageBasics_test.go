package languageBasics

import (
	"errors"
	"testing"
)

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil
}

// using functions
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("there is an error: ", err)
	}

	_, err2 := divide(10.0, 0)
	if err2 == nil {
		t.Error("there should be an error but the is no error here")
	}
}

// using data table
var tests = []struct {
	name     string
	x        float32
	y        float32
	expected float32
	isErr    bool
}{
	{"valid-data", 10.0, 1.0, 10.0, false},
	{"invalid-data", 10.0, 0.0, 0.0, true},
}

func TestDivisionWithDataTable(t *testing.T) {
	for _, tt := range tests {
		result, err := divide(tt.x, tt.y)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("didn't expected an error but got one", err.Error())
			}
		}
		if result != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, result)
		}
	}
}
