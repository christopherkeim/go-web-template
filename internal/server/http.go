package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

type GopherData struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

type ValidationResponse struct {
	GopherName string `json:"gopher_name"`
	State      string `json:"state"`
}

func getRoot(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello 🐍 🚀 ✨\n"))
}

func postGopherData(writer http.ResponseWriter, request *http.Request) {
	if dayOfWeek := chi.URLParam(request, "day_of_week"); dayOfWeek != "" {
		// Stdout on server
		fmt.Println(dayOfWeek)
	}
	// Parse the body and extract the GopherData properties
	request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
	decoder := json.NewDecoder((request.Body))
	var gopher GopherData
	parsingError := decoder.Decode(&gopher)

	if parsingError != nil {
		fmt.Printf("Parsing Error: %s", parsingError.Error())
		return
	}

	// Construct the validation response
	validationResponse := ValidationResponse{GopherName: gopher.Name, State: "success"}
	jsonResponse, jsonError := json.Marshal(validationResponse)

	if jsonError != nil {
		fmt.Printf("JSON Error: %s", jsonError.Error())
		return
	}

	// Send the response back to the client
	writer.Write([]byte(jsonResponse))
}

func RunHTTPServerOnAddress(address string) {
	router := chi.NewRouter()
	setMiddlewares(router)

	// GET endpoint
	router.Get("/", getRoot)
	// POST endpoint
	router.Post("/gopher", postGopherData)

	logrus.Info("Starting HTTP server")

	err := http.ListenAndServe(address, router)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}
