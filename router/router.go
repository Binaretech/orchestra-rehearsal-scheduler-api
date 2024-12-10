// Package router provides a simple and flexible routing mechanism for HTTP servers.
// It allows defining routes with specific HTTP methods, URL paths, and handler functions.
// Additionally, it supports middleware for both individual routes and groups of routes
// with common prefixes.
// Package router provides a simple and flexible routing mechanism for HTTP servers.
package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Middleware defines the type for middleware functions.
type Middleware func(*Context, HandlerFunc) error

// HandlerFunc defines the type for handler functions that use Context.
type HandlerFunc func(*Context) error

// ErrorHandler defines the type for centralized error handling.
type ErrorHandler func(*Context, error)

// Route represents a single route.
type Route struct {
	Method      string       // HTTP method (GET, POST, PUT, DELETE, etc.)
	Path        string       // URL path for the route
	HandlerFunc HandlerFunc  // Handler function to process the request
	Middlewares []Middleware // Middlewares specific to this route
}

// Group represents a group of routes with a common prefix and shared middlewares.
type Group struct {
	Prefix      string       // Common prefix for all routes in the group
	Middlewares []Middleware // Specific middlewares for all routes in this group
	Routes      []Route      // Specific routes of the group
}

// Router is an abstraction over the default ServeMux.
type Router struct {
	routes            []Route      // Registered routes
	globalMiddlewares []Middleware // Global middlewares applied to all routes
	groups            []*Group     // Route groups with prefixes
	errorHandler      ErrorHandler // Custom error handler
}

// Get allows registering a GET route.
func (r *Router) Get(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodGet, path, handlerFunc, middlewares...)
}

// Post allows registering a POST route.
func (r *Router) Post(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodPost, path, handlerFunc, middlewares...)
}

// Put allows registering a PUT route.
func (r *Router) Put(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodPut, path, handlerFunc, middlewares...)
}

// Delete allows registering a DELETE route.
func (r *Router) Delete(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	r.addRoute(http.MethodDelete, path, handlerFunc, middlewares...)
}

// Get allows registering a GET route on the group.
func (g *Group) Get(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodGet, path, handlerFunc, middlewares...)
}

// Post allows registering a POST route on the group.
func (g *Group) Post(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodPost, path, handlerFunc, middlewares...)
}

// Put allows registering a PUT route on the group.
func (g *Group) Put(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodPut, path, handlerFunc, middlewares...)
}

// Delete allows registering a DELETE route on the group.
func (g *Group) Delete(path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	g.addRoute(http.MethodDelete, path, handlerFunc, middlewares...)
}

// addRoute adds a new route to the group.
func (g *Group) addRoute(method, path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	route := Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
		Middlewares: middlewares,
	}

	g.Routes = append(g.Routes, route)
}

// Context holds the request and response objects for a route.
type Context struct {
	request  *http.Request
	response http.ResponseWriter
}

// New creates a new instance of Router.
func New() *Router {
	return &Router{
		globalMiddlewares: []Middleware{},
		groups:            []*Group{},
		errorHandler: func(c *Context, err error) {
			http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		},
	}
}

// AddMiddleware adds a global middleware that applies to all routes.
func (r *Router) AddMiddleware(mw Middleware) {
	r.globalMiddlewares = append(r.globalMiddlewares, mw)
}

// SetErrorHandler allows configuring a custom error handler.
func (r *Router) SetErrorHandler(handler ErrorHandler) {
	r.errorHandler = handler
}

// addRoute adds a new route to the router.
func (r *Router) addRoute(method, path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	route := Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
		Middlewares: middlewares,
	}

	r.routes = append(r.routes, route)
}

// Group creates a new group of routes with a common prefix and specific middlewares.
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
			finalHandler := r.createHandler(route.HandlerFunc, append(r.globalMiddlewares, group.Middlewares...)...)
			fmt.Printf("Registered %s %s%s\n", route.Method, group.Prefix, route.Path)

			mux.HandleFunc(fmt.Sprintf("%s %s%s", route.Method, group.Prefix, route.Path), finalHandler)
		}
	}

	for _, route := range r.routes {
		finalHandler := r.createHandler(route.HandlerFunc, append(r.globalMiddlewares, route.Middlewares...)...)
		fmt.Printf("Registered %s %s\n", route.Method, route.Path)
		mux.HandleFunc(fmt.Sprintf("%s %s", route.Method, route.Path), finalHandler)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "error/not-found"})

	})
}

// createHandler wraps a HandlerFunc with middlewares and the error handler.
func (r *Router) createHandler(handler HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	handler = wrapWithMiddlewares(handler, middlewares...)
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := &Context{
			request:  req,
			response: w,
		}

		if err := handler(ctx); err != nil {
			r.errorHandler(ctx, err)
		}
	}
}

// wrapWithMiddlewares wraps the handler with the provided middlewares.
func wrapWithMiddlewares(handler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]
		next := handler
		handler = func(c *Context) error {
			return middleware(c, next)
		}
	}
	return handler
}

func (c *Context) Param(key string) string {
	return c.request.PathValue(key)
}

func (c *Context) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Response() http.ResponseWriter {
	return c.response
}

func (c *Context) Error(status int, message string) error {
	return c.JSON(status, map[string]string{
		"error": message,
	})
}

func (c *Context) JSON(status int, data any) error {
	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(status)
	return json.NewEncoder(c.response).Encode(data)
}

func (c *Context) Parse(data any) error {

	if c.request.ContentLength == 0 {
		return nil
	}

	return json.NewDecoder(c.request.Body).Decode(data)
}
