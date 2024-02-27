package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Helper function for loading data store
func withDb() Db {
	return initDb()
}

func GetRoot(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello 🦫 🚀 ✨\n"))
}

func GetSingleUser(writer http.ResponseWriter, request *http.Request) {

	staticUser := User{
		Guid:      "abcd123",
		FirstName: "Bob",
		LastName:  "Smith",
		Age:       31,
		Admin:     false,
	}

	jsonResponse, jsonError := json.Marshal(staticUser)
	if jsonError != nil {
		fmt.Printf("JSON Error: %s", jsonError.Error())
		return
	}

	writer.Write([]byte(jsonResponse))

}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	users := initDb()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Printf("JSON Error: %s", jsonError.Error())
		return
	}

	writer.Write([]byte(jsonResponse))

}

func CreateNewUser(writer http.ResponseWriter, request *http.Request) {
	// Parse the body and extract the GopherData properties
	request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
	decoder := json.NewDecoder((request.Body))
	var user User
	parsingError := decoder.Decode(&user)

	if parsingError != nil {
		fmt.Printf("Parsing Error: %s", parsingError.Error())
		return
	}

	// Construct the validation response
	validationResponse := UserCreationResponse{UserGuid: user.Guid, State: "success"}
	jsonResponse, jsonError := json.Marshal(validationResponse)

	if jsonError != nil {
		fmt.Printf("JSON Error: %s", jsonError.Error())
		return
	}

	// Send the response back to the client
	writer.Write([]byte(jsonResponse))
}
