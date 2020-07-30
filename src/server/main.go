package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/scrummer123/golang-portfolio/src/server/database"
	"github.com/scrummer123/golang-portfolio/src/server/handlers"
	"github.com/scrummer123/golang-portfolio/src/server/helpers"
	"github.com/scrummer123/golang-portfolio/src/server/token"
)

func init() {
	token.Init()
}

func main() {
	// Chi stuff
	r := chi.NewRouter()
	r.Mount("/api", apiRoutes(r))
	r.Mount("/app", frontEndRoutes(r))
	// End chi stuff

	// Fire stuff
	database.InitializeFirestore()
	// End fire stuff

	err := http.ListenAndServe("127.0.0.1:8080", r)

	log.Printf("Listening on port 8080")

	if err != nil {
		log.Fatalf("Error while serving => %v", err)
	}
}

func frontEndRoutes(r chi.Router) chi.Router {
	r.Get("/", handlers.ReactApp)

	return r
}

// api routes
func apiRoutes(r chi.Router) chi.Router {
	r.Mount("/users", userRoutes())
	r.Mount("/auth", authRoutes())

	return r
}

func authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/new", handlers.Register)
	r.Post("/{id}", handlers.Login)

	return r
}

// post app routes
func userRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.AllUsers)
	r.Group(func(r chi.Router) {
		r.Use(helpers.UserAuth)

		r.Get("/{id}/find", handlers.UserByID)
		r.Put("/{id}/update", handlers.UpdateUser)
		r.Delete("/{id}/delete", handlers.DeleteUser)
	})

	return r
}