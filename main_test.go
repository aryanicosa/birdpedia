package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	//hit "GET /assets/" route to get the index.html
	resp, err := http.Get(mockServer.URL + "/assets/")
	
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, but got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, but got %s", expectedContentType, contentType)
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, but got %d", resp.StatusCode)
	}

	//response body is read and converted to string
	defer resp.Body.Close()

	//read body into a bunch type of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		t.Fatal(err)
	}

	//convert bytes to string
	respString := string(b)
	expected := `Halo dunia!`

	if respString != expected {
		t.Errorf("handler return unexpected body : got %v want %v", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status should be 405, but got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)

	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

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