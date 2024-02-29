package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRoot(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello 🦫 🚀 ✨\n"))
}

func GetSingleUser(db *Db) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		guid := chi.URLParam(request, "guid")

		foundUser := false
		var targetUser User
		for _, user := range *db {
			if guid == user.Guid {
				foundUser = true
				targetUser = user
				break
			}
		}

		if !foundUser {
			writer.Write([]byte(fmt.Sprintf("User not found: %d", http.StatusNotFound)))
			return
		}

		jsonResponse, jsonError := json.Marshal(targetUser)
		if jsonError != nil {
			fmt.Printf("JSON Error: %s", jsonError.Error())
			writer.Write([]byte(fmt.Sprintf("Internal server error: %d", http.StatusInternalServerError)))
			return
		}

		writer.Write([]byte(jsonResponse))
	}

}

func GetAllUsers(db *Db) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		jsonResponse, jsonError := json.Marshal(*db)
		if jsonError != nil {
			fmt.Printf("JSON Error: %s", jsonError.Error())
			writer.Write([]byte(fmt.Sprintf("Internal server error: %d", http.StatusInternalServerError)))
			return
		}

		writer.Write([]byte(jsonResponse))
	}

}

func CreateNewUser(db *Db) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Parse the body and extract the GopherData properties
		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
		decoder := json.NewDecoder((request.Body))
		var user User
		parsingError := decoder.Decode(&user)

		if parsingError != nil {
			fmt.Printf("Parsing Error: %s", parsingError.Error())
			writer.Write([]byte(fmt.Sprintf("Internal server error: %d", http.StatusInternalServerError)))
			return
		}

		*db = append(*db, user)

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
}
