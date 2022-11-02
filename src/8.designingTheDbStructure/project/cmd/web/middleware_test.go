package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mockHandler httpHandlerMock
	myHandler := NoSurf(&mockHandler)
	switch v := myHandler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var mockHandler httpHandlerMock
	myHandler := SessionLoad(&mockHandler)
	switch v := myHandler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but %T", v)
	}
}
