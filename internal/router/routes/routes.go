package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	middlewares "github.com/luisberga/phones-api/internal/middleware"
	"github.com/rs/cors"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Config(r *mux.Router) http.Handler {
	routes := CompaniesRoute
	routes = append(routes, PhonesRoute...)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
	})

	handler := c.Handler(r)

	return handler

}
