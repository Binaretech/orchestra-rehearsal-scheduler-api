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
type Middleware func(http.ResponseWriter, *http.Request, http.HandlerFunc) error

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

// Get allows registering a GET route on the group.
func (g *Group) Get(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodGet, path, handlerFunc, middlewares...)
}

// Post allows registering a POST route on the group.
func (g *Group) Post(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodPost, path, handlerFunc, middlewares...)
}

// Put allows registering a PUT route on the group.
func (g *Group) Put(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodPut, path, handlerFunc, middlewares...)
}

// Delete allows registering a DELETE route on the group.
func (g *Group) Delete(path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodDelete, path, handlerFunc, middlewares...)
}

// addRoute adds a new route to the group.
func (g *Group) addRoute(method, path string, handlerFunc http.HandlerFunc, middlewares ...Middleware) {
	route := Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
		Middlewares: middlewares,
	}

	g.Routes = append(g.Routes, route)
}

// Router is an abstraction over the default ServeMux.
type Router struct {
	routes            []Route      // Registered routes
	globalMiddlewares []Middleware // Global middlewares applied to all routes
	groups            []*Group     // Route groups with prefixes
}

// New creates a new instance of Router.
func New() *Router {
	return &Router{
		globalMiddlewares: []Middleware{}, // Initialize without global middlewares
		groups:            []*Group{},     // Initialize without route groups
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
	r.groups = append(r.groups, group)
	return group
}

// RegisterRoutes registers all routes and applies middlewares, both global and group-specific.
func (r *Router) RegisterRoutes(mux *http.ServeMux) {
	for _, group := range r.groups {
		for _, route := range group.Routes {
			handler := route.HandlerFunc

			handler = wrapWithMiddlewares(handler, r.globalMiddlewares...)
			handler = wrapWithMiddlewares(handler, group.Middlewares...)
			handler = wrapWithMiddlewares(handler, route.Middlewares...)

			mux.HandleFunc(fmt.Sprintf("%s%s", group.Prefix, route.Path), handler)
		}
	}

	for _, route := range r.routes {
		handler := route.HandlerFunc
		handler = wrapWithMiddlewares(handler, r.globalMiddlewares...)
		handler = wrapWithMiddlewares(handler, route.Middlewares...)
		mux.HandleFunc(route.Path, handler)
	}
}

// wrapMiddleware wraps a single middleware function around the handler.
func wrapWithMiddlewares(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var next http.HandlerFunc

		next = handler

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			next = func(handler http.HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					if err := middleware(w, r, handler); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}(next)
		}

		// Execute the composed middlewares and finally the handler
		next(w, r)
	})
}

// AddMiddleware adds a global middleware that applies to all routes.
func (r *Router) AddMiddleware(mw Middleware) {
	r.globalMiddlewares = append(r.globalMiddlewares, mw)
}
