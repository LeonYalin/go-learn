package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}


// mock interface for NoSurf
type httpHandlerMock struct{}

func (mh *httpHandlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}