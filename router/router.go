// Package router provides a simple and flexible routing mechanism for HTTP servers.
// It allows defining routes with specific HTTP methods, URL paths, and handler functions.
// Additionally, it supports middleware for both individual routes and groups of routes
// with common prefixes.
package router

import (
	"fmt"
	"net/http"
)

// Middleware defines the type for middleware functions.
type Middleware func(http.Handler) http.Handler

// Route represents a single route.
type Route struct {
	Method      string           // HTTP method (GET, POST, PUT, DELETE, etc.)
	Path        string           // URL path for the route
	HandlerFunc http.HandlerFunc // Handler function to process the request
	Middlewares []Middleware     // Middlewares specific to this route
}

// Group represents a group of routes with a common prefix and shared middlewares.
type Group struct {
	Prefix      string       // Common prefix for all routes in the group
	Middlewares []Middleware // Specific middlewares for all routes in this group
	Routes      []Route      // Specific routes of the group
}

// Router is an abstraction over the default ServeMux.
type Router struct {
	mux               *http.ServeMux // HTTP mux to handle routes
	routes            []Route        // Registered routes
	globalMiddlewares []Middleware   // Global middlewares applied to all routes
	groups            []Group        // Route groups with prefixes
}

// New creates a new instance of Router.
func New(mux *http.ServeMux) *Router {
	return &Router{
		mux:               mux,
		globalMiddlewares: []Middleware{}, // Initialize without global middlewares
		groups:            []Group{},      // Initialize without route groups
	}
}

// Get allows registering a GET route.
func (r *Router) Get(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodGet, path, handlerFunc, middlewares...)
}

// Post allows registering a POST route.
func (r *Router) Post(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodPost, path, handlerFunc, middlewares...)
}

// Put allows registering a PUT route.
func (r *Router) Put(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodPut, path, handlerFunc, middlewares...)
}

// Delete allows registering a DELETE route.
func (r *Router) Delete(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodDelete, path, handlerFunc, middlewares...)
}

// addRoute adds a new route to the router.
func (r *Router) addRoute(method, path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	route := Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
		Middlewares: middlewares,
	}

	r.routes = append(r.routes, route)
}

// Group allows creating a group of routes with a common prefix and specific middlewares.
func (r *Router) Group(prefix string, middlewares ...Middleware) *Group {
	group := &Group{
		Prefix:      prefix,
		Middlewares: middlewares,
	}
	r.groups = append(r.groups, *group)
	return group
}

// RegisterRoutes registers all routes and applies middlewares, both global and group-specific.
func (r *Router) RegisterRoutes() {
	// Register group routes first
	for _, group := range r.groups {
		// Create group routes with their prefix and middlewares
		for _, route := range group.Routes {
			handler := route.HandlerFunc

			// Apply global middlewares
			for _, mw := range r.globalMiddlewares {
				handler = wrapMiddleware(handler, mw)
			}

			// Apply group middlewares
			for _, mw := range group.Middlewares {
				handler = wrapMiddleware(handler, mw)
			}

			// Register the group route in the mux with its prefix
			r.mux.HandleFunc(fmt.Sprintf("%s %s", route.Method, group.Prefix+route.Path), handler)
		}
	}

	// Register routes without group prefix
	for _, route := range r.routes {
		handler := route.HandlerFunc

		// Apply global middlewares
		for _, mw := range r.globalMiddlewares {
			handler = wrapMiddleware(handler, mw)
		}

		// Register the route without group
		r.mux.HandleFunc(fmt.Sprintf("%s %s", route.Method, route.Path), handler)
	}
}

// wrapMiddleware applies the middleware to a handler and performs the appropriate conversion.
func wrapMiddleware(handler http.HandlerFunc, mw Middleware) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Convert http.HandlerFunc to http.Handler and then apply the middleware
		mw(http.HandlerFunc(handler)).ServeHTTP(w, r)
	})
}

// HandleRequest starts the server and listens for requests on the given port.
func (r *Router) HandleRequest(addr string) error {
	r.RegisterRoutes()
	return http.ListenAndServe(addr, r.mux)
}

// AddMiddleware adds a global middleware that applies to all routes.
func (r *Router) AddMiddleware(mw Middleware) {
	r.globalMiddlewares = append(r.globalMiddlewares, mw)
}

// Listen starts the server and listens for requests on the given port.
func (r *Router) Listen(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
