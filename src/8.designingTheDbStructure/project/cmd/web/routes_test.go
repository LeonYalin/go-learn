package main

import (
	"testing"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	mux := Routes()

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but %T", v)
	}
}