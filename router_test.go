package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestHelloRouter(t *testing.T) {
	server := httptest.NewServer(newRouter())
	response, err := http.Get(server.URL)

	if err != nil{
		t.Fatal("Bugger, i got an error doing a request", err)
	}

	if response.StatusCode != http.StatusOK{
		t.Error("Expected a 200 but i got", response.StatusCode)
	}
}