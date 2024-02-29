package server

import "github.com/go-chi/chi/v5"

func setRoutes(router *chi.Mux) {
	db := initDb()
	// GET /
	router.Get("/", GetRoot)
	// GET /users/{guid}
	router.Get("/users/{guid}", GetSingleUser(&db))
	// GET /users
	router.Get("/users", GetAllUsers(&db))
	// POST /users
	router.Post("/users", CreateNewUser(&db))
}
