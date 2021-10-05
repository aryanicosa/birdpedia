package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal()
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler) //handler got from main.go

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `Halo dunia!`
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("handler return unexpected body : got %v want %v", actual, expected)
	}
}