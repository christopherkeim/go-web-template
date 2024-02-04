package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christopherkeim/go-web-template/internal/server"
)

func TestGetHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://0.0.0.0:8000/", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %s", err.Error())
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetRoot)
	handler.ServeHTTP(requestRecorder, request)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	want := "Hello 🦫 🚀 ✨\n"
	if response := requestRecorder.Body.String(); response != want {
		t.Errorf("Handler returned wrong string. Expected: %s, Got: %s", want, response)
	}
}

func TestPostHandler(t *testing.T) {
	bodyData := server.GopherData{
		Name:        "bob",
		Age:         31,
		Description: "Had a good day.",
	}

	jsonBody, err := json.Marshal(&bodyData)
	if err != nil {
		t.Errorf("Error creating json: %s", err.Error())
	}

	request, err := http.NewRequest("POST", "http://0.0.0.0:8000/gopher", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf("Error creating a new request: %s", err.Error())
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(server.PostGopherData)
	handler.ServeHTTP(requestRecorder, request)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	want := `{"gopher_name":"bob","state":"success"}`
	if response := requestRecorder.Body.String(); response != want {
		t.Errorf("Handler returned wrong string. Expected: %s, Got: %s", want, response)
	}
}
