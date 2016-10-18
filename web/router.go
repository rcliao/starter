package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is to define client dependencies injection
type Handler func(*Client) http.HandlerFunc

// Route defines a route with its name, pattern, method and its handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc Handler
}

// Routes wrap a list of route in a typed model
type Routes []Route

// NewRouter is constructor pattern to create a new router mux
func NewRouter(client *Client) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc(client)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
