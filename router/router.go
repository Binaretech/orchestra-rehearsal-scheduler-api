// Package router provides a simple and flexible routing mechanism for HTTP servers.
// It allows defining routes with specific HTTP methods, URL paths, and handler functions.
// Additionally, it supports middleware for both individual routes and groups of routes
// with common prefixes.
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

// addRoute adds a new route to the router.
func (r *Router) addRoute(method, path string, handlerFunc HandlerFunc, middlewares ...Middleware) {
	route := Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
		Middlewares: middlewares,
	}

	fmt.Printf("Registered %s %s", method, path)

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

			handler = wrapWithMiddlewares(handler, append(r.globalMiddlewares, group.Middlewares...)...)
			handler = wrapWithMiddlewares(handler, route.Middlewares...)

			finalHandler := func(h HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, req *http.Request) {
					ctx := &Context{
						request:  req,
						response: w,
					}
					if err := h(ctx); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
				}
			}(handler)

			mux.HandleFunc(fmt.Sprintf("%s%s", group.Prefix, route.Path), finalHandler)
		}
	}

	for _, route := range r.routes {
		handler := route.HandlerFunc
		handler = wrapWithMiddlewares(handler, append(r.globalMiddlewares, route.Middlewares...)...)

		finalHandler := func(h HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, req *http.Request) {
				ctx := &Context{
					request:  req,
					response: w,
				}
				if err := h(ctx); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}(handler)

		mux.HandleFunc(route.Path, finalHandler)
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

// AddMiddleware adds a global middleware that applies to all routes.
func (r *Router) AddMiddleware(mw Middleware) {
	r.globalMiddlewares = append(r.globalMiddlewares, mw)
}

type Context struct {
	request  *http.Request
	response http.ResponseWriter
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Response() http.ResponseWriter {
	return c.response
}

func (c *Context) JSON(status int, data any) error {
	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(status)
	return json.NewEncoder(c.response).Encode(data)
}
