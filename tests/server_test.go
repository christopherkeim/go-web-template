package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christopherkeim/go-web-template/internal/server"
	"github.com/go-chi/chi/v5"
)

func setupMockUsers() []server.User {
	mockUsers := make([]server.User, 0, 5)
	for i := range 5 {
		newUser := server.User{
			Guid:      fmt.Sprintf("dummy%d", (i+1)*100),
			FirstName: "User",
			LastName:  fmt.Sprintf("%d", i),
			Age:       (i * 3) + 10,
			Admin:     (i%2 == 0),
		}
		mockUsers = append(mockUsers, newUser)
	}
	return mockUsers
}

func addChiURLParams(request *http.Request, params map[string]string) *http.Request {
	ctx := chi.NewRouteContext()
	for k, v := range params {
		ctx.URLParams.Add(k, v)
	}

	return request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))
}

func TestGetSingleUser(t *testing.T) {
	mockUsers := setupMockUsers()

	request, err := http.NewRequest("GET", "http://0.0.0.0:8000/users/{guid}", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %s", err.Error())
	}

	params := map[string]string{"guid": "dummy300"}
	requestWithParms := addChiURLParams(request, params)

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetSingleUser(&mockUsers))
	handler.ServeHTTP(requestRecorder, requestWithParms)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	want := `{"guid":"dummy300","first_name":"User","last_name":"2","age":16,"admin":true}`
	if response := requestRecorder.Body.String(); response != want {
		t.Errorf("Handler returned wrong JSON. Expected: %s, Got: %s", want, response)
	}
}

func TestGetAllUsers(t *testing.T) {
	mockUsers := setupMockUsers()

	request, err := http.NewRequest("GET", "http://0.0.0.0:8000/users", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %s", err.Error())
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetAllUsers(&mockUsers))
	handler.ServeHTTP(requestRecorder, request)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	want, _ := json.Marshal(mockUsers)

	if response := requestRecorder.Body.String(); response != string(want) {
		t.Errorf("Handler returned wrong JSON. Expected: %s, Got: %s", want, response)
	}
}

func TestPostHandler(t *testing.T) {
	mockUsers := setupMockUsers()

	bodyData := server.User{
		Guid:      "abcd123",
		FirstName: "Bob",
		LastName:  "Smith",
		Age:       31,
		Admin:     false,
	}

	jsonBody, err := json.Marshal(&bodyData)
	if err != nil {
		t.Errorf("Error creating json: %s", err.Error())
	}

	request, err := http.NewRequest("POST", "http://0.0.0.0:8000/users", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf("Error creating a new request: %s", err.Error())
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(server.CreateNewUser(&mockUsers))
	handler.ServeHTTP(requestRecorder, request)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	want := `{"user_guid":"abcd123","state":"success"}`
	if response := requestRecorder.Body.String(); response != want {
		t.Errorf("Handler returned wrong string. Expected: %s, Got: %s", want, response)
	}
}
