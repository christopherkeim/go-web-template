package server

import "github.com/go-chi/chi/v5"

func setRoutes(router *chi.Mux) {
	// GET /
	router.Get("/", GetRoot)
	// GET /users/{guid}
	router.Get("/users/{guid}", GetSingleUser)
	// GET /users
	router.Get("/users", GetAllUsers)
	// POST /users
	router.Post("/users", CreateNewUser)
}
